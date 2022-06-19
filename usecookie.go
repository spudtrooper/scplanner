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

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/scplanner/usecookie"
)

func main() {
	flag.Parse()
	check.Err(usecookie.Main())
}
