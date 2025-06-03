package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

type Permission struct {
	ent.Schema
}

func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").NotEmpty().Unique(),
    field.String("description").Optional().Nillable(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		// ความสัมพันธ์กับ Role ผ่าน RolePermission
		edge.To("role_permissions", RolePermission.Type),

		// ความสัมพันธ์กับ User ผ่าน UserPermission
		edge.To("user_permissions", UserPermission.Type),
	}
}
