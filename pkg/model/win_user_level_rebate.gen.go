// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserLevelRebate = "win_user_level_rebate"

// WinUserLevelRebate 会员等级返水
type WinUserLevelRebate struct {
	ID          int32   `gorm:"column:id;primaryKey;autoIncrement:true;comment:ID主键" json:"id"`                                                       // ID主键
	LevelID     int32   `gorm:"column:level_id;not null;comment:会员等级ID" json:"level_id"`                                                              // 会员等级ID
	GroupID     int32   `gorm:"column:group_id;not null;comment:类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能 11-table game" json:"group_id"` // 类型:1-体育 2-电子 3-真人 4-捕鱼 5-棋牌 6-电竞 7-彩票 8-动物 9-快速 10-技能 11-table game
	RebateRate  float64 `gorm:"column:rebate_rate;not null;default:0.0000;comment:返水比例" json:"rebate_rate"`                                           // 返水比例
	Status      int32   `gorm:"column:status;not null;comment:状态:1-启用 0-停用" json:"status"`                                                            // 状态:1-启用 0-停用
	CreatedAt   int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   int32   `gorm:"column:updated_at;not null" json:"updated_at"`
	UpdatedUser string  `gorm:"column:updated_user;comment:最后修改人" json:"updated_user"` // 最后修改人
}

// TableName WinUserLevelRebate's table name
func (*WinUserLevelRebate) TableName() string {
	return TableNameWinUserLevelRebate
}
