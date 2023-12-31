// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/skkugoon/strattonight/ent/predicate"
	"github.com/skkugoon/strattonight/ent/streamrequest"
)

// StreamRequestUpdate is the builder for updating StreamRequest entities.
type StreamRequestUpdate struct {
	config
	hooks    []Hook
	mutation *StreamRequestMutation
}

// Where appends a list predicates to the StreamRequestUpdate builder.
func (sru *StreamRequestUpdate) Where(ps ...predicate.StreamRequest) *StreamRequestUpdate {
	sru.mutation.Where(ps...)
	return sru
}

// SetRequestID sets the "request_id" field.
func (sru *StreamRequestUpdate) SetRequestID(i int) *StreamRequestUpdate {
	sru.mutation.ResetRequestID()
	sru.mutation.SetRequestID(i)
	return sru
}

// AddRequestID adds i to the "request_id" field.
func (sru *StreamRequestUpdate) AddRequestID(i int) *StreamRequestUpdate {
	sru.mutation.AddRequestID(i)
	return sru
}

// SetRequestType sets the "request_type" field.
func (sru *StreamRequestUpdate) SetRequestType(s string) *StreamRequestUpdate {
	sru.mutation.SetRequestType(s)
	return sru
}

// SetIsActive sets the "is_active" field.
func (sru *StreamRequestUpdate) SetIsActive(b bool) *StreamRequestUpdate {
	sru.mutation.SetIsActive(b)
	return sru
}

// Mutation returns the StreamRequestMutation object of the builder.
func (sru *StreamRequestUpdate) Mutation() *StreamRequestMutation {
	return sru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sru *StreamRequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sru.sqlSave, sru.mutation, sru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sru *StreamRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := sru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sru *StreamRequestUpdate) Exec(ctx context.Context) error {
	_, err := sru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sru *StreamRequestUpdate) ExecX(ctx context.Context) {
	if err := sru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sru *StreamRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(streamrequest.Table, streamrequest.Columns, sqlgraph.NewFieldSpec(streamrequest.FieldID, field.TypeInt))
	if ps := sru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sru.mutation.RequestID(); ok {
		_spec.SetField(streamrequest.FieldRequestID, field.TypeInt, value)
	}
	if value, ok := sru.mutation.AddedRequestID(); ok {
		_spec.AddField(streamrequest.FieldRequestID, field.TypeInt, value)
	}
	if value, ok := sru.mutation.RequestType(); ok {
		_spec.SetField(streamrequest.FieldRequestType, field.TypeString, value)
	}
	if value, ok := sru.mutation.IsActive(); ok {
		_spec.SetField(streamrequest.FieldIsActive, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{streamrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sru.mutation.done = true
	return n, nil
}

// StreamRequestUpdateOne is the builder for updating a single StreamRequest entity.
type StreamRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StreamRequestMutation
}

// SetRequestID sets the "request_id" field.
func (sruo *StreamRequestUpdateOne) SetRequestID(i int) *StreamRequestUpdateOne {
	sruo.mutation.ResetRequestID()
	sruo.mutation.SetRequestID(i)
	return sruo
}

// AddRequestID adds i to the "request_id" field.
func (sruo *StreamRequestUpdateOne) AddRequestID(i int) *StreamRequestUpdateOne {
	sruo.mutation.AddRequestID(i)
	return sruo
}

// SetRequestType sets the "request_type" field.
func (sruo *StreamRequestUpdateOne) SetRequestType(s string) *StreamRequestUpdateOne {
	sruo.mutation.SetRequestType(s)
	return sruo
}

// SetIsActive sets the "is_active" field.
func (sruo *StreamRequestUpdateOne) SetIsActive(b bool) *StreamRequestUpdateOne {
	sruo.mutation.SetIsActive(b)
	return sruo
}

// Mutation returns the StreamRequestMutation object of the builder.
func (sruo *StreamRequestUpdateOne) Mutation() *StreamRequestMutation {
	return sruo.mutation
}

// Where appends a list predicates to the StreamRequestUpdate builder.
func (sruo *StreamRequestUpdateOne) Where(ps ...predicate.StreamRequest) *StreamRequestUpdateOne {
	sruo.mutation.Where(ps...)
	return sruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sruo *StreamRequestUpdateOne) Select(field string, fields ...string) *StreamRequestUpdateOne {
	sruo.fields = append([]string{field}, fields...)
	return sruo
}

// Save executes the query and returns the updated StreamRequest entity.
func (sruo *StreamRequestUpdateOne) Save(ctx context.Context) (*StreamRequest, error) {
	return withHooks(ctx, sruo.sqlSave, sruo.mutation, sruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sruo *StreamRequestUpdateOne) SaveX(ctx context.Context) *StreamRequest {
	node, err := sruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sruo *StreamRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := sruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sruo *StreamRequestUpdateOne) ExecX(ctx context.Context) {
	if err := sruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sruo *StreamRequestUpdateOne) sqlSave(ctx context.Context) (_node *StreamRequest, err error) {
	_spec := sqlgraph.NewUpdateSpec(streamrequest.Table, streamrequest.Columns, sqlgraph.NewFieldSpec(streamrequest.FieldID, field.TypeInt))
	id, ok := sruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "StreamRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, streamrequest.FieldID)
		for _, f := range fields {
			if !streamrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != streamrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sruo.mutation.RequestID(); ok {
		_spec.SetField(streamrequest.FieldRequestID, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.AddedRequestID(); ok {
		_spec.AddField(streamrequest.FieldRequestID, field.TypeInt, value)
	}
	if value, ok := sruo.mutation.RequestType(); ok {
		_spec.SetField(streamrequest.FieldRequestType, field.TypeString, value)
	}
	if value, ok := sruo.mutation.IsActive(); ok {
		_spec.SetField(streamrequest.FieldIsActive, field.TypeBool, value)
	}
	_node = &StreamRequest{config: sruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{streamrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sruo.mutation.done = true
	return _node, nil
}
