package grpc

import (
	"context"
	"github.com/google/uuid"
	shelter "github.com/nourbol/zhandos/services/shelters/internal/delivery/grpc/interface"
	shelterDomain "github.com/nourbol/zhandos/services/shelters/internal/domain/shelter"
	shelterName "github.com/nourbol/zhandos/services/shelters/internal/domain/shelter/name"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (d *Delivery) CreateShelter(
	ctx context.Context,
	request *shelter.CreateShelterRequest) (*shelter.CreateShelterResponse, error) {
	name, err := shelterName.New(request.Name)
	if err != nil {
		return nil, err
	}

	var ownedBreeds []uuid.UUID

	for _, breedId := range request.OwnedBreeds {
		parsedBreedId, err := uuid.Parse(breedId)
		if err != nil {
			return nil, err
		}

		ownedBreeds = append(ownedBreeds, parsedBreedId)
	}

	createdShelter, err := shelterDomain.New(*name, ownedBreeds)
	if err != nil {
		return nil, err
	}

	insertedShelter, err := d.ucShelter.Insert(ctx, *createdShelter)
	if err != nil {
		return nil, err
	}

	var returnedBreeds []string

	for _, returnedBreedId := range insertedShelter.OwnedBreeds() {
		returnedBreeds = append(returnedBreeds, returnedBreedId.String())
	}

	return &shelter.CreateShelterResponse{Response: &shelter.ShelterResponse{
		Id:          insertedShelter.ID().String(),
		CreatedAt:   timestamppb.New(insertedShelter.CreatedAt()),
		Name:        insertedShelter.Name().String(),
		OwnedBreeds: returnedBreeds,
	}}, nil
}

//func (d *Delivery) UpdateShelter(context.Context, *UpdateShelterRequest) (*ShelterResponse, error)
//func (d *Delivery) DeleteShelter(context.Context, *DeleteShelterRequest) (*DeleteShelterResponse, error)
//func (d *Delivery) GetShelter(context.Context, *GetShelterRequest) (*GetShelterResponse, error)
//func (d *Delivery) GetShelters(context.Context, *GetSheltersRequest) (*GetSheltersResponse, error)
