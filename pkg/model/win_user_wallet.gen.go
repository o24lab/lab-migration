// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserWallet = "win_user_wallet"

// WinUserWallet 用户钱包表
type WinUserWallet struct {
	ID         int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UID        int32   `gorm:"column:uid;not null;comment:用户id" json:"uid"`                                 // 用户id
	Username   string  `gorm:"column:username;not null;comment:用户名" json:"username"`                        // 用户名
	MerchantID int32   `gorm:"column:merchant_id;not null;comment:商户id" json:"merchant_id"`                 // 商户id
	Category   int32   `gorm:"column:category;not null;default:1;comment:钱包类型:支付/游戏/活动/佣金" json:"category"` // 钱包类型:支付/游戏/活动/佣金
	Currency   int32   `gorm:"column:currency;not null;default:1;comment:币种" json:"currency"`               // 币种
	Coin       float64 `gorm:"column:coin;not null;default:0.0000;comment:账户余额" json:"coin"`                // 账户余额
	Frozen     float64 `gorm:"column:frozen;not null;default:0.0000;comment:冻结金额" json:"frozen"`            // 冻结金额
	Version    int32   `gorm:"column:version;default:1;comment:版本号" json:"version"`                         // 版本号
	ModifyAt   int64   `gorm:"column:modify_at;not null;comment:13位时间戳" json:"modify_at"`                   // 13位时间戳
	CreatedAt  int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  int32   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinUserWallet's table name
func (*WinUserWallet) TableName() string {
	return TableNameWinUserWallet
}
