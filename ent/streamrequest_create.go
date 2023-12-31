// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skkugoon/strattonight/ent/streamrequest"
)

// StreamRequestCreate is the builder for creating a StreamRequest entity.
type StreamRequestCreate struct {
	config
	mutation *StreamRequestMutation
	hooks    []Hook
}

// SetRequestID sets the "request_id" field.
func (src *StreamRequestCreate) SetRequestID(i int) *StreamRequestCreate {
	src.mutation.SetRequestID(i)
	return src
}

// SetRequestType sets the "request_type" field.
func (src *StreamRequestCreate) SetRequestType(s string) *StreamRequestCreate {
	src.mutation.SetRequestType(s)
	return src
}

// SetIsActive sets the "is_active" field.
func (src *StreamRequestCreate) SetIsActive(b bool) *StreamRequestCreate {
	src.mutation.SetIsActive(b)
	return src
}

// SetID sets the "id" field.
func (src *StreamRequestCreate) SetID(i int) *StreamRequestCreate {
	src.mutation.SetID(i)
	return src
}

// Mutation returns the StreamRequestMutation object of the builder.
func (src *StreamRequestCreate) Mutation() *StreamRequestMutation {
	return src.mutation
}

// Save creates the StreamRequest in the database.
func (src *StreamRequestCreate) Save(ctx context.Context) (*StreamRequest, error) {
	return withHooks(ctx, src.sqlSave, src.mutation, src.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (src *StreamRequestCreate) SaveX(ctx context.Context) *StreamRequest {
	v, err := src.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (src *StreamRequestCreate) Exec(ctx context.Context) error {
	_, err := src.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (src *StreamRequestCreate) ExecX(ctx context.Context) {
	if err := src.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (src *StreamRequestCreate) check() error {
	if _, ok := src.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "StreamRequest.request_id"`)}
	}
	if _, ok := src.mutation.RequestType(); !ok {
		return &ValidationError{Name: "request_type", err: errors.New(`ent: missing required field "StreamRequest.request_type"`)}
	}
	if _, ok := src.mutation.IsActive(); !ok {
		return &ValidationError{Name: "is_active", err: errors.New(`ent: missing required field "StreamRequest.is_active"`)}
	}
	return nil
}

func (src *StreamRequestCreate) sqlSave(ctx context.Context) (*StreamRequest, error) {
	if err := src.check(); err != nil {
		return nil, err
	}
	_node, _spec := src.createSpec()
	if err := sqlgraph.CreateNode(ctx, src.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	src.mutation.id = &_node.ID
	src.mutation.done = true
	return _node, nil
}

func (src *StreamRequestCreate) createSpec() (*StreamRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &StreamRequest{config: src.config}
		_spec = sqlgraph.NewCreateSpec(streamrequest.Table, sqlgraph.NewFieldSpec(streamrequest.FieldID, field.TypeInt))
	)
	if id, ok := src.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := src.mutation.RequestID(); ok {
		_spec.SetField(streamrequest.FieldRequestID, field.TypeInt, value)
		_node.RequestID = value
	}
	if value, ok := src.mutation.RequestType(); ok {
		_spec.SetField(streamrequest.FieldRequestType, field.TypeString, value)
		_node.RequestType = value
	}
	if value, ok := src.mutation.IsActive(); ok {
		_spec.SetField(streamrequest.FieldIsActive, field.TypeBool, value)
		_node.IsActive = value
	}
	return _node, _spec
}

// StreamRequestCreateBulk is the builder for creating many StreamRequest entities in bulk.
type StreamRequestCreateBulk struct {
	config
	err      error
	builders []*StreamRequestCreate
}

// Save creates the StreamRequest entities in the database.
func (srcb *StreamRequestCreateBulk) Save(ctx context.Context) ([]*StreamRequest, error) {
	if srcb.err != nil {
		return nil, srcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(srcb.builders))
	nodes := make([]*StreamRequest, len(srcb.builders))
	mutators := make([]Mutator, len(srcb.builders))
	for i := range srcb.builders {
		func(i int, root context.Context) {
			builder := srcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StreamRequestMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, srcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, srcb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, srcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (srcb *StreamRequestCreateBulk) SaveX(ctx context.Context) []*StreamRequest {
	v, err := srcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (srcb *StreamRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := srcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srcb *StreamRequestCreateBulk) ExecX(ctx context.Context) {
	if err := srcb.Exec(ctx); err != nil {
		panic(err)
	}
}
