package user

import "context"

type Repository interface {	
	Create(ctx context.Context, user *User) (*User, error)	
	
	Update(ctx context.Context, user *User) (*User, error)

	Delete(ctx context.Context, id int) error
	
	GetAll(ctx context.Context) ([]*User, error)

  GetByID(ctx context.Context, id int) (*User, error)

  GetByUsername(ctx context.Context, username string) (*User, error)
}
