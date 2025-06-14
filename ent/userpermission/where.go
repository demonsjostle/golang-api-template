// Code generated by ent, DO NOT EDIT.

package userpermission

import (
	"cpru-api/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldUserID, v))
}

// PermissionID applies equality check predicate on the "permission_id" field. It's identical to PermissionIDEQ.
func PermissionID(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldPermissionID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNotIn(FieldUserID, vs...))
}

// PermissionIDEQ applies the EQ predicate on the "permission_id" field.
func PermissionIDEQ(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldEQ(FieldPermissionID, v))
}

// PermissionIDNEQ applies the NEQ predicate on the "permission_id" field.
func PermissionIDNEQ(v int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNEQ(FieldPermissionID, v))
}

// PermissionIDIn applies the In predicate on the "permission_id" field.
func PermissionIDIn(vs ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldIn(FieldPermissionID, vs...))
}

// PermissionIDNotIn applies the NotIn predicate on the "permission_id" field.
func PermissionIDNotIn(vs ...int) predicate.UserPermission {
	return predicate.UserPermission(sql.FieldNotIn(FieldPermissionID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserPermission {
	return predicate.UserPermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserPermission {
	return predicate.UserPermission(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPermission applies the HasEdge predicate on the "permission" edge.
func HasPermission() predicate.UserPermission {
	return predicate.UserPermission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PermissionTable, PermissionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPermissionWith applies the HasEdge predicate on the "permission" edge with a given conditions (other predicates).
func HasPermissionWith(preds ...predicate.Permission) predicate.UserPermission {
	return predicate.UserPermission(func(s *sql.Selector) {
		step := newPermissionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserPermission) predicate.UserPermission {
	return predicate.UserPermission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserPermission) predicate.UserPermission {
	return predicate.UserPermission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserPermission) predicate.UserPermission {
	return predicate.UserPermission(sql.NotPredicates(p))
}
