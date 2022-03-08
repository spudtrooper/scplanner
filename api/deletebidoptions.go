package api

//go:generate genopts --opt_type=DeleteBidOption --prefix=DeleteBid --outfile=deletebidoptions.go "status:string" "dsp:string" "bidderUserId:string" "lastEvaluatedKey:LastEvaluatedKeyInfo"

type DeleteBidOption func(*deleteBidOptionImpl)

type DeleteBidOptions interface {
	Status() string
	Dsp() string
	BidderUserId() string
	LastEvaluatedKey() LastEvaluatedKeyInfo
}

func DeleteBidStatus(status string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.status = status
	}
}

func DeleteBidDsp(dsp string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.dsp = dsp
	}
}

func DeleteBidBidderUserId(bidderUserId string) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.bidderUserId = bidderUserId
	}
}

func DeleteBidLastEvaluatedKey(lastEvaluatedKey LastEvaluatedKeyInfo) DeleteBidOption {
	return func(opts *deleteBidOptionImpl) {
		opts.lastEvaluatedKey = lastEvaluatedKey
	}
}

type deleteBidOptionImpl struct {
	status           string
	dsp              string
	bidderUserId     string
	lastEvaluatedKey LastEvaluatedKeyInfo
}

func (d *deleteBidOptionImpl) Status() string                         { return d.status }
func (d *deleteBidOptionImpl) Dsp() string                            { return d.dsp }
func (d *deleteBidOptionImpl) BidderUserId() string                   { return d.bidderUserId }
func (d *deleteBidOptionImpl) LastEvaluatedKey() LastEvaluatedKeyInfo { return d.lastEvaluatedKey }

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
