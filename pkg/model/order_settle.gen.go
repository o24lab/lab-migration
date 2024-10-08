// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOrderSettle = "order_settle"

// OrderSettle 结算分批处理表
type OrderSettle struct {
	ID          int64     `gorm:"column:id;primaryKey;comment:主键" json:"id"`                                              // 主键
	OperationID string    `gorm:"column:operation_id;not null;comment:交易id" json:"operation_id"`                          // 交易id
	UserID      int64     `gorm:"column:user_id;not null;comment:用户ID" json:"user_id"`                                    // 用户ID
	OrderType   int32     `gorm:"column:order_type;not null;default:1;comment:注单类型:1-结算,2-重新结算,3-取消结算" json:"order_type"` // 注单类型:1-结算,2-重新结算,3-取消结算
	ReqJSON     string    `gorm:"column:req_json;not null;comment:请求注单信息" json:"req_json"`                                // 请求注单信息
	RetryCount  int32     `gorm:"column:retry_count;not null;comment:重试次数" json:"retry_count"`                            // 重试次数
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`    // 创建时间
}

// TableName OrderSettle's table name
func (*OrderSettle) TableName() string {
	return TableNameOrderSettle
}
