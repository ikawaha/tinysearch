package tinysearch

import (
	"sort"
)

type ScoredDocID {
	docID int
	score int
}

type SearchResult struct {
	query string
	docs []ScoredDocID
}

type Seacher struct {
	analyzer  Analyzer
	invertedIndex InvertedIndex
}


type tokenSorter struct {
	tokens []Token
	by ByIndexCounts
}

type ByIndexCounts struct{
	invertedIndex InvertedIndex
}

func (by ByIndexCounts) Sort(tokens []Token) {
	ts := tokenSorter{
		tokens: tokens,
		by: by,
	}
	sort.Sort(&ts)
}

func (s tokenSorter) Len() int {return len(s.tokens)}
func (s tokenSorter) Swap(i, j int) {s.tokens[i], s.tokens[j] = s.tokens[j], s.tokens[i]}
func (s tokenSorter) Less(i, j int) bool {
	a := s.by.invertedIndex[s.tokens[i].Term]
	b := s.by.invertedIndex[s.tokens[j].Term]
	return len(a) < len(b)
}

func (s Seacher) Search(query string) (SearchResult, error) {
	var ret SearchResult
	tokens := s.analyzer.Tokenize(query)
	ByIndexCounts{invertedIndex:s.invertedIndex}.Sort(tokens)

	cursors := map[string]int{}

	var docIDs []int
	a := s.invertedIndex[tokens[0].Term]

loop:
	for j := range a {
		for i := range tokens[1:]{
			b := s.invertedIndex[tokens[i].Term]
			begin := cursors[tokens[i].Term]
			for k:= begin; k <len(b); k++ { // カーソルを持っておくと begin を決められる
				if a[j].DocID >b[k].DocID {
					continue
				}
			}
			// 失敗したら
			goto loop
		}
		docIDs = append(docIDs, a[j])
	}
}
