package persqueue

// CreateStreamOption stores additional options for stream modification requests
type StreamOption func(streamOptions)

// WithReadRule add read rule setting to stream modification calls
func WithReadRule(r ReadRule) StreamOption {
	return func(o streamOptions) {
		o.AddReadRule(r)
	}
}

// WithRemoteMirrorRule add remote mirror settings to stream modification calls
func WithRemoteMirrorRule(r RemoteMirrorRule) StreamOption {
	return func(o streamOptions) {
		o.SetRemoteMirrorRule(r)
	}
}

type streamOptions interface {
	AddReadRule(r ReadRule)
	SetRemoteMirrorRule(r RemoteMirrorRule)
}

type StreamingWriteOption func()

type steramingWriteOption interface {
	SetCodec(Codec)
	SetFormat(Format)
	// Block encoding settings
}
