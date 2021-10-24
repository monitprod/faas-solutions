package constant

type ContextKey int

const (
	IsLocal       ContextKey = iota
	LocalMainFunc ContextKey = iota
)
