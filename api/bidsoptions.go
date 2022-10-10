// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type BidsOption func(*bidsOptionImpl)

type BidsOptions interface {
	Status() string
	HasStatus() bool
	Dsp() string
	HasDsp() bool
	BidderUserId() string
	HasBidderUserId() bool
	LastEvaluatedKey() LastEvaluatedKeyInfo
	HasLastEvaluatedKey() bool
}

func BidsStatus(status string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.has_status = true
		opts.status = status
	}
}
func BidsStatusFlag(status *string) BidsOption {
	return func(opts *bidsOptionImpl) {
		if status == nil {
			return
		}
		opts.has_status = true
		opts.status = *status
	}
}

func BidsDsp(dsp string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.has_dsp = true
		opts.dsp = dsp
	}
}
func BidsDspFlag(dsp *string) BidsOption {
	return func(opts *bidsOptionImpl) {
		if dsp == nil {
			return
		}
		opts.has_dsp = true
		opts.dsp = *dsp
	}
}

func BidsBidderUserId(bidderUserId string) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.has_bidderUserId = true
		opts.bidderUserId = bidderUserId
	}
}
func BidsBidderUserIdFlag(bidderUserId *string) BidsOption {
	return func(opts *bidsOptionImpl) {
		if bidderUserId == nil {
			return
		}
		opts.has_bidderUserId = true
		opts.bidderUserId = *bidderUserId
	}
}

func BidsLastEvaluatedKey(lastEvaluatedKey LastEvaluatedKeyInfo) BidsOption {
	return func(opts *bidsOptionImpl) {
		opts.has_lastEvaluatedKey = true
		opts.lastEvaluatedKey = lastEvaluatedKey
	}
}
func BidsLastEvaluatedKeyFlag(lastEvaluatedKey *LastEvaluatedKeyInfo) BidsOption {
	return func(opts *bidsOptionImpl) {
		if lastEvaluatedKey == nil {
			return
		}
		opts.has_lastEvaluatedKey = true
		opts.lastEvaluatedKey = *lastEvaluatedKey
	}
}

type bidsOptionImpl struct {
	status               string
	has_status           bool
	dsp                  string
	has_dsp              bool
	bidderUserId         string
	has_bidderUserId     bool
	lastEvaluatedKey     LastEvaluatedKeyInfo
	has_lastEvaluatedKey bool
}

func (b *bidsOptionImpl) Status() string                         { return b.status }
func (b *bidsOptionImpl) HasStatus() bool                        { return b.has_status }
func (b *bidsOptionImpl) Dsp() string                            { return b.dsp }
func (b *bidsOptionImpl) HasDsp() bool                           { return b.has_dsp }
func (b *bidsOptionImpl) BidderUserId() string                   { return b.bidderUserId }
func (b *bidsOptionImpl) HasBidderUserId() bool                  { return b.has_bidderUserId }
func (b *bidsOptionImpl) LastEvaluatedKey() LastEvaluatedKeyInfo { return b.lastEvaluatedKey }
func (b *bidsOptionImpl) HasLastEvaluatedKey() bool              { return b.has_lastEvaluatedKey }

type BidsParams struct {
	Status           string               `json:"status"`
	Dsp              string               `json:"dsp"`
	BidderUserId     string               `json:"bidder_user_id"`
	LastEvaluatedKey LastEvaluatedKeyInfo `json:"last_evaluated_key"`
}

func (o BidsParams) Options() []BidsOption {
	return []BidsOption{
		BidsStatus(o.Status),
		BidsDsp(o.Dsp),
		BidsBidderUserId(o.BidderUserId),
		BidsLastEvaluatedKey(o.LastEvaluatedKey),
	}
}

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
