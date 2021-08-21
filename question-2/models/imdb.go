package models

// QueryParams is the type to create the query parameters
type QueryParams struct {
	Search string
	Page   string
}

type ListResults struct {
	Title  string
	Year   string
	ImdbID string
	Type   string
	Poster string
}

type MovieDetail struct {
	Title    string
	Year     string
	ImdbID   string
	Type     string
	Poster   string
	Response string
	Error    string
}

type SearchResponse struct {
	Search       []ListResults
	Response     string
	Error        string
	totalResults int
}
