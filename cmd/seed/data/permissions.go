package data

var PermissionSeed = []struct {
	Code        string
	Description string
}{
	// User CRUD
	{"user:create", "สร้างผู้ใช้ใหม่"},
	{"user:read",   "อ่านข้อมูลผู้ใช้"},
	{"user:update", "แก้ไขข้อมูลผู้ใช้"},
	{"user:delete", "ลบผู้ใช้"},

	// Department CRUD
	{"department:create", "สร้างแผนกใหม่"},
	{"department:read",   "อ่านข้อมูลแผนก"},
	{"department:update", "แก้ไขข้อมูลแผนก"},
	{"department:delete", "ลบแผนก"},

	// Role CRUD
	{"role:create", "สร้างบทบาทใหม่"},
	{"role:read",   "อ่านข้อมูลบทบาท"},
	{"role:update", "แก้ไขข้อมูลบทบาท"},
	{"role:delete", "ลบบทบาท"},

	// Permission CRUD
	{"permission:create", "สร้างสิทธิ์ใหม่"},
	{"permission:read",   "อ่านข้อมูลสิทธิ์"},
	{"permission:update", "แก้ไขข้อมูลสิทธิ์"},
	{"permission:delete", "ลบสิทธิ์"},

	// RolePermission (join) CRUD
	{"role_permission:create", "สร้างความสัมพันธ์บทบาท-สิทธิ์"},
	{"role_permission:read",   "อ่านความสัมพันธ์บทบาท-สิทธิ์"},
	{"role_permission:update", "แก้ไขความสัมพันธ์บทบาท-สิทธิ์"},
	{"role_permission:delete", "ลบความสัมพันธ์บทบาท-สิทธิ์"},

	// UserRole (join) CRUD
	{"user_role:create", "สร้างความสัมพันธ์ผู้ใช้-บทบาท"},
	{"user_role:read",   "อ่านความสัมพันธ์ผู้ใช้-บทบาท"},
	{"user_role:update", "แก้ไขความสัมพันธ์ผู้ใช้-บทบาท"},
	{"user_role:delete", "ลบความสัมพันธ์ผู้ใช้-บทบาท"},

	// UserPermission (join) CRUD
	{"user_permission:create", "สร้างความสัมพันธ์ผู้ใช้-สิทธิ์"},
	{"user_permission:read",   "อ่านความสัมพันธ์ผู้ใช้-สิทธิ์"},
	{"user_permission:update", "แก้ไขความสัมพันธ์ผู้ใช้-สิทธิ์"},
	{"user_permission:delete", "ลบความสัมพันธ์ผู้ใช้-สิทธิ์"},
}
