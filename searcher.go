package tinysearch

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
			for k := docCursors[i] + 1; k < len(pl); k++ {
				if basis[j].DocID < pl[k].DocID {
					docCursors[i] = k - 1
					continue loop
				} else if basis[j].DocID == pl[k].DocID {
					docCursors[i] = k
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

		if score, ok := s.phraseCheck(phraseCursors); ok {
			ret.Docs = append(ret.Docs, ScoredDocID{
				DocID: basis[j].DocID,
				Score: score,
			})
		}
	}
	return &ret, nil
}

func (s Searcher) phraseCheck(ps []phraseCursor) (int, bool) {
	return 0, true
}
