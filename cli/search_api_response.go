package cli

type SearchAPIResponse struct {
	Results              []SearchResult `json:"Search"`
	AmoutOfResults       string         `json:"totalResults"`
	IsResponseSuccessful string         `json:"Response"`
	ErrorMessage         string         `json:"Error"`
}

type SearchResult struct {
	Title     string `json:"Title"`
	IMDBID    string `json:"imdbID"`
	Year      string `json:"Year"`
	Type      string `json:"Type"`
	PosterURL string `json:"Poster"`
}
