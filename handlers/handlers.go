// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/minimalcli/handler"
	"github.com/spudtrooper/opensecrets/api"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/scplanner/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Core) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("GetLegislator",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.GetLegislatorParams)
			return client.GetLegislator(p.Cid)
		},
		api.GetLegislatorParams{},
	)

	return b.Build()
}
