package role

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, role *Role) (*Role, error)
	Update(ctx context.Context, role *Role) (*Role, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]*Role, error)
  GetByID(ctx context.Context, id int) (*Role, error)
}
