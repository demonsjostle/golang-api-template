package userpermission

import "context"

type Repository interface {
  Assign(ctx context.Context, userID, permissionID int) (*UserPermission, error)
	Remove(ctx context.Context, userID, permissionID int) error
	GetByUserID(ctx context.Context, userID int) ([]*UserPermission, error)
	GetByPermissionID(ctx context.Context, permissionID int) ([]*UserPermission, error)
	Exists(ctx context.Context, userID, permissionID int) (bool, error)
}
