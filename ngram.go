package tinysearch

import (
	"unicode/utf8"
)

type NgramAnalyzer struct {
	N int
}

func NewNgramAnalyzer(n int) *NgramAnalyzer {
	return &NgramAnalyzer{N: n}
}

func (a NgramAnalyzer) Tokenize(s string) []Token {
	if a.N < 1 {
		return nil
	}
	var ret []Token
L:
	for i := 0; i < len(s); {
		l, l0 := 0, 0
		for k := 0; k < a.N; k++ {
			r, size := utf8.DecodeRuneInString(s[i+l:])
			if k == 0 {
				l0 = size
			}
			if l += size; r == utf8.RuneError || l > len(s) {
				break L
			}
		}
		ret = append(ret, Token{ID: len(ret), Term: s[i : i+l], Start: i, End: i + l})
		i += l0
	}
	return ret
}
