// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActivityBlackInfo = "activity_black_info"

// ActivityBlackInfo 活动黑名单信息表
type ActivityBlackInfo struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:规则ID，自增主键" json:"id"`                           // 规则ID，自增主键
	BlackType   int32  `gorm:"column:black_type;not null;comment:黑名单用户类型 (0: 用户名, 1: IP, 2: 真实姓名, 3: 银行卡)" json:"black_type"` // 黑名单用户类型 (0: 用户名, 1: IP, 2: 真实姓名, 3: 银行卡)
	BlackValue  string `gorm:"column:black_value;not null;comment:黑名单用户的值（例如，用户名、IP地址、真实姓名、银行卡号）" json:"black_value"`         // 黑名单用户的值（例如，用户名、IP地址、真实姓名、银行卡号）
	BlackStatus int32  `gorm:"column:black_status;not null;default:1;comment:状态：1-启用,2-禁用" json:"black_status"`               // 状态：1-启用,2-禁用
	CreateAt    int32  `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`                                       // 创建时间
	UpdateAt    int32  `gorm:"column:update_at;not null;comment:修改时间" json:"update_at"`                                       // 修改时间
	OpUser      string `gorm:"column:op_user;not null;default:system;comment:操作人" json:"op_user"`                             // 操作人
}

// TableName ActivityBlackInfo's table name
func (*ActivityBlackInfo) TableName() string {
	return TableNameActivityBlackInfo
}
