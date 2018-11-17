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

func (s Searcher) Search(query string) (SearchResult, error) {
	ret := SearchResult{
		Query: query,
	}
	tokens := s.analyzer.Tokenize(query)
	if len(tokens) == 0 {
		return ret, nil
	}
	ByIndexCounts{invertedIndex: s.invertedIndex}.Sort(tokens)
	cursors := make([]int, len(tokens)-1)
	basis := s.invertedIndex[tokens[0].Term]
loop:
	for j := range basis {
		for i, v := range tokens[1:] {
			pl := s.invertedIndex[v.Term]
			var ok bool
			for k := cursors[i]; k < len(pl); k++ {
				if basis[j].DocID < pl[k].DocID {
					cursors[i] = k
					continue loop
				} else if basis[j].DocID == pl[k].DocID {
					cursors[i] = k + 1
					ok = true
					break
				}
			}
			if !ok {
				break loop
			}
		}
		ret.Docs = append(ret.Docs, ScoredDocID{DocID: basis[j].DocID}) //XXX スコア未実装
	}
	return ret, nil
}
