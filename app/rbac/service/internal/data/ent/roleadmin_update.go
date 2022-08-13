// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vue-admin/app/rbac/service/internal/data/ent/predicate"
	"vue-admin/app/rbac/service/internal/data/ent/roleadmin"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleAdminUpdate is the builder for updating RoleAdmin entities.
type RoleAdminUpdate struct {
	config
	hooks    []Hook
	mutation *RoleAdminMutation
}

// Where appends a list predicates to the RoleAdminUpdate builder.
func (rau *RoleAdminUpdate) Where(ps ...predicate.RoleAdmin) *RoleAdminUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetRoleID sets the "role_id" field.
func (rau *RoleAdminUpdate) SetRoleID(i int) *RoleAdminUpdate {
	rau.mutation.ResetRoleID()
	rau.mutation.SetRoleID(i)
	return rau
}

// AddRoleID adds i to the "role_id" field.
func (rau *RoleAdminUpdate) AddRoleID(i int) *RoleAdminUpdate {
	rau.mutation.AddRoleID(i)
	return rau
}

// SetAdminID sets the "admin_id" field.
func (rau *RoleAdminUpdate) SetAdminID(i int) *RoleAdminUpdate {
	rau.mutation.ResetAdminID()
	rau.mutation.SetAdminID(i)
	return rau
}

// AddAdminID adds i to the "admin_id" field.
func (rau *RoleAdminUpdate) AddAdminID(i int) *RoleAdminUpdate {
	rau.mutation.AddAdminID(i)
	return rau
}

// Mutation returns the RoleAdminMutation object of the builder.
func (rau *RoleAdminUpdate) Mutation() *RoleAdminMutation {
	return rau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *RoleAdminUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rau.hooks) == 0 {
		affected, err = rau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleAdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rau.mutation = mutation
			affected, err = rau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rau.hooks) - 1; i >= 0; i-- {
			if rau.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rau *RoleAdminUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *RoleAdminUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *RoleAdminUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rau *RoleAdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleadmin.Table,
			Columns: roleadmin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roleadmin.FieldID,
			},
		},
	}
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldRoleID,
		})
	}
	if value, ok := rau.mutation.AddedRoleID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldRoleID,
		})
	}
	if value, ok := rau.mutation.AdminID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldAdminID,
		})
	}
	if value, ok := rau.mutation.AddedAdminID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldAdminID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RoleAdminUpdateOne is the builder for updating a single RoleAdmin entity.
type RoleAdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleAdminMutation
}

// SetRoleID sets the "role_id" field.
func (rauo *RoleAdminUpdateOne) SetRoleID(i int) *RoleAdminUpdateOne {
	rauo.mutation.ResetRoleID()
	rauo.mutation.SetRoleID(i)
	return rauo
}

// AddRoleID adds i to the "role_id" field.
func (rauo *RoleAdminUpdateOne) AddRoleID(i int) *RoleAdminUpdateOne {
	rauo.mutation.AddRoleID(i)
	return rauo
}

// SetAdminID sets the "admin_id" field.
func (rauo *RoleAdminUpdateOne) SetAdminID(i int) *RoleAdminUpdateOne {
	rauo.mutation.ResetAdminID()
	rauo.mutation.SetAdminID(i)
	return rauo
}

// AddAdminID adds i to the "admin_id" field.
func (rauo *RoleAdminUpdateOne) AddAdminID(i int) *RoleAdminUpdateOne {
	rauo.mutation.AddAdminID(i)
	return rauo
}

// Mutation returns the RoleAdminMutation object of the builder.
func (rauo *RoleAdminUpdateOne) Mutation() *RoleAdminMutation {
	return rauo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *RoleAdminUpdateOne) Select(field string, fields ...string) *RoleAdminUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated RoleAdmin entity.
func (rauo *RoleAdminUpdateOne) Save(ctx context.Context) (*RoleAdmin, error) {
	var (
		err  error
		node *RoleAdmin
	)
	if len(rauo.hooks) == 0 {
		node, err = rauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleAdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rauo.mutation = mutation
			node, err = rauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rauo.hooks) - 1; i >= 0; i-- {
			if rauo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rauo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rauo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RoleAdmin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RoleAdminMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *RoleAdminUpdateOne) SaveX(ctx context.Context) *RoleAdmin {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *RoleAdminUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *RoleAdminUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rauo *RoleAdminUpdateOne) sqlSave(ctx context.Context) (_node *RoleAdmin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleadmin.Table,
			Columns: roleadmin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: roleadmin.FieldID,
			},
		},
	}
	id, ok := rauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoleAdmin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roleadmin.FieldID)
		for _, f := range fields {
			if !roleadmin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roleadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.RoleID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldRoleID,
		})
	}
	if value, ok := rauo.mutation.AddedRoleID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldRoleID,
		})
	}
	if value, ok := rauo.mutation.AdminID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldAdminID,
		})
	}
	if value, ok := rauo.mutation.AddedAdminID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldAdminID,
		})
	}
	_node = &RoleAdmin{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
