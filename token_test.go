package tinysearch

import (
	"reflect"
	"testing"
)

func TestByIndexCounts_Sort(t *testing.T) {
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
	tokens := NewNgramAnalyzer(2).Tokenize("きょうは")
	ByIndexCounts{invertedIndex: ii}.Sort(tokens)
	expected := []Token{
		{Term: "うは", Start: 6, End: 12},
		{Term: "ょう", Start: 3, End: 9},
		{Term: "きょ", Start: 0, End: 6},
	}
	if !reflect.DeepEqual(expected, tokens) {
		t.Errorf("expected %+v, got %+v", expected, tokens)
	}
}
