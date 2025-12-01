package places

import (
	"github.com/google/uuid"
)

type Place struct {
	ID          uuid.UUID
	FIAS        uuid.UUID
	Coordinates Coordinates
}
