package userrole

import "context"

type Repository interface {
	Assign(ctx context.Context, userID, roleID int) (*UserRole, error)
	Remove(ctx context.Context, userID, roleID int) error
	GetByUserID(ctx context.Context, userID int) ([]*UserRole, error)
	GetByRoleID(ctx context.Context, roleID int) ([]*UserRole, error)
	Exists(ctx context.Context, userID, roleID int) (bool, error)
}
