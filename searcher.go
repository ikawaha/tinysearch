package tinysearch

import (
	"sort"
)

type ScoredDocID struct {
	DocID int
	Score int // 同じドキュメントで何回出てくるか
}

type SearchResult struct {
	Query string
	Docs  []ScoredDocID
}

type Searcher struct {
	analyzer      Analyzer
	invertedIndex InvertedIndex
}

type tokenSorter struct {
	tokens []Token
	by     ByIndexCounts
}

type ByIndexCounts struct {
	invertedIndex InvertedIndex
}

func (by ByIndexCounts) Sort(tokens []Token) {
	ts := tokenSorter{
		tokens: tokens,
		by:     by,
	}
	sort.Sort(&ts)
}

func (s tokenSorter) Len() int      { return len(s.tokens) }
func (s tokenSorter) Swap(i, j int) { s.tokens[i], s.tokens[j] = s.tokens[j], s.tokens[i] }
func (s tokenSorter) Less(i, j int) bool {
	a := s.by.invertedIndex[s.tokens[i].Term]
	b := s.by.invertedIndex[s.tokens[j].Term]
	return len(a) < len(b)
}

func (s Searcher) Search(query string) (SearchResult, error) {
	var ret SearchResult
	ret.Query = query

	tokens := s.analyzer.Tokenize(query)
	ByIndexCounts{invertedIndex: s.invertedIndex}.Sort(tokens)

	cursors := map[string]int{}

	basis := s.invertedIndex[tokens[0].Term] //基準となるPostingList
loop:
	for j := range basis { // 最初の token のポスティングリストを基準に
		for _, v := range tokens[1:] { // 残りの token のポスティングリストを見ていく
			pl := s.invertedIndex[v.Term]
			var ok bool
			for k := cursors[v.Term]; k < len(pl); k++ { // カーソルを持っておくと begin を決められる
				if basis[j].DocID < pl[k].DocID {
					cursors[v.Term] = k
					continue loop
				} else if basis[j].DocID == pl[k].DocID {
					cursors[v.Term] = k + 1
					ok = true
					break
				}
			}
			if !ok { // ok でないときは候補がなくなってしまったとき
				break loop
			}
		}
		ret.Docs = append(ret.Docs, ScoredDocID{DocID: basis[j].DocID}) //XXX スコア未実装
	}
	return ret, nil
}
