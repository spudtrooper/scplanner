// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type ResolveOption func(*resolveOptionImpl)

type ResolveOptions interface {
}

type resolveOptionImpl struct {
}

type ResolveParams struct {
	Url string `json:"url" required:"true"`
}

func (o ResolveParams) Options() []ResolveOption {
	return []ResolveOption{}
}

func makeResolveOptionImpl(opts ...ResolveOption) *resolveOptionImpl {
	res := &resolveOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeResolveOptions(opts ...ResolveOption) ResolveOptions {
	return makeResolveOptionImpl(opts...)
}
