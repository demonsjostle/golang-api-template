package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("ชื่อ"),

		field.String("surname").
			NotEmpty().
			Comment("นามสกุล"),

		field.String("username").
			NotEmpty().
			Unique().
			Comment("ชื่อผู้ใช้"),

		field.String("email").
			NotEmpty().
			Unique().
			Comment("อีเมล"),

		field.String("phone").
			Optional().
			Comment("เบอร์โทรศัพท์"),

		field.String("password").
			NotEmpty().
			Comment("รหัสผ่าน (hash แล้ว)"),
	
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{	
    edge.To("user_roles", UserRole.Type).
	    Comment("บทบาทของผู้ใช้"),
    edge.To("user_permissions", UserPermission.Type).
			Comment("สิทธิ์ที่ระบุให้ผู้ใช้นี้โดยตรง"),
	}
}
