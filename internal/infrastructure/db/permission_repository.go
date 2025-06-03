package db

import (
	"context"
	"cpru-api/ent"
  "cpru-api/internal/domain/permission"
)

type permissionRepository struct {
	client *ent.Client
}

func NewPermissionRepository(client *ent.Client) permission.Repository {
	return &permissionRepository{client: client}
}

func (r *permissionRepository) Create(ctx context.Context, p *permission.Permission) (*permission.Permission, error) {
	entPerm, err := r.client.Permission.
		Create().
		SetCode(p.Code).
		SetNillableDescription(p.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPermission(entPerm), nil
}

func (r *permissionRepository) Update(ctx context.Context, p *permission.Permission) (*permission.Permission, error) {
	entPerm, err := r.client.Permission.Get(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	entUpdated, err := entPerm.Update().
		SetCode(p.Code).
		SetNillableDescription(p.Description).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entToDomainPermission(entUpdated), nil
}

func (r *permissionRepository) Delete(ctx context.Context, id int) error {
	return r.client.Permission.DeleteOneID(id).Exec(ctx)
}

func (r *permissionRepository) GetByID(ctx context.Context, id int) (*permission.Permission, error) {
	entPerm, err := r.client.Permission.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entToDomainPermission(entPerm), nil
}

func (r *permissionRepository) GetAll(ctx context.Context) ([]*permission.Permission, error) {
	entPerms, err := r.client.Permission.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var result []*permission.Permission
	for _, entPerm := range entPerms {
		result = append(result, entToDomainPermission(entPerm))
	}
	return result, nil
}

// helper
func entToDomainPermission(entPerm *ent.Permission) *permission.Permission {
	return &permission.Permission{
		ID:          entPerm.ID,
		Code:        entPerm.Code,
		Description: entPerm.Description,
	}
}
