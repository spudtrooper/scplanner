// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type DeleteBidOption func(*deleteBidOptionImpl)

type DeleteBidOptions interface {
	Status() string
	HasStatus() bool
	Dsp() string
	HasDsp() bool
	BidderUserId() string
	HasBidderUserId() bool
	LastEvaluatedKey() LastEvaluatedKeyInfo
	HasLastEvaluatedKey() bool
}

func DeleteBidStatus(status string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.has_status = true
		opts.status = status
	}
}
func DeleteBidStatusFlag(status *string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		if status == nil {
			return
		}
		opts.has_status = true
		opts.status = *status
	}
}

func DeleteBidDsp(dsp string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.has_dsp = true
		opts.dsp = dsp
	}
}
func DeleteBidDspFlag(dsp *string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		if dsp == nil {
			return
		}
		opts.has_dsp = true
		opts.dsp = *dsp
	}
}

func DeleteBidBidderUserId(bidderUserId string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.has_bidderUserId = true
		opts.bidderUserId = bidderUserId
	}
}
func DeleteBidBidderUserIdFlag(bidderUserId *string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		if bidderUserId == nil {
			return
		}
		opts.has_bidderUserId = true
		opts.bidderUserId = *bidderUserId
	}
}

func DeleteBidLastEvaluatedKey(lastEvaluatedKey LastEvaluatedKeyInfo) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.has_lastEvaluatedKey = true
		opts.lastEvaluatedKey = lastEvaluatedKey
	}
}
func DeleteBidLastEvaluatedKeyFlag(lastEvaluatedKey *LastEvaluatedKeyInfo) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		if lastEvaluatedKey == nil {
			return
		}
		opts.has_lastEvaluatedKey = true
		opts.lastEvaluatedKey = *lastEvaluatedKey
	}
}

type deleteBidOptionImpl struct {
	status               string
	has_status           bool
	dsp                  string
	has_dsp              bool
	bidderUserId         string
	has_bidderUserId     bool
	lastEvaluatedKey     LastEvaluatedKeyInfo
	has_lastEvaluatedKey bool
}

func (d *deleteBidOptionImpl) Status() string                         { return d.status }
func (d *deleteBidOptionImpl) HasStatus() bool                        { return d.has_status }
func (d *deleteBidOptionImpl) Dsp() string                            { return d.dsp }
func (d *deleteBidOptionImpl) HasDsp() bool                           { return d.has_dsp }
func (d *deleteBidOptionImpl) BidderUserId() string                   { return d.bidderUserId }
func (d *deleteBidOptionImpl) HasBidderUserId() bool                  { return d.has_bidderUserId }
func (d *deleteBidOptionImpl) LastEvaluatedKey() LastEvaluatedKeyInfo { return d.lastEvaluatedKey }
func (d *deleteBidOptionImpl) HasLastEvaluatedKey() bool              { return d.has_lastEvaluatedKey }

type DeleteBidParams struct {
	Status           string               `json:"status"`
	Dsp              string               `json:"dsp"`
	BidderUserId     string               `json:"bidder_user_id"`
	LastEvaluatedKey LastEvaluatedKeyInfo `json:"last_evaluated_key"`
}

func (o DeleteBidParams) Options() []DeleteBidOption {
	return []DeleteBidOption{
		DeleteBidStatus(o.Status),
		DeleteBidDsp(o.Dsp),
		DeleteBidBidderUserId(o.BidderUserId),
		DeleteBidLastEvaluatedKey(o.LastEvaluatedKey),
	}
}

func makeDeleteBidOptionImpl(opts ...DeleteBidOption) *deleteBidOptionImpl {
	res := &deleteBidOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeDeleteBidOptions(opts ...DeleteBidOption) DeleteBidOptions {
	return makeDeleteBidOptionImpl(opts...)
}
