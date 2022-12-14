// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"vue-admin/app/rbac/service/internal/data/ent/roleadmin"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RoleAdminCreate is the builder for creating a RoleAdmin entity.
type RoleAdminCreate struct {
	config
	mutation *RoleAdminMutation
	hooks    []Hook
}

// SetRoleID sets the "role_id" field.
func (rac *RoleAdminCreate) SetRoleID(i int) *RoleAdminCreate {
	rac.mutation.SetRoleID(i)
	return rac
}

// SetAdminID sets the "admin_id" field.
func (rac *RoleAdminCreate) SetAdminID(i int64) *RoleAdminCreate {
	rac.mutation.SetAdminID(i)
	return rac
}

// SetID sets the "id" field.
func (rac *RoleAdminCreate) SetID(i int64) *RoleAdminCreate {
	rac.mutation.SetID(i)
	return rac
}

// Mutation returns the RoleAdminMutation object of the builder.
func (rac *RoleAdminCreate) Mutation() *RoleAdminMutation {
	return rac.mutation
}

// Save creates the RoleAdmin in the database.
func (rac *RoleAdminCreate) Save(ctx context.Context) (*RoleAdmin, error) {
	var (
		err  error
		node *RoleAdmin
	)
	if len(rac.hooks) == 0 {
		if err = rac.check(); err != nil {
			return nil, err
		}
		node, err = rac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleAdminMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rac.check(); err != nil {
				return nil, err
			}
			rac.mutation = mutation
			if node, err = rac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rac.hooks) - 1; i >= 0; i-- {
			if rac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rac.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (rac *RoleAdminCreate) SaveX(ctx context.Context) *RoleAdmin {
	v, err := rac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rac *RoleAdminCreate) Exec(ctx context.Context) error {
	_, err := rac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rac *RoleAdminCreate) ExecX(ctx context.Context) {
	if err := rac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rac *RoleAdminCreate) check() error {
	if _, ok := rac.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role_id", err: errors.New(`ent: missing required field "RoleAdmin.role_id"`)}
	}
	if _, ok := rac.mutation.AdminID(); !ok {
		return &ValidationError{Name: "admin_id", err: errors.New(`ent: missing required field "RoleAdmin.admin_id"`)}
	}
	return nil
}

func (rac *RoleAdminCreate) sqlSave(ctx context.Context) (*RoleAdmin, error) {
	_node, _spec := rac.createSpec()
	if err := sqlgraph.CreateNode(ctx, rac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (rac *RoleAdminCreate) createSpec() (*RoleAdmin, *sqlgraph.CreateSpec) {
	var (
		_node = &RoleAdmin{config: rac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: roleadmin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: roleadmin.FieldID,
			},
		}
	)
	if id, ok := rac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rac.mutation.RoleID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: roleadmin.FieldRoleID,
		})
		_node.RoleID = value
	}
	if value, ok := rac.mutation.AdminID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: roleadmin.FieldAdminID,
		})
		_node.AdminID = value
	}
	return _node, _spec
}

// RoleAdminCreateBulk is the builder for creating many RoleAdmin entities in bulk.
type RoleAdminCreateBulk struct {
	config
	builders []*RoleAdminCreate
}

// Save creates the RoleAdmin entities in the database.
func (racb *RoleAdminCreateBulk) Save(ctx context.Context) ([]*RoleAdmin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(racb.builders))
	nodes := make([]*RoleAdmin, len(racb.builders))
	mutators := make([]Mutator, len(racb.builders))
	for i := range racb.builders {
		func(i int, root context.Context) {
			builder := racb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RoleAdminMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, racb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, racb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, racb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (racb *RoleAdminCreateBulk) SaveX(ctx context.Context) []*RoleAdmin {
	v, err := racb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (racb *RoleAdminCreateBulk) Exec(ctx context.Context) error {
	_, err := racb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (racb *RoleAdminCreateBulk) ExecX(ctx context.Context) {
	if err := racb.Exec(ctx); err != nil {
		panic(err)
	}
}
