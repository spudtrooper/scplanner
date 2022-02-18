package api

// genopts --opt_type=BidsOption --prefix=Bids --outfile=api/bidsoptions.go 'status:string' 'dsp:string' 'bidderUserId:string' 'lastEvaluatedKey:LastEvaluatedKeyInfo'

type BidsOption func(*bidsOptionImpl)

type BidsOptions interface {
	Status() string
	Dsp() string
	BidderUserId() string
	LastEvaluatedKey() LastEvaluatedKeyInfo
}

func BidsStatus(status string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.status = status
	}
}

func BidsDsp(dsp string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.dsp = dsp
	}
}

func BidsBidderUserId(bidderUserId string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.bidderUserId = bidderUserId
	}
}

func BidsLastEvaluatedKey(lastEvaluatedKey LastEvaluatedKeyInfo) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.lastEvaluatedKey = lastEvaluatedKey
	}
}

type bidsOptionImpl struct {
	status           string
	dsp              string
	bidderUserId     string
	lastEvaluatedKey LastEvaluatedKeyInfo
}

func (b *bidsOptionImpl) Status() string                         { return b.status }
func (b *bidsOptionImpl) Dsp() string                            { return b.dsp }
func (b *bidsOptionImpl) BidderUserId() string                   { return b.bidderUserId }
func (b *bidsOptionImpl) LastEvaluatedKey() LastEvaluatedKeyInfo { return b.lastEvaluatedKey }

func makeBidsOptionImpl(opts ...BidsOption) *bidsOptionImpl {
	res := &bidsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeBidsOptions(opts ...BidsOption) BidsOptions {
	return makeBidsOptionImpl(opts...)
}
