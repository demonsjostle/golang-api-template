// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// 📦 ข้อมูลตอบกลับหลังจาก register หรือ login สำเร็จ
type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

// 📦 Input for creating a permission
type CreatePermissionInput struct {
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

// 📦 Input for creating a role
type CreateRoleInput struct {
	Name string `json:"name"`
}

// 📦 Input for creating a new user
type CreateUserInput struct {
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Phone        *string `json:"phone,omitempty"`
	Password     string  `json:"password"`
	DepartmentID *int    `json:"departmentId,omitempty"`
}

// 🔐 ข้อมูลสำหรับล็อกอิน (Login)
type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Permission struct {
	ID          string  `json:"id"`
	Code        string  `json:"code"`
	Description *string `json:"description,omitempty"`
}

type Query struct {
}

// 🔐 ข้อมูลสำหรับสมัครสมาชิก (Register)
type RegisterInput struct {
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Phone        *string `json:"phone,omitempty"`
	Password     string  `json:"password"`
	DepartmentID *int    `json:"departmentId,omitempty"`
}

type Role struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 🔗 Entity สำหรับความสัมพันธ์ระหว่าง Role และ Permission
type RolePermission struct {
	ID           string `json:"id"`
	RoleID       int    `json:"roleID"`
	PermissionID int    `json:"permissionID"`
}

// ✏️ Input for updating an existing permission
type UpdatePermissionInput struct {
	Code        *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// ✏️ Input for updating a role
type UpdateRoleInput struct {
	Name *string `json:"name,omitempty"`
}

// ✏️ Input for updating an existing user
type UpdateUserInput struct {
	Name         *string `json:"name,omitempty"`
	Surname      *string `json:"surname,omitempty"`
	Username     *string `json:"username,omitempty"`
	Email        *string `json:"email,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	Password     *string `json:"password,omitempty"`
	DepartmentID *int    `json:"departmentId,omitempty"`
}

// 👤 User entity
type User struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Surname      string  `json:"surname"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Phone        *string `json:"phone,omitempty"`
	DepartmentID *int    `json:"departmentId,omitempty"`
}

// 🔗 UserPermission entity (ความสัมพันธ์ระหว่างผู้ใช้กับสิทธิ์)
type UserPermission struct {
	ID           string `json:"id"`
	UserID       int    `json:"userID"`
	PermissionID int    `json:"permissionID"`
}

// 🔗 UserRole entity (ความสัมพันธ์ระหว่าง User และ Role)
type UserRole struct {
	ID     string `json:"id"`
	UserID int    `json:"userID"`
	RoleID int    `json:"roleID"`
}
