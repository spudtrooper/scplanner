// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type TradeContractsSearchOption func(*tradeContractsSearchOptionImpl)

type TradeContractsSearchOptions interface {
	Page() int
	HasPage() bool
	Size() int
	HasSize() bool
	Dsp() string
	HasDsp() bool
	Genre() string
	HasGenre() bool
	MinimumFollowers() int
	HasMinimumFollowers() bool
	MaximumFollowers() int
	HasMaximumFollowers() bool
}

func TradeContractsSearchPage(page int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_page = true
		opts.page = page
	}
}
func TradeContractsSearchPageFlag(page *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if page == nil {
			return
		}
		opts.has_page = true
		opts.page = *page
	}
}

func TradeContractsSearchSize(size int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_size = true
		opts.size = size
	}
}
func TradeContractsSearchSizeFlag(size *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if size == nil {
			return
		}
		opts.has_size = true
		opts.size = *size
	}
}

func TradeContractsSearchDsp(dsp string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_dsp = true
		opts.dsp = dsp
	}
}
func TradeContractsSearchDspFlag(dsp *string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if dsp == nil {
			return
		}
		opts.has_dsp = true
		opts.dsp = *dsp
	}
}

func TradeContractsSearchGenre(genre string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_genre = true
		opts.genre = genre
	}
}
func TradeContractsSearchGenreFlag(genre *string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if genre == nil {
			return
		}
		opts.has_genre = true
		opts.genre = *genre
	}
}

func TradeContractsSearchMinimumFollowers(minimumFollowers int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_minimumFollowers = true
		opts.minimumFollowers = minimumFollowers
	}
}
func TradeContractsSearchMinimumFollowersFlag(minimumFollowers *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if minimumFollowers == nil {
			return
		}
		opts.has_minimumFollowers = true
		opts.minimumFollowers = *minimumFollowers
	}
}

func TradeContractsSearchMaximumFollowers(maximumFollowers int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.has_maximumFollowers = true
		opts.maximumFollowers = maximumFollowers
	}
}
func TradeContractsSearchMaximumFollowersFlag(maximumFollowers *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		if maximumFollowers == nil {
			return
		}
		opts.has_maximumFollowers = true
		opts.maximumFollowers = *maximumFollowers
	}
}

type tradeContractsSearchOptionImpl struct {
	page                 int
	has_page             bool
	size                 int
	has_size             bool
	dsp                  string
	has_dsp              bool
	genre                string
	has_genre            bool
	minimumFollowers     int
	has_minimumFollowers bool
	maximumFollowers     int
	has_maximumFollowers bool
}

func (t *tradeContractsSearchOptionImpl) Page() int                 { return t.page }
func (t *tradeContractsSearchOptionImpl) HasPage() bool             { return t.has_page }
func (t *tradeContractsSearchOptionImpl) Size() int                 { return t.size }
func (t *tradeContractsSearchOptionImpl) HasSize() bool             { return t.has_size }
func (t *tradeContractsSearchOptionImpl) Dsp() string               { return t.dsp }
func (t *tradeContractsSearchOptionImpl) HasDsp() bool              { return t.has_dsp }
func (t *tradeContractsSearchOptionImpl) Genre() string             { return t.genre }
func (t *tradeContractsSearchOptionImpl) HasGenre() bool            { return t.has_genre }
func (t *tradeContractsSearchOptionImpl) MinimumFollowers() int     { return t.minimumFollowers }
func (t *tradeContractsSearchOptionImpl) HasMinimumFollowers() bool { return t.has_minimumFollowers }
func (t *tradeContractsSearchOptionImpl) MaximumFollowers() int     { return t.maximumFollowers }
func (t *tradeContractsSearchOptionImpl) HasMaximumFollowers() bool { return t.has_maximumFollowers }

type TradeContractsSearchParams struct {
	Page             int    `json:"page"`
	Size             int    `json:"size"`
	Dsp              string `json:"dsp"`
	Genre            string `json:"genre"`
	MinimumFollowers int    `json:"minimum_followers"`
	MaximumFollowers int    `json:"maximum_followers"`
}

func (o TradeContractsSearchParams) Options() []TradeContractsSearchOption {
	return []TradeContractsSearchOption{
		TradeContractsSearchPage(o.Page),
		TradeContractsSearchSize(o.Size),
		TradeContractsSearchDsp(o.Dsp),
		TradeContractsSearchGenre(o.Genre),
		TradeContractsSearchMinimumFollowers(o.MinimumFollowers),
		TradeContractsSearchMaximumFollowers(o.MaximumFollowers),
	}
}

func makeTradeContractsSearchOptionImpl(opts ...TradeContractsSearchOption) *tradeContractsSearchOptionImpl {
	res := &tradeContractsSearchOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeTradeContractsSearchOptions(opts ...TradeContractsSearchOption) TradeContractsSearchOptions {
	return makeTradeContractsSearchOptionImpl(opts...)
}
