// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type MakeClientOption func(*makeClientOptionImpl)

type MakeClientOptions interface {
	Debug() bool
	HasDebug() bool
}

func MakeClientDebug(debug bool) MakeClientOption {
	return func(opts *makeClientOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}
}
func MakeClientDebugFlag(debug *bool) MakeClientOption {
	return func(opts *makeClientOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}
}

type makeClientOptionImpl struct {
	debug     bool
	has_debug bool
}

func (m *makeClientOptionImpl) Debug() bool    { return m.debug }
func (m *makeClientOptionImpl) HasDebug() bool { return m.has_debug }

func makeMakeClientOptionImpl(opts ...MakeClientOption) *makeClientOptionImpl {
	res := &makeClientOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeMakeClientOptions(opts ...MakeClientOption) MakeClientOptions {
	return makeMakeClientOptionImpl(opts...)
}
