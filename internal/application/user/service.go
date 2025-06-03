package user

import (
	"context"
	"cpru-api/internal/domain/user"
  "cpru-api/pkg/dberror"
  "cpru-api/pkg/security"
  "log"  
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

// Create user
func (s *UserService) Create(ctx context.Context, u *user.User) (*user.User, error) {
  hashed, err := security.HashPassword(u.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return nil, err
	}
	u.Password = hashed

	createdUser, err := s.repo.Create(ctx, u)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr 
    }
		log.Printf("failed to create user: %v", err)
		return nil, err
	}
	return createdUser, nil
}

// Update user
func (s *UserService) Update(ctx context.Context, u *user.User) (*user.User, error) {
  if u.Password != "" {
		hashed, err := security.HashPassword(u.Password)
		if err != nil {
			log.Printf("failed to hash password: %v", err)
			return nil, err
		}
		u.Password = hashed
	}

	updatedUser, err := s.repo.Update(ctx, u)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr // ส่ง error ที่สวยงามกลับไปเลย
    }
		log.Printf("failed to update user: %v", err)
		return nil, err
	}
	return updatedUser, nil
}

// Delete user
func (s *UserService) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return notFoundErr
	  }	
    log.Printf("failed to delete user with ID %d: %v", id, err)
	} 
	return err
}

// Get all users
func (s *UserService) GetAll(ctx context.Context) ([]*user.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("failed to list users: %v", err)
		return nil, err
	}
	return users, nil
}

// Get user by ID
func (s *UserService) GetByID(ctx context.Context, id int) (*user.User, error) {
	u, err := s.repo.GetByID(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return nil, notFoundErr
	  }
		log.Printf("failed to get user with ID %d: %v", id, err)
		return nil, err
	}
	return u, nil
}
