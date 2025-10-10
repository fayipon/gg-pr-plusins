package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type Merchant struct {
    ent.Schema
}

func (Merchant) Fields() []ent.Field {
    return []ent.Field{
        field.String("merchant_code").Unique(),
        field.String("name"),
        field.String("contact_email"),
        field.String("contact_phone").Optional(),
        field.String("domain").Optional(),
        field.Int64("plan_id").Optional(),
        field.String("status").Default("active"),
    }
}
