package api

//go:generate genopts --prefix=CreateBid --outfile=api/createbidoptions.go "dsp:string" "bidderUserId:string" "contractID:string" "debugBody"

type CreateBidOption func(*createBidOptionImpl)

type CreateBidOptions interface {
	Dsp() string
	BidderUserId() string
	ContractID() string
	DebugBody() bool
}

func CreateBidDsp(dsp string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.dsp = dsp
	}
}
func CreateBidDspFlag(dsp *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.dsp = *dsp
	}
}

func CreateBidBidderUserId(bidderUserId string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.bidderUserId = bidderUserId
	}
}
func CreateBidBidderUserIdFlag(bidderUserId *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.bidderUserId = *bidderUserId
	}
}

func CreateBidContractID(contractID string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.contractID = contractID
	}
}
func CreateBidContractIDFlag(contractID *string) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.contractID = *contractID
	}
}

func CreateBidDebugBody(debugBody bool) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.debugBody = debugBody
	}
}
func CreateBidDebugBodyFlag(debugBody *bool) CreateBidOption {
	return func(opts *createBidOptionImpl) {
		opts.debugBody = *debugBody
	}
}

type createBidOptionImpl struct {
	dsp          string
	bidderUserId string
	contractID   string
	debugBody    bool
}

func (c *createBidOptionImpl) Dsp() string          { return c.dsp }
func (c *createBidOptionImpl) BidderUserId() string { return c.bidderUserId }
func (c *createBidOptionImpl) ContractID() string   { return c.contractID }
func (c *createBidOptionImpl) DebugBody() bool      { return c.debugBody }

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
