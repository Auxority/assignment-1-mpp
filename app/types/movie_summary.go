package types

type MovieSummary struct {
	OK     string `json:"Response"`
	Error  string `json:"Error"`
	Plot   string `json:"Plot"`
	IMDbId string `json:"imdbID"`
}
