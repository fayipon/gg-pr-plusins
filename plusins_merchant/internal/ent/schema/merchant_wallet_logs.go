package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
)

// PlusinsMerchantWalletLogs holds the schema definition for the plusins_merchant_wallet_logs table.
type PlusinsMerchantWalletLogs struct {
	ent.Schema
}

func (PlusinsMerchantWalletLogs) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("merchant_id").Comment("商户 ID"),
		field.Int("wallet_id").Comment("钱包 ID"),
		field.String("change_type").Comment("变动类型：deposit / withdraw / adjust 等"),
		field.Float("before_amount").Default(0).Comment("变动前余额"),
		field.Float("after_amount").Default(0).Comment("变动后余额"),
		field.Float("amount").Default(0).Comment("变动金额"),
		field.Int("status").Default(1).Comment("状态：1=有效 0=无效"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (PlusinsMerchantWalletLogs) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("merchant", PlusinsMerchant.Type).
			Ref("wallet_logs").
			Field("merchant_id").
			Unique().
			Required(),
		edge.From("wallet", PlusinsMerchantWallet.Type).
			Ref("logs").
			Field("wallet_id").
			Unique().
			Required(),
	}
}
