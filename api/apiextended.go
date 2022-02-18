package api

import (
	"math"
	"sync"

	"github.com/spudtrooper/goutil/or"
)

// type extended has functions the use the Core API to produce multiple results.
type extended struct {
	*core
}

func MakeExtended(c *core) *extended {
	return &extended{c}
}

type OffsetTradeContractsSearchResults struct {
	Offset  int
	Results []TradeContractsSearchResult
}

func (c *extended) TradeContractsSearches(tOpts ...TradeContractsSearchOption) (chan OffsetTradeContractsSearchResults, chan error, error) {
	initialInfo, err := c.TradeContractsSearch(tOpts...)
	if err != nil {
		return nil, nil, err
	}

	opts := MakeTradeContractsSearchOptions(tOpts...)
	start := or.Int(opts.Page(), 1)

	offsetResults := make(chan OffsetTradeContractsSearchResults)
	offsets := make(chan int)
	errs := make(chan error)

	go func() {
		limit := initialInfo.Pages
		for offset := start; offset < limit; offset++ {
			offsets <- offset
		}
		close(offsets)
	}()

	threads := 5

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for offset := range offsets {
					var cOpts []TradeContractsSearchOption
					cOpts = append(cOpts, tOpts...)
					cOpts = append(cOpts, TradeContractsSearchPage(offset))
					ress, err := c.TradeContractsSearch(cOpts...)
					if err != nil {
						errs <- err
						break
					}
					if len(ress.Results) == 0 {
						break
					}
					offsetResults <- OffsetTradeContractsSearchResults{
						Results: ress.Results,
						Offset:  offset,
					}
				}
			}()
		}
		wg.Wait()
		close(offsetResults)
		close(errs)
	}()

	return offsetResults, errs, nil
}

type OffsetBidsResults struct {
	Offset  int
	Results []BidsResult
}

func (c *extended) Bidss(bOpts ...BidsOption) (chan OffsetBidsResults, chan error, error) {
	offsetResults := make(chan OffsetBidsResults)
	offsets := make(chan int)
	errs := make(chan error)

	go func() {
		start := 1
		limit := math.MaxInt
		for offset := start; offset < limit; offset++ {
			offsets <- offset
		}
		close(offsets)
	}()

	go func() {
		var lastEvaluatedKey LastEvaluatedKeyInfo
		for offset := 0; ; offset++ {
			var cOpts []BidsOption
			cOpts = append(cOpts, bOpts...)
			if !lastEvaluatedKey.Empty() {
				cOpts = append(cOpts, BidsLastEvaluatedKey(lastEvaluatedKey))
			}
			ress, err := c.Bids(cOpts...)
			if err != nil {
				errs <- err
				break
			}
			if len(ress.Results) == 0 {
				break
			}
			offsetResults <- OffsetBidsResults{
				Results: ress.Results,
				Offset:  offset,
			}
			lastEvaluatedKey = ress.LastEvaluatedKey
			if lastEvaluatedKey.Empty() {
				break
			}
		}
		close(offsetResults)
		close(errs)
	}()

	return offsetResults, errs, nil
}
