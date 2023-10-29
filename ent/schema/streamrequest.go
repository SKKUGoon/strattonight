package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StreamRequest holds the schema definition for the StreamRequest entity.
type StreamRequest struct {
	ent.Schema
}

// Fields of the StreamRequest.
func (StreamRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("request_id"),
		field.String("request_type"),
		field.Bool("is_active"),
	}
}

// Edges of the StreamRequest.
func (StreamRequest) Edges() []ent.Edge {
	return nil
}

func (StreamRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{
			Table: "stream_request",
		},
	}
}
