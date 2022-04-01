// Code generated by gtrace. DO NOT EDIT.

package trace

// schemeComposeOptions is a holder of options
type schemeComposeOptions struct {
	panicCallback func(e interface{})
}

// SchemeOption specified Scheme compose option
type SchemeComposeOption func(o *schemeComposeOptions)

// WithSchemePanicCallback specified behavior on panic
func WithSchemePanicCallback(cb func(e interface{})) SchemeComposeOption {
	return func(o *schemeComposeOptions) {
		o.panicCallback = cb
	}
}

// Compose returns a new Scheme which has functional fields composed both from t and x.
func (t Scheme) Compose(x Scheme, opts ...SchemeComposeOption) (ret Scheme) {
	return ret
}
