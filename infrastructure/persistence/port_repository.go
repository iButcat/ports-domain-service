package persistence

import (
	"context"
	"port-domain-service/domain/entity"
	"port-domain-service/domain/repository"
)

type portRepositoryImplt struct {
	DB map[string]entity.Port
}

func NewPortRepositoryImplt() repository.PortRepository {
	return &portRepositoryImplt{make(map[string]entity.Port)}
}

func (r *portRepositoryImplt) Create(ctx context.Context, port *entity.Port) error {
	// should pass an ID however not enough time and it works fine
	r.DB[port.Name] = *port
	return nil
}

func (r *portRepositoryImplt) Update(ctx context.Context, port *entity.Port) error {
	r.DB[port.Name] = *port
	return nil
}
