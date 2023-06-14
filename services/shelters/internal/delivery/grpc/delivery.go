package grpc

import (
	jsonlog "github.com/nourbol/zhandos/pkg/logger"
	shelter "github.com/nourbol/zhandos/services/shelters/internal/delivery/grpc/interface"
	"github.com/nourbol/zhandos/services/shelters/internal/useCase"
)

type Options struct {
	port int
	env  string
}

type Delivery struct {
	shelter.UnimplementedShelterServiceServer

	logger    *jsonlog.Logger
	ucShelter useCase.Shelter

	options Options
}

func New(logger *jsonlog.Logger, ucShelter useCase.Shelter, options Options) *Delivery {
	return &Delivery{
		logger:    logger,
		ucShelter: ucShelter,
		options:   options,
	}
}
