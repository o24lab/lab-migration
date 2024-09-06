// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinCoinAdminTransfer = "win_coin_admin_transfer"

// WinCoinAdminTransfer 后台调账记录
type WinCoinAdminTransfer struct {
	ID         int64   `gorm:"column:id;primaryKey" json:"id"`
	AdminID    int32   `gorm:"column:admin_id;not null;comment:管理员ID" json:"admin_id"`                            // 管理员ID
	Coin       float64 `gorm:"column:coin;not null;default:0.0000;comment:调账金额" json:"coin"`                      // 调账金额
	CoinBefore float64 `gorm:"column:coin_before;not null;default:0.0000;comment:调账前金额" json:"coin_before"`       // 调账前金额
	UID        int32   `gorm:"column:uid;not null;comment:用户ID" json:"uid"`                                       // 用户ID
	Username   string  `gorm:"column:username;not null;comment:用户名" json:"username"`                              // 用户名
	Category   int32   `gorm:"column:category;not null;comment:调账原因:0-其他 1-误存调账 2-活动调账 3-用户强制提款" json:"category"` // 调账原因:0-其他 1-误存调账 2-活动调账 3-用户强制提款
	Mark       string  `gorm:"column:mark;not null;comment:调账原因" json:"mark"`                                     // 调账原因
	CreatedAt  int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt  int32   `gorm:"column:updated_at;not null" json:"updated_at"`
	FlowClaim  int32   `gorm:"column:flow_claim;not null;default:1;comment:流水倍数" json:"flow_claim"` // 流水倍数
	Message    string  `gorm:"column:message;comment:通知客户信息" json:"message"`                        // 通知客户信息
	MerchantID int32   `gorm:"column:merchant_id;not null;comment:商户id" json:"merchant_id"`         // 商户id
}

// TableName WinCoinAdminTransfer's table name
func (*WinCoinAdminTransfer) TableName() string {
	return TableNameWinCoinAdminTransfer
}
