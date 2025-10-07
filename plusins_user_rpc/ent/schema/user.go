package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "time"
)

// User holds the schema definition for the User entity.
type User struct {
    ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("username").NotEmpty().Unique(),
        field.String("password").NotEmpty(),
        field.String("email").Optional().Unique(),
        field.Time("created_at").Default(time.Now),
    }
}
