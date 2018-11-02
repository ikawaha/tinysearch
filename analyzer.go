package tinysearch

type Analyzer interface {
	Tokenize(s string) []Token
}
