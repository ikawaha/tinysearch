package tinysearch

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearcher_Search(t *testing.T) {
	t.Run("図4-1", func(t *testing.T) {
		ii := InvertedIndex{
			"きょ": {
				{DocID: 15, Positions: []int{4, 36, 100}},
				{DocID: 18, Positions: []int{15, 30}},
				{DocID: 30, Positions: []int{457}},
				{DocID: 87, Positions: []int{76, 543}},
				{DocID: 213, Positions: []int{43, 68}},
			},
			"ょう": {
				{DocID: 13, Positions: []int{10}},
				{DocID: 17, Positions: []int{65}},
				{DocID: 18, Positions: []int{8, 31}},
				{DocID: 114, Positions: []int{4, 67, 117}},
			},
			"うは": {
				{DocID: 1, Positions: []int{0, 2}},
				{DocID: 18, Positions: []int{4, 32}},
				{DocID: 196, Positions: []int{5}},
			},
		}

		s := Searcher{
			analyzer:      NewNgramAnalyzer(2),
			invertedIndex: ii,
		}
		docs, err := s.Search("きょうは")
		fmt.Println(docs, err)
	})
	t.Run("基準より小さいものしかないとき", func(t *testing.T) {
		ii := InvertedIndex{
			"きょ": {
				{DocID: 15, Positions: []int{4, 36, 100}},
				{DocID: 16, Positions: []int{15, 30}},
				{DocID: 17, Positions: []int{457}},
			},
			"ょう": {
				{DocID: 3, Positions: []int{10}},
				{DocID: 13, Positions: []int{10}},
				{DocID: 17, Positions: []int{65}},
			},
			"うは": {
				{DocID: 18, Positions: []int{4, 32}},
			},
		}

		s := Searcher{
			analyzer:      NewNgramAnalyzer(2),
			invertedIndex: ii,
		}
		docs, err := s.Search("きょうは")
		fmt.Println(docs, err)
	})
	t.Run("どれも被らないとき", func(t *testing.T) {
		ii := InvertedIndex{
			"きょ": {
				{DocID: 4, Positions: []int{4, 36, 100}},
				{DocID: 5, Positions: []int{15, 30}},
				{DocID: 6, Positions: []int{457}},
			},
			"ょう": {
				{DocID: 2, Positions: []int{10}},
				{DocID: 3, Positions: []int{65}},
			},
			"うは": {
				{DocID: 1, Positions: []int{4, 32}},
			},
		}

		s := Searcher{
			analyzer:      NewNgramAnalyzer(2),
			invertedIndex: ii,
		}
		docs, err := s.Search("きょうは")
		fmt.Println(docs, err)
	})

}
