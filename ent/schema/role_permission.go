package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
  "entgo.io/ent/schema/index"
)

// RolePermission holds the schema definition for the join table between Role and Permission.
type RolePermission struct {
	ent.Schema
}

// Fields of the RolePermission.
func (RolePermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("role_id").
			Comment("FK ไปยัง role"),

		field.Int("permission_id").
			Comment("FK ไปยัง permission"),
	}
}

// Edges of the RolePermission.
func (RolePermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("role", Role.Type).
			Ref("role_permissions").
			Field("role_id").
			Unique().
			Required().Comment("ความสัมพันธ์กับ role"),

		edge.From("permission", Permission.Type).
			Ref("role_permissions").
			Field("permission_id").
			Unique().
			Required().Comment("ความสัมพันธ์กับ permission"),
	}
}

// Indexes เพิ่ม unique constraint กันข้อมูลซ้ำ
func (RolePermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("role_id", "permission_id").Unique(),
	}
}
