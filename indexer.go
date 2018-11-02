package tinysearch

import "io"

type Indexer struct {
	InternalID int
	Analyzer   Analyzer
	Index      InvertedIndex
	Storage    Storage
}

func NewIndexer(a Analyzer, s Storage) *Indexer {
	return &Indexer{
		Analyzer: a,
		Index:    InvertedIndex{},
		Storage:  s,
	}
}

//TODO id が被る場合
func (ii *Indexer) AddDocument(id int, d Document) {
	m := map[string]*Posting{}
	ts := ii.Analyzer.Tokenize(d.Text)
	for i, v := range ts {
		p, ok := m[v.Term]
		if !ok {
			p = &Posting{DocID: d.ID, Positions: []int{}}
			m[v.Term] = p
		}
		p.Positions = append(p.Positions, i) // タームの位置
	}
	for k, v := range m {
		if v == nil {
			continue
		}
		ii.Index[k] = append(ii.Index[k], *v)
	}
}

func (ii Indexer) WriteTo(r io.Writer) error {
	//TODO
	ii.Storage.Persist(ii.Index)
	return nil
}
