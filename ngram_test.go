package tinysearch

import (
	"testing"
)

func TestNgram(t *testing.T) {
	testdata := []struct {
		N        int
		Input    string
		Expected []Token
	}{
		{
			N:     1,
			Input: "すもも",
			Expected: []Token{
				{Term: "す", Start: 0, End: 3},
				{Term: "も", Start: 3, End: 6},
				{Term: "も", Start: 6, End: 9},
			},
		},
		{
			N:     2,
			Input: "すもも",
			Expected: []Token{
				{Term: "すも", Start: 0, End: 6},
				{Term: "もも", Start: 3, End: 9},
			},
		},
		{
			N:     2,
			Input: "aあbいc",
			Expected: []Token{
				{Term: "aあ", Start: 0, End: 4},
				{Term: "あb", Start: 1, End: 5},
				{Term: "bい", Start: 4, End: 8},
				{Term: "いc", Start: 5, End: 9},
			},
		},
		{
			N:     3,
			Input: "aあbいc",
			Expected: []Token{
				{Term: "aあb", Start: 0, End: 5},
				{Term: "あbい", Start: 1, End: 8},
				{Term: "bいc", Start: 4, End: 9},
			},
		},
		{
			N:        10,
			Input:    "aあbいc",
			Expected: []Token{},
		},
		{
			N:        0,
			Input:    "aあbいc",
			Expected: []Token{},
		},
	}
	for i, v := range testdata {
		a := NewNgramAnalyzer(v.N)
		tokens := a.Tokenize(v.Input)
		if l, expected := len(tokens), len(v.Expected); l != expected {
			t.Errorf("%d, got %+v, expected %+v", i, tokens, v.Expected)
			break
		}
		for i := range tokens {
			if tokens[i] != v.Expected[i] {
				t.Errorf("%d, got %+v, expected %+v", i, tokens[i], v.Expected[i])
				break
			}
		}
	}
}
