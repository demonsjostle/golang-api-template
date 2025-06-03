package userrole

import (
	"context"
	"cpru-api/internal/domain/userrole"
	"cpru-api/pkg/dberror"
	"fmt"
	"log"
)

type UserRoleService struct {
	repo userrole.Repository
}

func NewUserRoleService(repo userrole.Repository) *UserRoleService {
	return &UserRoleService{repo: repo}
}

// Assign เชื่อม user กับ role
func (s *UserRoleService) Assign(ctx context.Context, userID, roleID int) (*userrole.UserRole, error) {
	ur, err := s.repo.Assign(ctx, userID, roleID)
	if err != nil {
		if parsed := dberror.ParseUniqueConstraintError(err); parsed != nil {
			return nil, parsed
		}
		if parsed := dberror.ParseForeignKeyConstraintError(err); parsed != nil {
			return nil, parsed
		}
		log.Printf("failed to assign role %d to user %d: %v", roleID, userID, err)
		return nil, err
	}
	return ur, nil
}

// Remove แยก user กับ role
func (s *UserRoleService) Remove(ctx context.Context, userID, roleID int) error {
	exists, err := s.repo.Exists(ctx, userID, roleID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user-role relation not found")
	}

	err = s.repo.Remove(ctx, userID, roleID)
	if err != nil {
		log.Printf("failed to remove role %d from user %d: %v", roleID, userID, err)
	}
	return err
}

// GetByUserID คืน list ของ userrole โดยใช้ userID
func (s *UserRoleService) GetByUserID(ctx context.Context, userID int) ([]*userrole.UserRole, error) {
	urs, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		log.Printf("failed to get userroles by userID %d: %v", userID, err)
		return nil, err
	}
	return urs, nil
}

// GetByRoleID คืน list ของ userrole โดยใช้ roleID
func (s *UserRoleService) GetByRoleID(ctx context.Context, roleID int) ([]*userrole.UserRole, error) {
	urs, err := s.repo.GetByRoleID(ctx, roleID)
	if err != nil {
		log.Printf("failed to get userroles by roleID %d: %v", roleID, err)
		return nil, err
	}
	return urs, nil
}
