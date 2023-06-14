package _shelter

import (
	"github.com/google/uuid"
	"github.com/nourbol/zhandos/services/shelters/internal/domain/shelter/name"
	"github.com/pkg/errors"
	"time"
)

var (
	ErrNameRequired = errors.New("name is required")
)

type Shelter struct {
	id        uuid.UUID
	createdAt time.Time

	name name.Name

	ownedBreeds []uuid.UUID
}

func NewWithId(
	id uuid.UUID,
	createdAt time.Time,
	name name.Name,
	ownedBreeds []uuid.UUID) (*Shelter, error) {
	if name.IsEmpty() {
		return nil, ErrNameRequired
	}

	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Shelter{
		id:          id,
		createdAt:   createdAt.UTC(),
		name:        name,
		ownedBreeds: ownedBreeds,
	}, nil
}

func New(name name.Name, ownedBreeds []uuid.UUID) (*Shelter, error) {
	if name.IsEmpty() {
		return nil, ErrNameRequired
	}

	return &Shelter{
		id:          uuid.New(),
		createdAt:   time.Now(),
		name:        name,
		ownedBreeds: ownedBreeds,
	}, nil
}

func (s Shelter) ID() uuid.UUID {
	return s.id
}

func (s Shelter) CreatedAt() time.Time {
	return s.createdAt
}

func (s Shelter) Name() name.Name {
	return s.name
}

func (s Shelter) OwnedBreeds() []uuid.UUID {
	return s.ownedBreeds
}

func (s Shelter) Equals(shelter Shelter) bool {
	return s.id == shelter.id
}
