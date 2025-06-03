package db

import (
	"context"
	"cpru-api/ent"
	entUserRole "cpru-api/ent/userrole"
	"cpru-api/internal/domain/userrole"
)

type userRoleRepository struct {
	client *ent.Client
}

func NewUserRoleRepository(client *ent.Client) *userRoleRepository {
	return &userRoleRepository{client: client}
}

func (r *userRoleRepository) Assign(ctx context.Context, userID, roleID int) (*userrole.UserRole, error) {
	ur, err := r.client.UserRole.
		Create().
		SetUserID(userID).
		SetRoleID(roleID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &userrole.UserRole{
		ID:     ur.ID,
		UserID: ur.UserID,
		RoleID: ur.RoleID,
	}, nil
}

func (r *userRoleRepository) Remove(ctx context.Context, userID, roleID int) error {
	_, err := r.client.UserRole.
		Delete().
		Where(
			entUserRole.UserIDEQ(userID),
			entUserRole.RoleIDEQ(roleID),
		).
		Exec(ctx)
	return err
}

func (r *userRoleRepository) GetByUserID(ctx context.Context, userID int) ([]*userrole.UserRole, error) {
	results, err := r.client.UserRole.
		Query().
		Where(entUserRole.UserIDEQ(userID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var list []*userrole.UserRole
	for _, ur := range results {
		list = append(list, &userrole.UserRole{
			ID:     ur.ID,
			UserID: ur.UserID,
			RoleID: ur.RoleID,
		})
	}
	return list, nil
}

func (r *userRoleRepository) GetByRoleID(ctx context.Context, roleID int) ([]*userrole.UserRole, error) {
	results, err := r.client.UserRole.
		Query().
		Where(entUserRole.RoleIDEQ(roleID)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	var list []*userrole.UserRole
	for _, ur := range results {
		list = append(list, &userrole.UserRole{
			ID:     ur.ID,
			UserID: ur.UserID,
			RoleID: ur.RoleID,
		})
	}
	return list, nil
}

func (r *userRoleRepository) Exists(ctx context.Context, userID, roleID int) (bool, error) {
	return r.client.UserRole.
		Query().
		Where(
			entUserRole.UserIDEQ(userID),
			entUserRole.RoleIDEQ(roleID),
		).
		Exist(ctx)
}
