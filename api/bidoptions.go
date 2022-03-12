package api

//go:generate genopts --prefix=Bid --outfile=api/bidoptions.go "dsp:string" "bidderUserId:string" "contractID:string" "debugBody"

type BidOption func(*bidOptionImpl)

type BidOptions interface {
	Dsp() string
	BidderUserId() string
	ContractID() string
	DebugBody() bool
}

func BidDsp(dsp string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.dsp = dsp
	}
}
func BidDspFlag(dsp *string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.dsp = *dsp
	}
}

func BidBidderUserId(bidderUserId string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.bidderUserId = bidderUserId
	}
}
func BidBidderUserIdFlag(bidderUserId *string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.bidderUserId = *bidderUserId
	}
}

func BidContractID(contractID string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.contractID = contractID
	}
}
func BidContractIDFlag(contractID *string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.contractID = *contractID
	}
}

func BidDebugBody(debugBody bool) BidOption {
	return func(opts *bidOptionImpl) {
		opts.debugBody = debugBody
	}
}
func BidDebugBodyFlag(debugBody *bool) BidOption {
	return func(opts *bidOptionImpl) {
		opts.debugBody = *debugBody
	}
}

type bidOptionImpl struct {
	dsp          string
	bidderUserId string
	contractID   string
	debugBody    bool
}

func (b *bidOptionImpl) Dsp() string          { return b.dsp }
func (b *bidOptionImpl) BidderUserId() string { return b.bidderUserId }
func (b *bidOptionImpl) ContractID() string   { return b.contractID }
func (b *bidOptionImpl) DebugBody() bool      { return b.debugBody }

func makeBidOptionImpl(opts ...BidOption) *bidOptionImpl {
	res := &bidOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeBidOptions(opts ...BidOption) BidOptions {
	return makeBidOptionImpl(opts...)
}
