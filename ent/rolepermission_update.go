// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"cpru-api/ent/permission"
	"cpru-api/ent/predicate"
	"cpru-api/ent/role"
	"cpru-api/ent/rolepermission"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RolePermissionUpdate is the builder for updating RolePermission entities.
type RolePermissionUpdate struct {
	config
	hooks    []Hook
	mutation *RolePermissionMutation
}

// Where appends a list predicates to the RolePermissionUpdate builder.
func (rpu *RolePermissionUpdate) Where(ps ...predicate.RolePermission) *RolePermissionUpdate {
	rpu.mutation.Where(ps...)
	return rpu
}

// SetRoleID sets the "role_id" field.
func (rpu *RolePermissionUpdate) SetRoleID(i int) *RolePermissionUpdate {
	rpu.mutation.SetRoleID(i)
	return rpu
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rpu *RolePermissionUpdate) SetNillableRoleID(i *int) *RolePermissionUpdate {
	if i != nil {
		rpu.SetRoleID(*i)
	}
	return rpu
}

// SetPermissionID sets the "permission_id" field.
func (rpu *RolePermissionUpdate) SetPermissionID(i int) *RolePermissionUpdate {
	rpu.mutation.SetPermissionID(i)
	return rpu
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (rpu *RolePermissionUpdate) SetNillablePermissionID(i *int) *RolePermissionUpdate {
	if i != nil {
		rpu.SetPermissionID(*i)
	}
	return rpu
}

// SetRole sets the "role" edge to the Role entity.
func (rpu *RolePermissionUpdate) SetRole(r *Role) *RolePermissionUpdate {
	return rpu.SetRoleID(r.ID)
}

// SetPermission sets the "permission" edge to the Permission entity.
func (rpu *RolePermissionUpdate) SetPermission(p *Permission) *RolePermissionUpdate {
	return rpu.SetPermissionID(p.ID)
}

// Mutation returns the RolePermissionMutation object of the builder.
func (rpu *RolePermissionUpdate) Mutation() *RolePermissionMutation {
	return rpu.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rpu *RolePermissionUpdate) ClearRole() *RolePermissionUpdate {
	rpu.mutation.ClearRole()
	return rpu
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (rpu *RolePermissionUpdate) ClearPermission() *RolePermissionUpdate {
	rpu.mutation.ClearPermission()
	return rpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rpu *RolePermissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rpu.sqlSave, rpu.mutation, rpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpu *RolePermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := rpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rpu *RolePermissionUpdate) Exec(ctx context.Context) error {
	_, err := rpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpu *RolePermissionUpdate) ExecX(ctx context.Context) {
	if err := rpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpu *RolePermissionUpdate) check() error {
	if rpu.mutation.RoleCleared() && len(rpu.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RolePermission.role"`)
	}
	if rpu.mutation.PermissionCleared() && len(rpu.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RolePermission.permission"`)
	}
	return nil
}

func (rpu *RolePermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(rolepermission.Table, rolepermission.Columns, sqlgraph.NewFieldSpec(rolepermission.FieldID, field.TypeInt))
	if ps := rpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rpu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.RoleTable,
			Columns: []string{rolepermission.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.RoleTable,
			Columns: []string{rolepermission.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rpu.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.PermissionTable,
			Columns: []string{rolepermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpu.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.PermissionTable,
			Columns: []string{rolepermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolepermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rpu.mutation.done = true
	return n, nil
}

// RolePermissionUpdateOne is the builder for updating a single RolePermission entity.
type RolePermissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RolePermissionMutation
}

// SetRoleID sets the "role_id" field.
func (rpuo *RolePermissionUpdateOne) SetRoleID(i int) *RolePermissionUpdateOne {
	rpuo.mutation.SetRoleID(i)
	return rpuo
}

// SetNillableRoleID sets the "role_id" field if the given value is not nil.
func (rpuo *RolePermissionUpdateOne) SetNillableRoleID(i *int) *RolePermissionUpdateOne {
	if i != nil {
		rpuo.SetRoleID(*i)
	}
	return rpuo
}

// SetPermissionID sets the "permission_id" field.
func (rpuo *RolePermissionUpdateOne) SetPermissionID(i int) *RolePermissionUpdateOne {
	rpuo.mutation.SetPermissionID(i)
	return rpuo
}

// SetNillablePermissionID sets the "permission_id" field if the given value is not nil.
func (rpuo *RolePermissionUpdateOne) SetNillablePermissionID(i *int) *RolePermissionUpdateOne {
	if i != nil {
		rpuo.SetPermissionID(*i)
	}
	return rpuo
}

// SetRole sets the "role" edge to the Role entity.
func (rpuo *RolePermissionUpdateOne) SetRole(r *Role) *RolePermissionUpdateOne {
	return rpuo.SetRoleID(r.ID)
}

// SetPermission sets the "permission" edge to the Permission entity.
func (rpuo *RolePermissionUpdateOne) SetPermission(p *Permission) *RolePermissionUpdateOne {
	return rpuo.SetPermissionID(p.ID)
}

// Mutation returns the RolePermissionMutation object of the builder.
func (rpuo *RolePermissionUpdateOne) Mutation() *RolePermissionMutation {
	return rpuo.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (rpuo *RolePermissionUpdateOne) ClearRole() *RolePermissionUpdateOne {
	rpuo.mutation.ClearRole()
	return rpuo
}

// ClearPermission clears the "permission" edge to the Permission entity.
func (rpuo *RolePermissionUpdateOne) ClearPermission() *RolePermissionUpdateOne {
	rpuo.mutation.ClearPermission()
	return rpuo
}

// Where appends a list predicates to the RolePermissionUpdate builder.
func (rpuo *RolePermissionUpdateOne) Where(ps ...predicate.RolePermission) *RolePermissionUpdateOne {
	rpuo.mutation.Where(ps...)
	return rpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rpuo *RolePermissionUpdateOne) Select(field string, fields ...string) *RolePermissionUpdateOne {
	rpuo.fields = append([]string{field}, fields...)
	return rpuo
}

// Save executes the query and returns the updated RolePermission entity.
func (rpuo *RolePermissionUpdateOne) Save(ctx context.Context) (*RolePermission, error) {
	return withHooks(ctx, rpuo.sqlSave, rpuo.mutation, rpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rpuo *RolePermissionUpdateOne) SaveX(ctx context.Context) *RolePermission {
	node, err := rpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rpuo *RolePermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := rpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rpuo *RolePermissionUpdateOne) ExecX(ctx context.Context) {
	if err := rpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rpuo *RolePermissionUpdateOne) check() error {
	if rpuo.mutation.RoleCleared() && len(rpuo.mutation.RoleIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RolePermission.role"`)
	}
	if rpuo.mutation.PermissionCleared() && len(rpuo.mutation.PermissionIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "RolePermission.permission"`)
	}
	return nil
}

func (rpuo *RolePermissionUpdateOne) sqlSave(ctx context.Context) (_node *RolePermission, err error) {
	if err := rpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(rolepermission.Table, rolepermission.Columns, sqlgraph.NewFieldSpec(rolepermission.FieldID, field.TypeInt))
	id, ok := rpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RolePermission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rolepermission.FieldID)
		for _, f := range fields {
			if !rolepermission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rolepermission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rpuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.RoleTable,
			Columns: []string{rolepermission.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.RoleTable,
			Columns: []string{rolepermission.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rpuo.mutation.PermissionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.PermissionTable,
			Columns: []string{rolepermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rpuo.mutation.PermissionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rolepermission.PermissionTable,
			Columns: []string{rolepermission.PermissionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &RolePermission{config: rpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rolepermission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rpuo.mutation.done = true
	return _node, nil
}
