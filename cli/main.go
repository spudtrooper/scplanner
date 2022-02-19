package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"strings"
	"sync"

	"github.com/pkg/errors"
	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/flags"
	"github.com/spudtrooper/goutil/parallel"
	minimalcli "github.com/spudtrooper/minimalcli/app"
	"github.com/spudtrooper/scplanner/api"
	"github.com/spudtrooper/scplanner/log"
)

var (
	page             = flags.Int("page", "global page")
	genre            = flags.String("genre", "global genre")
	id               = flags.String("id", "global ID")
	url              = flags.String("url", "global URL")
	maximumFollowers = flags.Int("maximum_followers", "global maximumFollowers")
	mininumFollowers = flags.Int("mininum_followers", "global mininumFollowers")
	threads          = flag.Int("threads", 10, "number of threads to use in async things")
)

func Main(ctx context.Context) error {
	app := minimalcli.Make()
	app.Init()

	coreClient, err := api.MakeClientFromFlags()
	if err != nil {
		return err
	}
	client := api.MakeExtended(coreClient)

	tradeContractsSearch := func() (*api.TradeContractsSearchInfo, error) {
		return client.TradeContractsSearch(
			api.TradeContractsSearchPage(*page),
			api.TradeContractsSearchGenre(*genre),
			api.TradeContractsSearchMaximumFollowers(*maximumFollowers),
			api.TradeContractsSearchMinimumFollowers(*mininumFollowers),
		)
	}

	tradeContractsSearches := func() (chan api.OffsetTradeContractsSearchResults, chan error, error) {
		return client.TradeContractsSearches(
			api.TradeContractsSearchPage(*page),
			api.TradeContractsSearchGenre(*genre),
			api.TradeContractsSearchMaximumFollowers(*maximumFollowers),
			api.TradeContractsSearchMinimumFollowers(*mininumFollowers),
		)
	}

	tradeContractsSearchesFlat := func() (chan api.TradeContractsSearchResult, chan error, error) {
		return client.TradeContractsSearchesFlat(
			api.TradeContractsSearchPage(*page),
			api.TradeContractsSearchGenre(*genre),
			api.TradeContractsSearchMaximumFollowers(*maximumFollowers),
			api.TradeContractsSearchMinimumFollowers(*mininumFollowers),
		)
	}

	app.Register("TradeContractsSearch", func(context.Context) error {
		info, err := tradeContractsSearch()
		if err != nil {
			return err
		}
		log.Printf("TradeContractsSearch: %s", mustFormatString(info))
		return nil
	})

	app.Register("TradeContractsSearches", func(context.Context) error {
		resultss, errs, err := tradeContractsSearches()
		if err != nil {
			return err
		}
		parallel.WaitFor(func() {
			for rs := range resultss {
				for _, r := range rs.Results {
					log.Printf("result: %d %v", rs.Offset, r.MinimumFollowers)
				}
			}
		}, func() {
			for e := range errs {
				log.Printf("error: %v", e)
			}
		})
		return nil
	})

	app.Register("Bids", func(context.Context) error {
		info, err := client.Bids()
		if err != nil {
			return err
		}
		log.Printf("Bids: %s", mustFormatString(info))
		return nil
	})

	app.Register("Bidss", func(context.Context) error {
		bidss, errs, err := client.Bidss()
		if err != nil {
			return err
		}
		parallel.WaitFor(func() {
			for rs := range bidss {
				for _, r := range rs.Results {
					log.Printf("result: %v %v", rs.Offset, r)
				}
			}
		}, func() {
			for e := range errs {
				log.Printf("error: %v", e)
			}
		})
		return nil
	})

	app.Register("Auth", func(context.Context) error {
		info, err := client.Auth()
		if err != nil {
			return err
		}
		log.Printf("Auth: %s", mustFormatString(info))
		return nil
	})

	app.Register("TradeContracts", func(context.Context) error {
		requireStringFlag(id, "id")
		info, err := client.TradeContracts(*id)
		if err != nil {
			return err
		}
		log.Printf("TradeContracts: %s", mustFormatString(info))
		return nil
	})

	app.Register("Resolve", func(context.Context) error {
		requireStringFlag(url, "url")
		info, err := client.Resolve(*url)
		if err != nil {
			return err
		}
		log.Printf("Resolve: %s", mustFormatString(info))
		return nil
	})

	app.Register("CreateBid", func(context.Context) error {
		requireStringFlag(id, "id")
		requireStringFlag(url, "url")

		authInfo, err := client.Auth()
		if err != nil {
			return err
		}

		resolvedTarget, err := client.Resolve(*url)
		if err != nil {
			return err
		}

		bidInfo, err := client.CreateBid(*id, *authInfo, *resolvedTarget)
		if err != nil {
			return err
		}
		log.Printf("Bid: %s", mustFormatString(bidInfo))

		return nil
	})

	app.Register("DeleteBid", func(context.Context) error {
		requireStringFlag(id, "id")

		bidInfo, err := client.DeleteBid(*id)
		if err != nil {
			return err
		}
		log.Printf("DeleteBid: %s", mustFormatString(bidInfo))

		return nil
	})

	app.Register("DeleteAllBids", func(context.Context) error {
		bidss, errs, err := client.Bidss()
		if err != nil {
			return err
		}
		parallel.WaitFor(func() {
			del := 0
			for rs := range bidss {
				for _, r := range rs.Results {
					id := r.ID
					if _, err := client.DeleteBid(id); err != nil {
						log.Printf("error: %v", err)
						continue
					}
					log.Printf("deleted[%d]: %s", del, id)
					del++
				}
			}
		}, func() {
			for e := range errs {
				log.Printf("error: %v", e)
			}
		})

		return nil
	})

	app.Register("SearchAndBid", func(context.Context) error {
		requireStringFlag(url, "url")

		authInfo, err := client.Auth()
		if err != nil {
			return err
		}

		existingSet := map[string]bool{}
		var numBids int
		bidss, errs, err := client.Bidss()
		if err != nil {
			return err
		}
		parallel.WaitFor(func() {
			for rs := range bidss {
				for _, r := range rs.Results {
					existingSet[r.ContractID] = true
					numBids++
				}
			}
		}, func() {
			for e := range errs {
				log.Printf("error: %v", e)
			}
		})
		log.Printf("have %d existing bids", numBids)
		log.Printf("have %d unique bids", len(existingSet))

		resolvedTarget, err := client.Resolve(*url)
		if err != nil {
			return err
		}

		resultss, errs, err := tradeContractsSearchesFlat()
		if err != nil {
			return err
		}
		parallel.WaitFor(func() {
			var wg sync.WaitGroup
			for i := 0; i < *threads; i++ {
				wg.Add(1)
				go func() {
					for r := range resultss {
						if exists := existingSet[r.ID]; exists {
							log.Printf("skipping %s because it already exists", r.ID)
							continue
						}
						bidInfo, err := client.CreateBid(r.ID, *authInfo, *resolvedTarget)
						if err != nil {
							log.Printf("error: %v", err)
						} else {
							log.Printf("created bid: %s", bidInfo.ID)
						}
					}
				}()
			}
			wg.Done()
		}, func() {
			for e := range errs {
				log.Printf("error: %v", e)
			}
		})

		return nil
	})

	if err := app.Run(ctx); err != nil {
		return err
	}

	return nil
}

func requireStringFlag(flag *string, name string) {
	if *flag == "" {
		log.Fatalf("--%s required", name)
	}
}

func mustFormatString(x interface{}) string {
	b, err := json.Marshal(x)
	check.Err(err)
	res, err := prettyPrintJSON(b)
	check.Err(err)
	return res
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
