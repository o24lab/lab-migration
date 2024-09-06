// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinSpecialRecordStakeGame = "win_special_record_stake_game"

// WinSpecialRecordStakeGame 打码量特殊算法游戏配置表
type WinSpecialRecordStakeGame struct {
	ID                    int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                                  // 自增主键
	GameTypeID            string `gorm:"column:game_type_id;not null;comment:游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔" json:"game_type_id"`                         // 游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔
	GameProviderSubtypeID string `gorm:"column:game_provider_subtype_id;not null;comment:游戏厂商：0-全部，默认字符串ex:id1,id2 逗号分隔" json:"game_provider_subtype_id"` // 游戏厂商：0-全部，默认字符串ex:id1,id2 逗号分隔
	GameListID            string `gorm:"column:game_list_id;comment:子游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔" json:"game_list_id"`                                 // 子游戏类型：0-全部，默认字符串ex:id1,id2 逗号分隔
	CreatedAt             int32  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                                       // 创建时间
	UpdatedAt             int32  `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`                                                       // 修改时间
	UpdatedUser           string `gorm:"column:updated_user;comment:最后修改人" json:"updated_user"`                                                           // 最后修改人
}

// TableName WinSpecialRecordStakeGame's table name
func (*WinSpecialRecordStakeGame) TableName() string {
	return TableNameWinSpecialRecordStakeGame
}
