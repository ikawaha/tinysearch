package tinysearch

type Posting struct {
	DocID     int
	Positions []int
}

type Postings []Posting
