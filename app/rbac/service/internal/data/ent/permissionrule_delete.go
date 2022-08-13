// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"vue-admin/app/rbac/service/internal/data/ent/permissionrule"
	"vue-admin/app/rbac/service/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PermissionRuleDelete is the builder for deleting a PermissionRule entity.
type PermissionRuleDelete struct {
	config
	hooks    []Hook
	mutation *PermissionRuleMutation
}

// Where appends a list predicates to the PermissionRuleDelete builder.
func (prd *PermissionRuleDelete) Where(ps ...predicate.PermissionRule) *PermissionRuleDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *PermissionRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(prd.hooks) == 0 {
		affected, err = prd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PermissionRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			prd.mutation = mutation
			affected, err = prd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(prd.hooks) - 1; i >= 0; i-- {
			if prd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = prd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, prd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *PermissionRuleDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *PermissionRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: permissionrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: permissionrule.FieldID,
			},
		},
	}
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PermissionRuleDeleteOne is the builder for deleting a single PermissionRule entity.
type PermissionRuleDeleteOne struct {
	prd *PermissionRuleDelete
}

// Exec executes the deletion query.
func (prdo *PermissionRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{permissionrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *PermissionRuleDeleteOne) ExecX(ctx context.Context) {
	prdo.prd.ExecX(ctx)
}