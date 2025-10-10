package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PlusinsMerchant holds the schema definition for the plusins_merchant table.
type PlusinsMerchant struct {
	ent.Schema
}

// Fields of the PlusinsMerchant.
func (PlusinsMerchant) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("merchant_name").NotEmpty().Comment("商户名称"),
		field.String("merchant_token").Unique().NotEmpty().Comment("商户授权 Token"),
		field.String("merchant_setting").Optional().Comment("商户设置 JSON 字符串"),
		field.Int("status").Default(1).Comment("状态：1=启用 0=停用"),
		field.Time("created_at").Default(time.Now).Comment("创建时间"),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Comment("更新时间"),
	}
}
