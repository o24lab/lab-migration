// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGameProviderSubtype = "game_provider_subtype"

// GameProviderSubtype 厂商游戏类型表
type GameProviderSubtype struct {
	ID             int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code           string  `gorm:"column:code;not null;comment:平台编码" json:"code"`                                                                                                 // 平台编码
	Name           string  `gorm:"column:name;not null;comment:名称" json:"name"`                                                                                                   // 名称
	Icon           string  `gorm:"column:icon;not null;comment:图标" json:"icon"`                                                                                                   // 图标
	GameTypeID     int32   `gorm:"column:game_type_id;not null;default:1;comment:自定义游戏类型ID:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能 11-table game" json:"game_type_id"` // 自定义游戏类型ID:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能 11-table game
	GameProviderID int32   `gorm:"column:game_provider_id;not null;comment:游戏供应商列表ID" json:"game_provider_id"`                                                                    // 游戏供应商列表ID
	ProviderRate   float64 `gorm:"column:provider_rate;not null;default:0.0000;comment:厂商税收比例" json:"provider_rate"`                                                              // 厂商税收比例
	Maintenance    string  `gorm:"column:maintenance;comment:维护时间" json:"maintenance"`                                                                                            // 维护时间
	GameCount      int32   `gorm:"column:game_count;comment:分类游戏数量" json:"game_count"`                                                                                            // 分类游戏数量
	Remark         string  `gorm:"column:remark;not null;comment:备注" json:"remark"`                                                                                               // 备注
	Sort           int32   `gorm:"column:sort;not null;comment:排序: 从高到低" json:"sort"`                                                                                             // 排序: 从高到低
	Status         int32   `gorm:"column:status;not null;comment:状态: 1-启用 0-停用" json:"status"`                                                                                    // 状态: 1-启用 0-停用
	CreatedAt      int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      int32   `gorm:"column:updated_at;not null" json:"updated_at"`
	UpdatedBy      string  `gorm:"column:updated_by;comment:最后更新人" json:"updated_by"`                     // 最后更新人
	CreatedBy      string  `gorm:"column:created_by;comment:创建人" json:"created_by"`                       // 创建人
	GamePagcorID   string  `gorm:"column:game_pagcor_id;comment:pagcor游戏分类表ID" json:"game_pagcor_id"`     // pagcor游戏分类表ID
	PagcorRate     float64 `gorm:"column:pagcor_rate;default:0.0000;comment:pagcor税率" json:"pagcor_rate"` // pagcor税率
}

// TableName GameProviderSubtype's table name
func (*GameProviderSubtype) TableName() string {
	return TableNameGameProviderSubtype
}
