package rolepermission

import "context"

type Repository interface {
    Assign(ctx context.Context, roleID, permissionID int) (*RolePermission, error)
    Remove(ctx context.Context, roleID, permissionID int) error
    GetByRoleID(ctx context.Context, roleID int) ([]*RolePermission, error)
    GetByPermissionID(ctx context.Context, permissionID int) ([]*RolePermission, error)
    Exists(ctx context.Context, roleID, permissionID int) (bool, error)
}
