package data

var UserPermissionSeed = []struct {
	Username       string
	PermissionCode string
}{
	{"manager", "user:delete"}, // MANAGER ได้ user:delete
	{"user",    "user:read"},   // USER ได้ user:read
}
