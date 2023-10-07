package types

import (
	_ "github.com/go-playground/validator/v10"
)

type Movie struct {
	IMDbId       *string  `json:"imdb_id" validate:"required,min=9,max=9"`
	Title        *string  `json:"title" validate:"required,min=1,max=255"`
	IMDbRating   *float64 `json:"rating,string" validate:"required,min=1,max=10"`
	ReleaseYear  *int     `json:"year,string" validate:"required,min=1,max=9999"`
	Plot_summary *string  `json:"plot,omitempty" validate:"omitempty,min=1,max=255"`
}
