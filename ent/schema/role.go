package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

type Role struct {
	ent.Schema
}

func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		// ความสัมพันธ์กับ User ผ่าน UserRole
		edge.To("user_roles", UserRole.Type),

		// ความสัมพันธ์กับ Permission ผ่าน RolePermission
		edge.To("role_permissions", RolePermission.Type),
	}
}
