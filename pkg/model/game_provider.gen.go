// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGameProvider = "game_provider"

// GameProvider 游戏供应商列表
type GameProvider struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code         string `gorm:"column:code;not null;comment:平台编码" json:"code"` // 平台编码
	Name         string `gorm:"column:name;not null;comment:平台名称" json:"name"` // 平台名称
	Icon         string `gorm:"column:icon" json:"icon"`
	Config       string `gorm:"column:config;comment:配置信息" json:"config"`                                 // 配置信息
	Sort         int32  `gorm:"column:sort;not null;comment:排序: 从低到高" json:"sort"`                        // 排序: 从低到高
	ConfigStatus int32  `gorm:"column:config_status;not null;comment:状态: 1-启用 0-停用" json:"config_status"` // 状态: 1-启用 0-停用
	CreatedAt    int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    int32  `gorm:"column:updated_at;not null" json:"updated_at"`
	CreatedBy    string `gorm:"column:created_by;comment:操作人姓名" json:"created_by"` // 操作人姓名
	UpdatedBy    string `gorm:"column:updated_by;comment:最后更新人" json:"updated_by"` // 最后更新人
}

// TableName GameProvider's table name
func (*GameProvider) TableName() string {
	return TableNameGameProvider
}
