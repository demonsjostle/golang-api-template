package db

import (
    "context"
    "cpru-api/ent"
    "cpru-api/internal/domain/rolepermission"
    entRolePermission "cpru-api/ent/rolepermission"
)

type rolePermissionRepository struct {
    client *ent.Client
}

func NewRolePermissionRepository(client *ent.Client) *rolePermissionRepository {
    return &rolePermissionRepository{client: client}
}

func (r *rolePermissionRepository) Assign(ctx context.Context, roleID, permissionID int) (*rolepermission.RolePermission, error) {
    rp, err := r.client.RolePermission.
        Create().
        SetRoleID(roleID).
        SetPermissionID(permissionID).
        Save(ctx)
    if err != nil {
        return nil, err
    }
    return &rolepermission.RolePermission{
        ID:           rp.ID,
        RoleID:       rp.RoleID,
        PermissionID: rp.PermissionID,
    }, nil
}

func (r *rolePermissionRepository) Remove(ctx context.Context, roleID, permissionID int) error {
    _, err := r.client.RolePermission.
      Delete().
      Where(
        entRolePermission.RoleIDEQ(roleID),
        entRolePermission.PermissionIDEQ(permissionID),
      ).
      Exec(ctx)
    return err
}

func (r *rolePermissionRepository) GetByRoleID(ctx context.Context, roleID int) ([]*rolepermission.RolePermission, error) {
    results, err := r.client.RolePermission.
        Query().
        Where(entRolePermission.RoleIDEQ(roleID)).
        All(ctx)
    if err != nil {
        return nil, err
    }
    var list []*rolepermission.RolePermission
    for _, rp := range results {
        list = append(list, &rolepermission.RolePermission{
            ID:           rp.ID,
            RoleID:       rp.RoleID,
            PermissionID: rp.PermissionID,
        })
    }
    return list, nil
}

func (r *rolePermissionRepository) GetByPermissionID(ctx context.Context, permissionID int) ([]*rolepermission.RolePermission, error) {
    results, err := r.client.RolePermission.
        Query().
        Where(entRolePermission.PermissionIDEQ(permissionID)).
        All(ctx)
    if err != nil {
        return nil, err
    }
    var list []*rolepermission.RolePermission
    for _, rp := range results {
        list = append(list, &rolepermission.RolePermission{
            ID:           rp.ID,
            RoleID:       rp.RoleID,
            PermissionID: rp.PermissionID,
        })
    }
    return list, nil
}

func (r *rolePermissionRepository) Exists(ctx context.Context, roleID, permissionID int) (bool, error) {
	return r.client.RolePermission.
		Query().
		Where(
			entRolePermission.RoleIDEQ(roleID),
			entRolePermission.PermissionIDEQ(permissionID),
		).Exist(ctx)
}
