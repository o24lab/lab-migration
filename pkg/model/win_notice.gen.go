// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinNotice = "win_notice"

// WinNotice 公告信息
type WinNotice struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增ID" json:"id"`          // 自增ID
	Title        string `gorm:"column:title;not null;comment:标题" json:"title"`                           // 标题
	Content      string `gorm:"column:content;comment:内容" json:"content"`                                // 内容
	Category     int32  `gorm:"column:category;default:4;comment:类型:1-系统公告2-站内信 3-系统消息" json:"category"` // 类型:1-系统公告2-站内信 3-系统消息
	Status       int32  `gorm:"column:status;not null;default:1;comment:状态:1-启用 0-停用" json:"status"`     // 状态:1-启用 0-停用
	Sort         int32  `gorm:"column:sort;not null;default:1;comment:排序:从大到小" json:"sort"`              // 排序:从大到小
	CreatedAt    int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    int32  `gorm:"column:updated_at;not null" json:"updated_at"`
	OperatorName string `gorm:"column:operator_name;comment:操作人姓名" json:"operator_name"` // 操作人姓名
	UID          int32  `gorm:"column:uid;not null;comment:用户ID" json:"uid"`             // 用户ID
}

// TableName WinNotice's table name
func (*WinNotice) TableName() string {
	return TableNameWinNotice
}
