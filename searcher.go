package tinysearch

import (
	"fmt"
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

type phraseCursor struct {
	index     int
	positions []int
}

type phraseCursorSorter struct {
	cursors []phraseCursor
}

func (s phraseCursorSorter) Len() int      { return len(s.cursors) }
func (s phraseCursorSorter) Swap(i, j int) { s.cursors[i], s.cursors[j] = s.cursors[j], s.cursors[i] }
func (s phraseCursorSorter) Less(i, j int) bool {
	return len(s.cursors[i].positions) < len(s.cursors[j].positions)
}

func (s Searcher) Search(query string) (*SearchResult, error) {
	ret := SearchResult{
		Query: query,
	}
	tokens := s.analyzer.Tokenize(query)
	if len(tokens) == 0 {
		return &ret, nil
	}
	ByIndexCounts{invertedIndex: s.invertedIndex}.Sort(tokens)
	docCursors := make([]int, len(tokens)-1)
	basis := s.invertedIndex[tokens[0].Term]
	phraseCursors := make([]phraseCursor, len(tokens))

loop:
	for j := range basis {
		phraseCursors[0].index = tokens[0].ID
		phraseCursors[0].positions = basis[j].Positions
		for i, v := range tokens[1:] {
			pl := s.invertedIndex[v.Term]
			var ok bool
			for k := docCursors[i]; k < len(pl); k++ {
				if basis[j].DocID < pl[k].DocID {
					docCursors[i] = k
					continue loop
				} else if basis[j].DocID == pl[k].DocID {
					docCursors[i] = k + 1
					phraseCursors[i+1].index = v.ID
					phraseCursors[i+1].positions = pl[k].Positions
					ok = true
					break
				}
			}
			if !ok {
				break loop
			}
		}

		if ok := s.phraseCheck(phraseCursors); ok {
			ret.Docs = append(ret.Docs, ScoredDocID{
				DocID: basis[j].DocID, // score はまだ決まらない
			})
		}
	}
	return &ret, nil
}

func (s Searcher) phraseCheck(ps []phraseCursor) bool {
	sort.Sort(&phraseCursorSorter{cursors: ps})
	cursors := make([]int, len(ps))
	var ret bool
loop:
	for n := 0; n < len(ps[0].positions); n++ {
		x := ps[0].positions[n] - ps[0].index
		for i := 1; i < len(ps); i++ {
			ph := ps[i]
			var ok bool
			for j := cursors[i]; j < len(ph.positions); j++ {
				v := ph.positions[j] - ph.index
				if x < v {
					cursors[j] = j
					continue loop
				} else if v == x {
					cursors[i] = j + 1
					ok = true
					break
				}
			}
			if !ok {
				continue loop
			}
		}
		ret = true
		fmt.Printf("phrase detect!, pos=%v\n", x)
	}

	return ret
}
