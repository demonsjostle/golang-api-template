package graphql

import (
  "cpru-api/internal/application/auth"
  "cpru-api/internal/application/user"
  "cpru-api/internal/application/role"
  "cpru-api/internal/application/permission"
  "cpru-api/internal/application/rolepermission"
  "cpru-api/internal/application/userpermission"
  "cpru-api/internal/application/userrole" 
)

type Resolver struct {
  AuthService *auth.AuthService
	UserService *user.UserService	
  RoleService *role.RoleService
  PermissionService *permission.PermissionService
  RolePermissionService *rolepermission.RolePermissionService
  UserPermissionService *userpermission.UserPermissionService
  UserRoleService *userrole.UserRoleService
}
