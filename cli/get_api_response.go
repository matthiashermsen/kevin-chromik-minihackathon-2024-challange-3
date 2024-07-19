package cli

type GetAPIResponse struct {
	Title                string `json:"Title"`
	Year                 string `json:"Year"`
	Released             string `json:"Released"`
	Runtime              string `json:"Runtime"`
	Genre                string `json:"Genre"`
	Director             string `json:"Director"`
	Writer               string `json:"Writer"`
	Actors               string `json:"Actors"`
	Plot                 string `json:"Plot"`
	Language             string `json:"Language"`
	Country              string `json:"Country"`
	PosterURL            string `json:"Poster"`
	Metascore            string `json:"Metascore"`
	IMDBRating           string `json:"imdbRating"`
	IMDBVotes            string `json:"imdbVotes"`
	Type                 string `json:"type"`
	AmountOfSeasons      string `json:"totalSeasons"`
	IsResponseSuccessful string `json:"Response"`
	ErrorMessage         string `json:"Error"`
}
