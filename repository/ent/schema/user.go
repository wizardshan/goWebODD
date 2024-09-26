package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("hash_id").Default(""),
		field.String("mobile").Default(""),
		field.String("password").Sensitive(),
		field.Int("age").Default(0),
		field.Int("level").Default(0),
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
