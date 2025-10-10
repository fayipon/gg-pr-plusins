package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// PlusinsMerchantWallet holds the schema definition for the plusins_merchant_wallet table.
type PlusinsMerchantWallet struct {
	ent.Schema
}

func (PlusinsMerchantWallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("merchant_id").Comment("商户 ID (外键)"),
		field.String("wallet_type").Comment("钱包类型，如 main / bonus / rebate"),
		field.Float("amount").Default(0).Comment("钱包余额"),
		field.Int("status").Default(1).Comment("状态：1=启用 0=停用"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (PlusinsMerchantWallet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", PlusinsMerchant.Type).
			Ref("wallets").
			Field("merchant_id").
			Unique().
			Required(),
	}
}
