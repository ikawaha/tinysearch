package tinysearch

type Encoder interface {
	Encode(ii InvertedIndex) ([]byte, error)
}
