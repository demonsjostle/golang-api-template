package data

var RolePermissionSeed = []struct {
	RoleName       string
	PermissionCode string
}{
	// ADMIN ได้ทุก permission
	{"ADMIN", "user:create"},
	{"ADMIN", "user:read"},
	{"ADMIN", "user:update"},
	{"ADMIN", "user:delete"},
	{"ADMIN", "department:create"},
	{"ADMIN", "department:read"},
	{"ADMIN", "department:update"},
	{"ADMIN", "department:delete"},
	{"ADMIN", "role:create"},
	{"ADMIN", "role:read"},
	{"ADMIN", "role:update"},
	{"ADMIN", "role:delete"},
	{"ADMIN", "permission:create"},
	{"ADMIN", "permission:read"},
	{"ADMIN", "permission:update"},
	{"ADMIN", "permission:delete"},
	{"ADMIN", "role_permission:create"},
	{"ADMIN", "role_permission:read"},
	{"ADMIN", "role_permission:update"},
	{"ADMIN", "role_permission:delete"},
	{"ADMIN", "user_role:create"},
	{"ADMIN", "user_role:read"},
	{"ADMIN", "user_role:update"},
	{"ADMIN", "user_role:delete"},
	{"ADMIN", "user_permission:create"},
	{"ADMIN", "user_permission:read"},
	{"ADMIN", "user_permission:update"},
	{"ADMIN", "user_permission:delete"},

	// MANAGER: user:create, user:read, user:update
	{"MANAGER", "user:create"},
	{"MANAGER", "user:read"},
	{"MANAGER", "user:update"},

	// USER: user:read, department:read
	{"USER", "user:read"},
	{"USER", "department:read"},
}
