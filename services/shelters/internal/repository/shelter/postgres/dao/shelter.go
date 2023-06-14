package dao

import (
	"github.com/google/uuid"
	"github.com/nourbol/zhandos/services/shelters/internal/domain/shelter"
	"github.com/nourbol/zhandos/services/shelters/internal/domain/shelter/name"
	"time"
)

type Shelter struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	Name        string
	OwnedBreeds []uuid.UUID
}

func (s *Shelter) ToDomainShelter() (*_shelter.Shelter, error) {
	shelterName, err := name.New(s.Name)
	if err != nil {
		return nil, err
	}

	return _shelter.NewWithId(
		s.ID,
		s.CreatedAt,
		*shelterName,
		s.OwnedBreeds)
}
