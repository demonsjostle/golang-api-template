package db

import (
	"context"
	"cpru-api/ent"
	"cpru-api/internal/domain/role" 
)

type roleRepository struct {
	client *ent.Client
}

// NewRoleRepository creates a new instance of roleRepository.
func NewRoleRepository(client *ent.Client) role.Repository {
	return &roleRepository{client: client}
}

// Create inserts a new role into the database.
func (r *roleRepository) Create(ctx context.Context, rl *role.Role) (*role.Role, error) {
	entRole, err := r.client.Role.
		Create().
		SetName(rl.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &role.Role{
		ID:   entRole.ID,
		Name: entRole.Name,
	}, nil
}

// Update modifies an existing role by ID.
func (r *roleRepository) Update(ctx context.Context, rl *role.Role) (*role.Role, error) {
	existing, err := r.client.Role.Get(ctx, rl.ID)
	if err != nil {
		return nil, err
	}

	entRole, err := existing.Update().
		SetName(rl.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &role.Role{
		ID:   entRole.ID,
		Name: entRole.Name,
	}, nil
}

// Delete removes a role by ID.
func (r *roleRepository) Delete(ctx context.Context, id int) error {
	return r.client.Role.DeleteOneID(id).Exec(ctx)
}

// GetByID retrieves a role by ID.
func (r *roleRepository) GetByID(ctx context.Context, id int) (*role.Role, error) {
	entRole, err := r.client.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return &role.Role{
		ID:   entRole.ID,
		Name: entRole.Name,
	}, nil
}

// GetAll retrieves all roles.
func (r *roleRepository) GetAll(ctx context.Context) ([]*role.Role, error) {
	entRoles, err := r.client.Role.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	roles := make([]*role.Role, 0, len(entRoles))
	for _, entRole := range entRoles {
		roles = append(roles, &role.Role{
			ID:   entRole.ID,
			Name: entRole.Name,
		})
	}
	return roles, nil
}
