// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinCoinDepositRecord = "win_coin_deposit_record"

// WinCoinDepositRecord 充值记录表
type WinCoinDepositRecord struct {
	ID               int64   `gorm:"column:id;primaryKey" json:"id"`
	OrderID          string  `gorm:"column:order_id;not null;comment:订单号(三方平台用)" json:"order_id"`                                                              // 订单号(三方平台用)
	PlatOrderID      string  `gorm:"column:plat_order_id;not null;comment:三方平台订单号" json:"plat_order_id"`                                                       // 三方平台订单号
	UID              int32   `gorm:"column:uid;not null;comment:UID" json:"uid"`                                                                               // UID
	Username         string  `gorm:"column:username;not null;comment:用户名" json:"username"`                                                                     // 用户名
	MerchantID       int32   `gorm:"column:merchant_id;not null;comment:商户id" json:"merchant_id"`                                                              // 商户id
	Code             string  `gorm:"column:code;not null;comment:支付通道编码" json:"code"`                                                                          // 支付通道编码
	PlatName         string  `gorm:"column:plat_name;not null;comment:平台名称" json:"plat_name"`                                                                  // 平台名称
	PlatNickName     string  `gorm:"column:plat_nick_name;comment:平台自定义名称" json:"plat_nick_name"`                                                              // 平台自定义名称
	CoinBefore       float64 `gorm:"column:coin_before;not null;comment:充值前金额" json:"coin_before"`                                                             // 充值前金额
	PayAddress       string  `gorm:"column:pay_address;comment:加密地址" json:"pay_address"`                                                                       // 加密地址
	PayAmount        float64 `gorm:"column:pay_amount;not null;default:0.0000;comment:充值金额" json:"pay_amount"`                                                 // 充值金额
	ExchangeRate     float64 `gorm:"column:exchange_rate;not null;default:0.0000;comment:汇率" json:"exchange_rate"`                                             // 汇率
	RealAmount       float64 `gorm:"column:real_amount;not null;default:0.0000;comment:到账金额" json:"real_amount"`                                               // 到账金额
	Currency         string  `gorm:"column:currency;not null;comment:币种" json:"currency"`                                                                      // 币种
	DepStatus        int32   `gorm:"column:dep_status;not null;default:9;comment:充值标识:1-首充 2-二充 9-其他" json:"dep_status"`                                       // 充值标识:1-首充 2-二充 9-其他
	Category         int32   `gorm:"column:category;not null;comment:类型:0-钱包充值 1-佣金钱包转账充值" json:"category"`                                                    // 类型:0-钱包充值 1-佣金钱包转账充值
	CategoryCurrency int32   `gorm:"column:category_currency;not null;default:1;comment:货币类型:0-数字货币 1-法币" json:"category_currency"`                            // 货币类型:0-数字货币 1-法币
	CategoryTransfer int32   `gorm:"column:category_transfer;not null;default:3;comment:转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH" json:"category_transfer"` // 转账类型:1-TRC-20 2-ERC-20 3-BANK 4-PIX 5-GCASH
	AdminUID         int32   `gorm:"column:admin_uid;not null;comment:审核ID" json:"admin_uid"`                                                                  // 审核ID
	Mark             string  `gorm:"column:mark;not null;comment:备注" json:"mark"`                                                                              // 备注
	Status           int32   `gorm:"column:status;not null;default:1;comment:状态: 0-申请中 1-成功 2-失败" json:"status"`                                               // 状态: 0-申请中 1-成功 2-失败
	CreatedAt        int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        int32   `gorm:"column:updated_at;not null" json:"updated_at"`
	Track            int32   `gorm:"column:track;not null;comment:埋点状态: 0-未埋点 1-埋点完成 2-埋点失败" json:"track"` // 埋点状态: 0-未埋点 1-埋点完成 2-埋点失败
	ActivityID       int64   `gorm:"column:activity_id;not null;comment:活动ID" json:"activity_id"`          // 活动ID
}

// TableName WinCoinDepositRecord's table name
func (*WinCoinDepositRecord) TableName() string {
	return TableNameWinCoinDepositRecord
}
