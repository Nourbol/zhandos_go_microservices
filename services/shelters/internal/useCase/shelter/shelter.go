package shelter

import (
	"context"
	"github.com/google/uuid"
	_filters "github.com/nourbol/zhandos/pkg/tools/filters"
	_shelter "github.com/nourbol/zhandos/services/shelters/internal/domain/shelter"
)

func (uc *UseCase) Insert(ctx context.Context, shelter _shelter.Shelter) (*_shelter.Shelter, error) {
	return uc.adapterStorage.Insert(ctx, shelter)
}

func (uc *UseCase) Update(ctx context.Context, shelter _shelter.Shelter) error {
	return uc.adapterStorage.Update(ctx, shelter)
}

func (uc *UseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.adapterStorage.Delete(ctx, id)
}

func (uc *UseCase) GetAll(
	ctx context.Context,
	name string,
	ownedBreeds []uuid.UUID,
	filters _filters.Filters) ([]*_shelter.Shelter, _filters.Metadata, error) {
	return uc.adapterStorage.GetAll(ctx, name, ownedBreeds, filters)
}

func (uc *UseCase) Get(ctx context.Context, id uuid.UUID) (*_shelter.Shelter, error) {
	return uc.adapterStorage.Get(ctx, id)
}
