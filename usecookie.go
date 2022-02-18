// Binary to set the toke value in .user_creds.json
// Usage:
//   - Copy the whole cookie value from the dev console
//   - Invoke this binary with that cookie as the only argument, e.g.
//
//       % scripts/usecookie.sh <cookie>
//
package main

import (
	"flag"
	"strings"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
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

func main() {
	flag.Parse()
	check.Err(useCookie())
}
