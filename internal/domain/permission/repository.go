package permission

import "context"

type Repository interface {
	Create(ctx context.Context, permission *Permission) (*Permission, error)
	Update(ctx context.Context, permission *Permission) (*Permission, error)
	Delete(ctx context.Context, id int) error
  GetAll(ctx context.Context) ([]*Permission, error)
	GetByID(ctx context.Context, id int) (*Permission, error)	
}
