package api

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/scplanner/log"
)

var (
	userID       = flags.String("user_id", "user ID")
	token        = flags.String("token", "auth token")
	userCreds    = flag.String("user_creds", ".user_creds.json", "file with user credentials")
	clientDebug  = flags.Bool("client_debug", "whether to debug requests")
	requestStats = flags.Bool("request_stats", "print verbose debugging of request timing")
)

// type core represents the core gettr core
type core struct {
	userID string
	jwt    string
	debug  bool
}

func MakeClientFromFlags() (*core, error) {
	if *userID != "" && *token != "" {
		client := MakeClient(*userID, *token, MakeClientDebug(*clientDebug))
		return client, nil
	}
	if *userCreds != "" {
		client, err := MakeClientFromFile(*userCreds, MakeClientDebug(*clientDebug))
		if err != nil {
			return nil, err
		}
		return client, nil
	}
	return nil, errors.Errorf("Must set --user & --token or --creds_file")
}

//go:generate genopts --function MakeClient "debug:bool"
func MakeClient(userID, token string, mOpts ...MakeClientOption) *core {
	opts := MakeMakeClientOptions(mOpts...)
	return &core{
		userID: userID,
		jwt:    token,
		debug:  opts.Debug(),
	}
}

type Creds struct {
	UserID string `json:"userId"`
	Token  string `json:"token"`
}

func ReadCredsFromFlags() (Creds, error) {
	return readCreds(*userCreds)
}

func WriteCredsFromFlags(creds Creds) error {
	b, err := json.Marshal(&creds)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(*userCreds, b, 0755); err != nil {
		return err
	}
	log.Printf("wrote to %s", *userCreds)
	return nil
}

func readCreds(credsFile string) (creds Creds, ret error) {
	credsBytes, err := ioutil.ReadFile(credsFile)
	if err != nil {
		ret = err
		return
	}
	if err := json.Unmarshal(credsBytes, &creds); err != nil {
		ret = err
		return
	}
	return
}

func MakeClientFromFile(credsFile string, mOpts ...MakeClientOption) (*core, error) {
	opts := MakeMakeClientOptions(mOpts...)
	creds, err := readCreds(credsFile)
	if err != nil {
		return nil, err
	}
	return &core{
		userID: creds.UserID,
		jwt:    creds.Token,
		debug:  opts.Debug(),
	}, nil
}

type param struct {
	key string
	val interface{}
}

//go:generate genopts --function Request "extraHeaders:map[string]string" "host:string" "customPayload:interface{}"

func createRoute(base string, ps ...param) string {
	if len(ps) == 0 {
		return base
	}
	var ss []string
	for _, p := range ps {
		s := fmt.Sprintf("%s=%s", p.key, url.QueryEscape(fmt.Sprintf("%v", p.val)))
		ss = append(ss, s)
	}
	return fmt.Sprintf("%s?%s", base, strings.Join(ss, "&"))
}

func (c *core) get(route string, result interface{}, rOpts ...RequestOption) (*http.Response, error) {
	return c.request("GET", route, result, nil, rOpts...)
}

// TODO: Move body to a RequestOption
func (c *core) post(route string, result interface{}, body io.Reader, rOpts ...RequestOption) (*http.Response, error) {
	return c.request("POST", route, result, body, rOpts...)
}

// TODO: Move body to a RequestOption
func (c *core) patch(route string, result interface{}, body io.Reader, rOpts ...RequestOption) (*http.Response, error) {
	return c.request("PATCH", route, result, body, rOpts...)
}

func (c *core) delete(route string, result interface{}, rOpts ...RequestOption) (*http.Response, error) {
	return c.request("DELETE", route, result, nil, rOpts...)
}

func (c *core) request(method, route string, result interface{}, body io.Reader, rOpts ...RequestOption) (*http.Response, error) {
	opts := MakeRequestOptions(rOpts...)
	host := or.String(opts.Host(), "scplanner.services.scplanner.net")
	url := fmt.Sprintf("https://%s/%s", host, route)

	start := time.Now()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	cookie := fmt.Sprintf("jwt=%s;", c.jwt)
	req.Header.Set("cookie", cookie)
	for k, v := range opts.ExtraHeaders() {
		req.Header.Set(k, v)
	}
	if c.debug {
		log.Printf("requesting %s %s", method, url)
		if len(opts.ExtraHeaders()) > 0 {
			log.Printf("  with extra headers:")
			for k, v := range opts.ExtraHeaders() {
				log.Printf("    %s: %s", k, v)
			}
		}
		log.Printf("  headers:")
		for k, v := range req.Header {
			log.Printf("    %s: %s", k, v)
		}
		log.Printf("  body: %v", body)
	}
	doRes, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reqStop := time.Now()

	data, err := ioutil.ReadAll(doRes.Body)
	if err != nil {
		return nil, err
	}

	doRes.Body.Close()

	readStop := time.Now()

	if *requestStats {
		reqDur := reqStop.Sub(start)
		readDur := readStop.Sub(reqStop)
		totalDur := readStop.Sub(start)
		log.Printf("request stats: total:%v request:%v read:%v", totalDur, reqDur, readDur)
	}

	if c.debug {
		prettyJSON, err := prettyPrintJSON(data)
		if err != nil {
			log.Printf("ignoring prettyPrintJSON error: %v", err)
			prettyJSON = string(data)
		}
		log.Printf("from route %q have response %s", route, prettyJSON)
	}

	if len(data) > 0 {
		if opts.CustomPayload() != nil {
			if err := json.Unmarshal(data, opts.CustomPayload()); err != nil {
				return nil, err
			}
		} else {
			if err := json.Unmarshal(data, &result); err != nil {
				return nil, err
			}
			if c.debug {
				log.Printf("got response: %+v", result)
			}
		}
	}

	return doRes, nil
}

func prettyPrintJSON(b []byte) (string, error) {
	b = []byte(strings.TrimSpace(string(b)))
	if len(b) == 0 {
		return "", nil
	}
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, b, "", "\t"); err != nil {
		return "", errors.Errorf("json.Indent: payload=%q: %v", string(b), err)
	}
	return prettyJSON.String(), nil
}

// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func jsonMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(t); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func formatAsJSON(x interface{}) (string, error) {
	b, err := json.Marshal(x)
	if err != nil {
		return "", err
	}
	res, err := prettyPrintJSON(b)
	if err != nil {
		return "", err
	}
	return res, nil
}
