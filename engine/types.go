package engine

// the struct are saving *citylist.go* matches city-url
type Request struct {
	Url string
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request // *Requests* typing is Request of slice
	Items []interface{}
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}