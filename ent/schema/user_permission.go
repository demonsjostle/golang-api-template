package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
  "entgo.io/ent/schema/index"
)

// UserPermission คือ join table ระหว่าง User กับ Permission
type UserPermission struct {
	ent.Schema
}

func (UserPermission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("FK ไปยัง user"),

		field.Int("permission_id").
			Comment("FK ไปยัง permission"),
	}
}

func (UserPermission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_permissions").
			Field("user_id").
			Unique().
			Required().
			Comment("ความสัมพันธ์กับ user"),

		edge.From("permission", Permission.Type).
			Ref("user_permissions").
			Field("permission_id").
			Unique().
			Required().
			Comment("ความสัมพันธ์กับ permission"),
	}
}

// Indexes เพิ่ม unique constraint กันข้อมูลซ้ำ
func (UserPermission) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "permission_id").Unique(),
	}
}
