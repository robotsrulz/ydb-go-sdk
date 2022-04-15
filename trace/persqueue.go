package trace

// tool gtrace used from ./internal/cmd/gtrace

//go:generate gtrace

type (
	// Persqueue specified trace of persqueue client activity.
	// gtrace:gen
	Persqueue struct{}
)
