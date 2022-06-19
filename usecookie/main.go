// usecookie is a package to set the toke value in .user_creds.json
package usecookie

import (
	"flag"
	"strings"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/spudtrooper/scplanner/api"
	"github.com/spudtrooper/scplanner/log"
)

func findToken(cookie string) (string, error) {
	for _, part := range strings.Split(cookie, " ") {
		parts := strings.Split(part, "=")
		if len(parts) != 2 {
			return "", errors.Errorf("invalid cookie part: %s", part)
		}
		key, val := parts[0], parts[1]
		log.Printf("%s = %s", color.RedString(key), color.YellowString(val))
		if key != "jwt" {
			continue
		}
		if strings.HasSuffix(val, ";") {
			val = string(val[0 : len(val)-1])
		}
		return val, nil
	}
	return "", nil
}

func useCookie() error {
	cookie := flag.Arg(0)
	if cookie == "" {
		return errors.Errorf("cookie should be the first arg")
	}
	token, err := findToken(cookie)
	if err != nil {
		return errors.Errorf("findToken: %v", err)
	}
	if token == "" {
		return errors.Errorf("could not find token from cookie: %s", cookie)
	}
	log.Printf("token: %s", token)
	creds, err := api.ReadCredsFromFlags()
	if err != nil {
		return errors.Errorf("api.ReadCredsFromFlags: %v", err)
	}
	creds.Token = token
	if err := api.WriteCredsFromFlags(creds); err != nil {
		return err
	}
	return nil
}

func Main() error {
	return useCookie()
}
