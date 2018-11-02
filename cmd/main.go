package main

import (
	"github.com/ikawaha/tinysearch"
)

func main() {

	store := tinysearch.NewJSONStorage("/tmp", "index.json")
	analyzer := tinysearch.NewNgramAnalyzer(2)
	indexer := tinysearch.NewIndexer(analyzer, store)

	docs := []string{
		"Doc1!", "Doc2!", "ドキュメント3",
	}

	for i, d := range docs {
		indexer.AddDocument(i, tinysearch.Document{ID: i, Text: d})
	}
	indexer.WriteTo(nil)
}
