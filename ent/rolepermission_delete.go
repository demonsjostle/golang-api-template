// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"cpru-api/ent/predicate"
	"cpru-api/ent/rolepermission"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RolePermissionDelete is the builder for deleting a RolePermission entity.
type RolePermissionDelete struct {
	config
	hooks    []Hook
	mutation *RolePermissionMutation
}

// Where appends a list predicates to the RolePermissionDelete builder.
func (rpd *RolePermissionDelete) Where(ps ...predicate.RolePermission) *RolePermissionDelete {
	rpd.mutation.Where(ps...)
	return rpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rpd *RolePermissionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rpd.sqlExec, rpd.mutation, rpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rpd *RolePermissionDelete) ExecX(ctx context.Context) int {
	n, err := rpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rpd *RolePermissionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rolepermission.Table, sqlgraph.NewFieldSpec(rolepermission.FieldID, field.TypeInt))
	if ps := rpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rpd.mutation.done = true
	return affected, err
}

// RolePermissionDeleteOne is the builder for deleting a single RolePermission entity.
type RolePermissionDeleteOne struct {
	rpd *RolePermissionDelete
}

// Where appends a list predicates to the RolePermissionDelete builder.
func (rpdo *RolePermissionDeleteOne) Where(ps ...predicate.RolePermission) *RolePermissionDeleteOne {
	rpdo.rpd.mutation.Where(ps...)
	return rpdo
}

// Exec executes the deletion query.
func (rpdo *RolePermissionDeleteOne) Exec(ctx context.Context) error {
	n, err := rpdo.rpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rolepermission.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rpdo *RolePermissionDeleteOne) ExecX(ctx context.Context) {
	if err := rpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
