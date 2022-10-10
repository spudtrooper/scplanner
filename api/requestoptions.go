// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

type RequestOption func(*requestOptionImpl)

type RequestOptions interface {
	ExtraHeaders() map[string]string
	HasExtraHeaders() bool
	Host() string
	HasHost() bool
	CustomPayload() interface{}
	HasCustomPayload() bool
}

func RequestExtraHeaders(extraHeaders map[string]string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.has_extraHeaders = true
		opts.extraHeaders = extraHeaders
	}
}
func RequestExtraHeadersFlag(extraHeaders *map[string]string) RequestOption {
	return func(opts *requestOptionImpl) {
		if extraHeaders == nil {
			return
		}
		opts.has_extraHeaders = true
		opts.extraHeaders = *extraHeaders
	}
}

func RequestHost(host string) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.has_host = true
		opts.host = host
	}
}
func RequestHostFlag(host *string) RequestOption {
	return func(opts *requestOptionImpl) {
		if host == nil {
			return
		}
		opts.has_host = true
		opts.host = *host
	}
}

func RequestCustomPayload(customPayload interface{}) RequestOption {
	return func(opts *requestOptionImpl) {
		opts.has_customPayload = true
		opts.customPayload = customPayload
	}
}
func RequestCustomPayloadFlag(customPayload *interface{}) RequestOption {
	return func(opts *requestOptionImpl) {
		if customPayload == nil {
			return
		}
		opts.has_customPayload = true
		opts.customPayload = *customPayload
	}
}

type requestOptionImpl struct {
	extraHeaders      map[string]string
	has_extraHeaders  bool
	host              string
	has_host          bool
	customPayload     interface{}
	has_customPayload bool
}

func (r *requestOptionImpl) ExtraHeaders() map[string]string { return r.extraHeaders }
func (r *requestOptionImpl) HasExtraHeaders() bool           { return r.has_extraHeaders }
func (r *requestOptionImpl) Host() string                    { return r.host }
func (r *requestOptionImpl) HasHost() bool                   { return r.has_host }
func (r *requestOptionImpl) CustomPayload() interface{}      { return r.customPayload }
func (r *requestOptionImpl) HasCustomPayload() bool          { return r.has_customPayload }

func makeRequestOptionImpl(opts ...RequestOption) *requestOptionImpl {
	res := &requestOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRequestOptions(opts ...RequestOption) RequestOptions {
	return makeRequestOptionImpl(opts...)
}
