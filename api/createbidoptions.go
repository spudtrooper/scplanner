// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type CreateBidOption func(*createBidOptionImpl)

type CreateBidOptions interface {
	Dsp() string
	HasDsp() bool
	BidderUserId() string
	HasBidderUserId() bool
	ContractID() string
	HasContractID() bool
	DebugBody() bool
	HasDebugBody() bool
}

func CreateBidDsp(dsp string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.has_dsp = true
		opts.dsp = dsp
	}
}
func CreateBidDspFlag(dsp *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		if dsp == nil {
			return
		}
		opts.has_dsp = true
		opts.dsp = *dsp
	}
}

func CreateBidBidderUserId(bidderUserId string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.has_bidderUserId = true
		opts.bidderUserId = bidderUserId
	}
}
func CreateBidBidderUserIdFlag(bidderUserId *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		if bidderUserId == nil {
			return
		}
		opts.has_bidderUserId = true
		opts.bidderUserId = *bidderUserId
	}
}

func CreateBidContractID(contractID string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.has_contractID = true
		opts.contractID = contractID
	}
}
func CreateBidContractIDFlag(contractID *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		if contractID == nil {
			return
		}
		opts.has_contractID = true
		opts.contractID = *contractID
	}
}

func CreateBidDebugBody(debugBody bool) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func CreateBidDebugBodyFlag(debugBody *bool) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

type createBidOptionImpl struct {
	dsp              string
	has_dsp          bool
	bidderUserId     string
	has_bidderUserId bool
	contractID       string
	has_contractID   bool
	debugBody        bool
	has_debugBody    bool
}

func (c *createBidOptionImpl) Dsp() string           { return c.dsp }
func (c *createBidOptionImpl) HasDsp() bool          { return c.has_dsp }
func (c *createBidOptionImpl) BidderUserId() string  { return c.bidderUserId }
func (c *createBidOptionImpl) HasBidderUserId() bool { return c.has_bidderUserId }
func (c *createBidOptionImpl) ContractID() string    { return c.contractID }
func (c *createBidOptionImpl) HasContractID() bool   { return c.has_contractID }
func (c *createBidOptionImpl) DebugBody() bool       { return c.debugBody }
func (c *createBidOptionImpl) HasDebugBody() bool    { return c.has_debugBody }

func makeCreateBidOptionImpl(opts ...CreateBidOption) *createBidOptionImpl {
	res := &createBidOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeCreateBidOptions(opts ...CreateBidOption) CreateBidOptions {
	return makeCreateBidOptionImpl(opts...)
}
