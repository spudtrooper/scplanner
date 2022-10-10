// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type BidOption func(*bidOptionImpl)

type BidOptions interface {
	Dsp() string
	HasDsp() bool
	BidderUserId() string
	HasBidderUserId() bool
	ContractID() string
	HasContractID() bool
	DebugBody() bool
	HasDebugBody() bool
}

func BidDsp(dsp string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.has_dsp = true
		opts.dsp = dsp
	}
}
func BidDspFlag(dsp *string) BidOption {
	return func(opts *bidOptionImpl) {
		if dsp == nil {
			return
		}
		opts.has_dsp = true
		opts.dsp = *dsp
	}
}

func BidBidderUserId(bidderUserId string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.has_bidderUserId = true
		opts.bidderUserId = bidderUserId
	}
}
func BidBidderUserIdFlag(bidderUserId *string) BidOption {
	return func(opts *bidOptionImpl) {
		if bidderUserId == nil {
			return
		}
		opts.has_bidderUserId = true
		opts.bidderUserId = *bidderUserId
	}
}

func BidContractID(contractID string) BidOption {
	return func(opts *bidOptionImpl) {
		opts.has_contractID = true
		opts.contractID = contractID
	}
}
func BidContractIDFlag(contractID *string) BidOption {
	return func(opts *bidOptionImpl) {
		if contractID == nil {
			return
		}
		opts.has_contractID = true
		opts.contractID = *contractID
	}
}

func BidDebugBody(debugBody bool) BidOption {
	return func(opts *bidOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func BidDebugBodyFlag(debugBody *bool) BidOption {
	return func(opts *bidOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

type bidOptionImpl struct {
	dsp              string
	has_dsp          bool
	bidderUserId     string
	has_bidderUserId bool
	contractID       string
	has_contractID   bool
	debugBody        bool
	has_debugBody    bool
}

func (b *bidOptionImpl) Dsp() string           { return b.dsp }
func (b *bidOptionImpl) HasDsp() bool          { return b.has_dsp }
func (b *bidOptionImpl) BidderUserId() string  { return b.bidderUserId }
func (b *bidOptionImpl) HasBidderUserId() bool { return b.has_bidderUserId }
func (b *bidOptionImpl) ContractID() string    { return b.contractID }
func (b *bidOptionImpl) HasContractID() bool   { return b.has_contractID }
func (b *bidOptionImpl) DebugBody() bool       { return b.debugBody }
func (b *bidOptionImpl) HasDebugBody() bool    { return b.has_debugBody }

type BidParams struct {
	ContractId   string     `json:"contract_id" required:"true"`
	Auth         AuthInfo   `json:"auth" required:"true"`
	BidMedia     TargetInfo `json:"bid_media" required:"true"`
	Dsp          string     `json:"dsp"`
	BidderUserId string     `json:"bidder_user_id"`
	ContractID   string     `json:"contract_id"`
	DebugBody    bool       `json:"debug_body"`
}

func (o BidParams) Options() []BidOption {
	return []BidOption{
		BidDsp(o.Dsp),
		BidBidderUserId(o.BidderUserId),
		BidContractID(o.ContractID),
		BidDebugBody(o.DebugBody),
	}
}

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
