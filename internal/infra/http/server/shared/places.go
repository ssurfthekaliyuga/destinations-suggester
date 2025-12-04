package shared

import (
	"destinations-suggester/internal/domain/models/places"
	"github.com/google/uuid"
)

type Place struct {
	ID          uuid.UUID   `json:"id"`
	FIAS        uuid.UUID   `json:"fias"`
	Coordinates Coordinates `json:"Coordinates"`
}

func (Place) FromModel(place places.Place) Place {
	return Place{
		ID:          place.ID,
		FIAS:        place.FIAS,
		Coordinates: Coordinates(place.Coordinates),
	}
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
