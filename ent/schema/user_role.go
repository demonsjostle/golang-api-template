package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
  "entgo.io/ent/schema/index"
)

// UserRole คือ join table ระหว่าง User กับ Role
type UserRole struct {
	ent.Schema
}

func (UserRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("FK ไปยัง user"),

		field.Int("role_id").
			Comment("FK ไปยัง role"),
	}
}

func (UserRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("user_roles").
			Field("user_id").
			Unique().
			Required().
			Comment("ความสัมพันธ์กับ user"),

		edge.From("role", Role.Type).
			Ref("user_roles").
			Field("role_id").
			Unique().
			Required().
			Comment("ความสัมพันธ์กับ role"),
	}
}

func (UserRole) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "role_id").Unique(),
	}
}
