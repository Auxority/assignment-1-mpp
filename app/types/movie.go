package types

type Movie struct {
	IMDbId      *string  `json:"imdb_id"`
	Title       *string  `json:"title"`
	IMDbRating  *float64 `json:"rating"`
	ReleaseYear *int     `json:"year"`
}
