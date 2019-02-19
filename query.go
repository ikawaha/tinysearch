package tinysearch

type QueryType int

const (
	Default QueryType = iota
	Phrase
	Fuzzy
)

type Query struct {
	Raw       string
	QueryType QueryType
}
