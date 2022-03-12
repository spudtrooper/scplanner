package api

//go:generate genopts --prefix=TradeContractsSearch --outfile=api/tradecontractssearchoptions.go "page:int" "size:int" "dsp:string" "genre:string" "minimumFollowers:int" "maximumFollowers:int"

type TradeContractsSearchOption func(*tradeContractsSearchOptionImpl)

type TradeContractsSearchOptions interface {
	Page() int
	Size() int
	Dsp() string
	Genre() string
	MinimumFollowers() int
	MaximumFollowers() int
}

func TradeContractsSearchPage(page int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.page = page
	}
}
func TradeContractsSearchPageFlag(page *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.page = *page
	}
}

func TradeContractsSearchSize(size int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.size = size
	}
}
func TradeContractsSearchSizeFlag(size *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.size = *size
	}
}

func TradeContractsSearchDsp(dsp string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.dsp = dsp
	}
}
func TradeContractsSearchDspFlag(dsp *string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.dsp = *dsp
	}
}

func TradeContractsSearchGenre(genre string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.genre = genre
	}
}
func TradeContractsSearchGenreFlag(genre *string) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.genre = *genre
	}
}

func TradeContractsSearchMinimumFollowers(minimumFollowers int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.minimumFollowers = minimumFollowers
	}
}
func TradeContractsSearchMinimumFollowersFlag(minimumFollowers *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.minimumFollowers = *minimumFollowers
	}
}

func TradeContractsSearchMaximumFollowers(maximumFollowers int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.maximumFollowers = maximumFollowers
	}
}
func TradeContractsSearchMaximumFollowersFlag(maximumFollowers *int) TradeContractsSearchOption {
	return func(opts *tradeContractsSearchOptionImpl) {
		opts.maximumFollowers = *maximumFollowers
	}
}

type tradeContractsSearchOptionImpl struct {
	page             int
	size             int
	dsp              string
	genre            string
	minimumFollowers int
	maximumFollowers int
}

func (t *tradeContractsSearchOptionImpl) Page() int             { return t.page }
func (t *tradeContractsSearchOptionImpl) Size() int             { return t.size }
func (t *tradeContractsSearchOptionImpl) Dsp() string           { return t.dsp }
func (t *tradeContractsSearchOptionImpl) Genre() string         { return t.genre }
func (t *tradeContractsSearchOptionImpl) MinimumFollowers() int { return t.minimumFollowers }
func (t *tradeContractsSearchOptionImpl) MaximumFollowers() int { return t.maximumFollowers }

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
