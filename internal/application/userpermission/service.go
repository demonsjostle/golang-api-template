package userpermission

import (
	"context"
	"cpru-api/internal/domain/userpermission"
	"cpru-api/pkg/dberror"
	"fmt"
	"log"
)

type UserPermissionService struct {
	repo userpermission.Repository
}

func NewUserPermissionService(repo userpermission.Repository) *UserPermissionService {
	return &UserPermissionService{repo: repo}
}

// Assign เชื่อม user กับ permission
func (s *UserPermissionService) Assign(ctx context.Context, userID, permissionID int) (*userpermission.UserPermission, error) {
	up, err := s.repo.Assign(ctx, userID, permissionID)
	if err != nil {
		if parsed := dberror.ParseUniqueConstraintError(err); parsed != nil {
			return nil, parsed
		}
    if parsed := dberror.ParseForeignKeyConstraintError(err); parsed != nil {
		  return nil, parsed
	  }
		log.Printf("failed to assign permission %d to user %d: %v", permissionID, userID, err)
		return nil, err
	}
	return up, nil
}

// Remove แยก user กับ permission
func (s *UserPermissionService) Remove(ctx context.Context, userID, permissionID int) error {
	exists, err := s.repo.Exists(ctx, userID, permissionID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user-permission relation not found")
	}

	err = s.repo.Remove(ctx, userID, permissionID)
	if err != nil {
		log.Printf("failed to remove permission %d from user %d: %v", permissionID, userID, err)
	}
	return err
}

// GetByUserID คืน list ของ userpermission โดยใช้ userID
func (s *UserPermissionService) GetByUserID(ctx context.Context, userID int) ([]*userpermission.UserPermission, error) {
	ups, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		log.Printf("failed to get userpermissions by userID %d: %v", userID, err)
		return nil, err
	}
	return ups, nil
}

// GetByPermissionID คืน list ของ userpermission โดยใช้ permissionID
func (s *UserPermissionService) GetByPermissionID(ctx context.Context, permissionID int) ([]*userpermission.UserPermission, error) {
	ups, err := s.repo.GetByPermissionID(ctx, permissionID)
	if err != nil {
		log.Printf("failed to get userpermissions by permissionID %d: %v", permissionID, err)
		return nil, err
	}
	return ups, nil
}
