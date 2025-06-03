package permission

import (
	"context"
	"cpru-api/internal/domain/permission"
  "cpru-api/pkg/dberror"
	"log"
)

type PermissionService struct {
	repo permission.Repository
}

func NewPermissionService(repo permission.Repository) *PermissionService {
	return &PermissionService{repo: repo}
}

// Create สร้าง Permission ใหม่
func (s *PermissionService) Create(ctx context.Context, p *permission.Permission) (*permission.Permission, error) {
	created, err := s.repo.Create(ctx, p)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr 
    }
		log.Printf("failed to create permission: %v", err)
		return nil, err
	}
	return created, nil
}

// Update แก้ไข Permission
func (s *PermissionService) Update(ctx context.Context, p *permission.Permission) (*permission.Permission, error) {
	updated, err := s.repo.Update(ctx, p)
	if err != nil {
    if parsedErr := dberror.ParseUniqueConstraintError(err); parsedErr != nil {
        return nil, parsedErr // ส่ง error ที่สวยงามกลับไปเลย
    }
		log.Printf("failed to update permission: %v", err)
		return nil, err
	}
	return updated, nil
}

// Delete ลบ Permission
func (s *PermissionService) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return notFoundErr
	  }
		log.Printf("failed to delete permission with ID %d: %v", id, err)
	}
	return err
}

// GetAll ดึง Permission ทั้งหมด
func (s *PermissionService) GetAll(ctx context.Context) ([]*permission.Permission, error) {
	list, err := s.repo.GetAll(ctx)
	if err != nil {
		log.Printf("failed to list permissions: %v", err)
		return nil, err
	}
	return list, nil
}

// GetByID ดึง Permission ตาม ID
func (s *PermissionService) GetByID(ctx context.Context, id int) (*permission.Permission, error) {
	p, err := s.repo.GetByID(ctx, id)
	if err != nil {
    if notFoundErr := dberror.ParseNotFoundError(err); notFoundErr != nil {
		  return nil, notFoundErr
	  }
		log.Printf("failed to get permission with ID %d: %v", id, err)
		return nil, err
	}
	return p, nil
}
