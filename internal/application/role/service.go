package role

import (
	"context"
	"cpru-api/internal/domain/role"
  "cpru-api/pkg/dberror"
	"log"
)

type RoleService struct {
	repo role.Repository
}

func NewRoleService(repo role.Repository) *RoleService {
	return &RoleService{repo: repo}
}

// Create สร้าง Role ใหม่
func (s *RoleService) Create(ctx context.Context, r *role.Role) (*role.Role, error) {
	created, err := s.repo.Create(ctx, r)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr 
    }
		log.Printf("failed to create role: %v", err)
		return nil, err
	}
	return created, nil
}

// Update แก้ไข Role
func (s *RoleService) Update(ctx context.Context, r *role.Role) (*role.Role, error) {
	updated, err := s.repo.Update(ctx, r)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr 
    }
		log.Printf("failed to update role: %v", err)
		return nil, err
	}
	return updated, nil
}

// Delete ลบ Role
func (s *RoleService) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return notFoundErr
	  }	
    log.Printf("failed to delete role with ID %d: %v", id, err)
	}
	return err
}

// GetAll ดึง Role ทั้งหมด
func (s *RoleService) GetAll(ctx context.Context) ([]*role.Role, error) {
	roles, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("failed to list roles: %v", err)
		return nil, err
	}
	return roles, nil
}

// GetByID ดึง Role ตาม ID
func (s *RoleService) GetByID(ctx context.Context, id int) (*role.Role, error) {
	r, err := s.repo.GetByID(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return nil, notFoundErr
	  }
	  log.Printf("failed to get role with ID %d: %v", id, err)	
		return nil, err
	}
	return r, nil
}
