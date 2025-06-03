package rolepermission

import (
	"context"
	"cpru-api/internal/domain/rolepermission"
	"log"
  "cpru-api/pkg/dberror"
  "fmt"
)

type RolePermissionService struct {
	repo rolepermission.Repository
}

func NewRolePermissionService(repo rolepermission.Repository) *RolePermissionService {
	return &RolePermissionService{repo: repo}
}

// Assign เชื่อม role กับ permission
func (s *RolePermissionService) Assign(ctx context.Context, roleID, permissionID int) (*rolepermission.RolePermission, error) {
	rp, err := s.repo.Assign(ctx, roleID, permissionID)
	if err != nil {
		if parsed := dberror.ParseUniqueConstraintError(err); parsed != nil {
			return nil, parsed
		}
    if parsed := dberror.ParseForeignKeyConstraintError(err); parsed != nil {
		  return nil, parsed
	  }
		log.Printf("failed to assign permission %d to role %d: %v", permissionID, roleID, err)
		return nil, err
	}
	return rp, nil
}

// Remove แยก role กับ permission
func (s *RolePermissionService) Remove(ctx context.Context, roleID, permissionID int) error {
  exists, err := s.repo.Exists(ctx, roleID, permissionID)
	if !exists {
		return fmt.Errorf("role-permission relation not found")
	}

	err = s.repo.Remove(ctx, roleID, permissionID)
	if err != nil {
		log.Printf("failed to remove permission %d from role %d: %v", permissionID, roleID, err)
	}
	return err
}

// GetByRoleID คืน list ของ rolepermission โดยใช้ roleID
func (s *RolePermissionService) GetByRoleID(ctx context.Context, roleID int) ([]*rolepermission.RolePermission, error) {
	rps, err := s.repo.GetByRoleID(ctx, roleID)
	if err != nil {
		log.Printf("failed to get rolepermissions by roleID %d: %v", roleID, err)
		return nil, err
	}
	return rps, nil
}

// GetByPermissionID คืน list ของ rolepermission โดยใช้ permissionID
func (s *RolePermissionService) GetByPermissionID(ctx context.Context, permissionID int) ([]*rolepermission.RolePermission, error) {
	rps, err := s.repo.GetByPermissionID(ctx, permissionID)
	if err != nil {
		log.Printf("failed to get rolepermissions by permissionID %d: %v", permissionID, err)
		return nil, err
	}
	return rps, nil
}
