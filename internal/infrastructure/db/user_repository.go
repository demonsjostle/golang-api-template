package db

import (
	"context"
	"cpru-api/ent"
	"cpru-api/internal/domain/user"
  entUser "cpru-api/ent/user"
  "cpru-api/pkg/ptr"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *userRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, u *user.User) (*user.User, error) {
	entUser, err := r.client.User.
		Create().
		SetName(u.Name).
		SetSurname(u.Surname).
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetNillablePhone(u.Phone).
		SetNillableDepartmentID(u.DepartmentID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return entToDomainUser(entUser), nil
}

func (r *userRepository) Update(ctx context.Context, u *user.User) (*user.User, error) {
	existing, err := r.client.User.Get(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	entUser, err := existing.Update().
		SetName(u.Name).
		SetSurname(u.Surname).
		SetUsername(u.Username).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetNillablePhone(u.Phone).
		SetNillableDepartmentID(u.DepartmentID).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return entToDomainUser(entUser), nil
}

func (r *userRepository) Delete(ctx context.Context, id int) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}

func (r *userRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	entUsers, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var users []*user.User
	for _, entUser := range entUsers {
		users = append(users, entToDomainUser(entUser))
	}
	return users, nil
}

func (r *userRepository) GetByID(ctx context.Context, id int) (*user.User, error) {
	entUser, err := r.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return entToDomainUser(entUser), nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*user.User, error) {
    entU, err := r.client.User.Query().
        Where(entUser.UsernameEQ(username)).
        Only(ctx)
    if err != nil {
        return nil, err
    }
    return entToDomainUser(entU), nil
}

// === Helper ===

func entToDomainUser(entUser *ent.User) *user.User {
	return &user.User{
		ID:           entUser.ID,
		Name:         entUser.Name,
		Surname:      entUser.Surname,
		Username:     entUser.Username,
		Email:        entUser.Email,
		Phone:        ptr.StringToPointer(entUser.Phone),
		Password:     entUser.Password,
		DepartmentID: ptr.IntToPointer(entUser.DepartmentID),
	}
}


