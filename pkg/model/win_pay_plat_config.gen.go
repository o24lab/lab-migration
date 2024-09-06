// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinPayPlatConfig = "win_pay_plat_config"

// WinPayPlatConfig 三方支付配置信息
type WinPayPlatConfig struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PlatName     string `gorm:"column:plat_name;not null;comment:平台名称" json:"plat_name"`                  // 平台名称
	PlatNickName string `gorm:"column:plat_nick_name;comment:平台昵称" json:"plat_nick_name"`                 // 平台昵称
	MerchantID   string `gorm:"column:merchant_id;not null;comment:商户号" json:"merchant_id"`               // 商户号
	APIKey       string `gorm:"column:api_key;not null;comment:商户秘钥" json:"api_key"`                      // 商户秘钥
	Status       int32  `gorm:"column:status;not null;default:1;comment:状态:0-停用 1-启用 2-删除" json:"status"` // 状态:0-停用 1-启用 2-删除
	PlatConfig   string `gorm:"column:plat_config;comment:平台特殊配置" json:"plat_config"`                     // 平台特殊配置
	CreatedAt    int32  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                // 创建时间
	UpdatedAt    int32  `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`                // 修改时间
	Operator     string `gorm:"column:operator;comment:操作人" json:"operator"`                              // 操作人
	OperatorName string `gorm:"column:operator_name;comment:操作人姓名" json:"operator_name"`                  // 操作人姓名
}

// TableName WinPayPlatConfig's table name
func (*WinPayPlatConfig) TableName() string {
	return TableNameWinPayPlatConfig
}
