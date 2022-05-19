package repository

import (
	"context"
	"port-domain-service/domain/entity"
)

type PortRepository interface {
	Create(ctx context.Context, port *entity.Port) error
	Update(ctx context.Context, port *entity.Port) error
}
