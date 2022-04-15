package persqueue

// CreateStreamOption stores additional options for stream modification requests
type StreamOption func() // FIXME: сделать накат на internal описание топиков

// WithReadRule add read rule setting to stream modification calls
func WithReadRule(ReadRule) StreamOption {
	return nil
}

// WithRemoteMirrorRule add remote mirror settings to stream modification calls
func WithRemoteMirrorRule(RemoteMirrorRule) StreamOption {
	return nil
}
