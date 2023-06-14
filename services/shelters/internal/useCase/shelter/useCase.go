package shelter

import "github.com/nourbol/zhandos/services/shelters/internal/useCase/adapters/storage"

type UseCase struct {
	adapterStorage storage.Storage
	options        Options
}

type Options struct {
}

func New(adapterStorage storage.Storage, options Options) *UseCase {
	return &UseCase{
		adapterStorage: adapterStorage,
		options:        options,
	}
}
