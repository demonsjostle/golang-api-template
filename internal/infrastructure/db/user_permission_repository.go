package db

import (
	"context"
	"cpru-api/ent"
	"cpru-api/internal/domain/userpermission"
  entUserPermission "cpru-api/ent/userpermission"
)

type userPermissionRepository struct {
	client *ent.Client
}

func NewUserPermissionRepository(client *ent.Client) *userPermissionRepository {
	return &userPermissionRepository{client: client}
}

func (r *userPermissionRepository) Assign(ctx context.Context, userID, permissionID int) (*userpermission.UserPermission, error) {
	up, err := r.client.UserPermission.
		Create().
		SetUserID(userID).
		SetPermissionID(permissionID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &userpermission.UserPermission{
		ID:           up.ID,
		UserID:       up.UserID,
		PermissionID: up.PermissionID,
	}, nil
}

func (r *userPermissionRepository) Remove(ctx context.Context, userID, permissionID int) error {
	_, err := r.client.UserPermission.
		Delete().
		Where(
			entUserPermission.UserIDEQ(userID),
			entUserPermission.PermissionIDEQ(permissionID),
		).
		Exec(ctx)
	return err
}

func (r *userPermissionRepository) GetByUserID(ctx context.Context, userID int) ([]*userpermission.UserPermission, error) {
	results, err := r.client.UserPermission.
		Query().
		Where(entUserPermission.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var list []*userpermission.UserPermission
	for _, up := range results {
		list = append(list, &userpermission.UserPermission{
			ID:           up.ID,
			UserID:       up.UserID,
			PermissionID: up.PermissionID,
		})
	}
	return list, nil
}

func (r *userPermissionRepository) GetByPermissionID(ctx context.Context, permissionID int) ([]*userpermission.UserPermission, error) {
	results, err := r.client.UserPermission.
		Query().
		Where(entUserPermission.PermissionIDEQ(permissionID)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	var list []*userpermission.UserPermission
	for _, up := range results {
		list = append(list, &userpermission.UserPermission{
			ID:           up.ID,
			UserID:       up.UserID,
			PermissionID: up.PermissionID,
		})
	}
	return list, nil
}

func (r *userPermissionRepository) Exists(ctx context.Context, userID, permissionID int) (bool, error) {
	return r.client.UserPermission.
		Query().
		Where(
			entUserPermission.UserIDEQ(userID),
			entUserPermission.PermissionIDEQ(permissionID),
		).
		Exist(ctx)
}
