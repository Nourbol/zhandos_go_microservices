package useCase

import (
	"context"
	"github.com/google/uuid"
	_filters "github.com/nourbol/zhandos/pkg/tools/filters"
	_shelter "github.com/nourbol/zhandos/services/shelters/internal/domain/shelter"
)

type Shelter interface {
	Insert(ctx context.Context, shelter _shelter.Shelter) (*_shelter.Shelter, error)
	Update(ctx context.Context, shelter _shelter.Shelter) error
	Delete(ctx context.Context, id uuid.UUID) error

	ShelterReader
}

type ShelterReader interface {
	Get(ctx context.Context, id uuid.UUID) (*_shelter.Shelter, error)
	GetAll(
		ctx context.Context,
		name string,
		ownedBreeds []uuid.UUID,
		filters _filters.Filters) ([]*_shelter.Shelter, _filters.Metadata, error)
}
