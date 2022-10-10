// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type AuthOption func(*authOptionImpl)

type AuthOptions interface {
}

type authOptionImpl struct {
}

type AuthParams struct {
}

func (o AuthParams) Options() []AuthOption {
	return []AuthOption{}
}

func makeAuthOptionImpl(opts ...AuthOption) *authOptionImpl {
	res := &authOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAuthOptions(opts ...AuthOption) AuthOptions {
	return makeAuthOptionImpl(opts...)
}
