package types

type MovieDetails struct {
	IMDbId       string `json:"imdbID"`
	Title        string `json:"Title"`
	IMDbRating   string `json:"imdbRating,omitempty"`
	ReleaseYear  string `json:"Year"`
	Plot_summary string `json:"Plot"`
	OK           string `json:"Response"`
	Error        string `json:"Error,omitempty"`
}
