package tinysearch

import (
	"sort"
)

type Token struct {
	Term  string
	Start int
	End   int
}

type StringComparator interface {
	Less(string, string) bool
}

type tokenSorter struct {
	tokens []Token
	by     StringComparator
}

func (s tokenSorter) Len() int      { return len(s.tokens) }
func (s tokenSorter) Swap(i, j int) { s.tokens[i], s.tokens[j] = s.tokens[j], s.tokens[i] }
func (s tokenSorter) Less(i, j int) bool {
	return s.by.Less(s.tokens[i].Term, s.tokens[j].Term)
}

type ByIndexCounts struct {
	invertedIndex InvertedIndex
}

func (by ByIndexCounts) Less(a, b string) bool {
	return len(by.invertedIndex[a]) < len(by.invertedIndex[b])
}

func (by ByIndexCounts) Sort(tokens []Token) {
	ts := tokenSorter{
		tokens: tokens,
		by:     by,
	}
	sort.Sort(&ts)
}
