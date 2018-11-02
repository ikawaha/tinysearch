package tinysearch

import (
	"fmt"
)

type Storage struct {
	Directory string
	FileName  string
	Encoder   Encoder
}

func NewJSONStorage(dir, file string) Storage {
	return Storage{
		Directory: dir,
		FileName:  file,
		Encoder:   JSONEncoder{},
	}
}

func (s Storage) Persist(ii InvertedIndex) error {
	b, err := s.Encoder.Encode(ii)
	if err != nil {
		return err
	}
	fmt.Printf("%s", b) //TODO output to the file
	return nil
}
