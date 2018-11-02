package tinysearch

import (
	"encoding/json"
)

type JSONEncoder struct{}

func (e JSONEncoder) Encode(ii InvertedIndex) ([]byte, error) {
	return json.MarshalIndent(ii, "", "\t")
}
