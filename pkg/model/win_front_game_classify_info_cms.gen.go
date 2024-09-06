// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinFrontGameClassifyInfoCm = "win_front_game_classify_info_cms"

// WinFrontGameClassifyInfoCm 首页游戏分类页面配置详情表
type WinFrontGameClassifyInfoCm struct {
	ID          int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CID         int32  `gorm:"column:c_id;not null;comment:分类id" json:"c_id"`                     // 分类id
	SlotID      string `gorm:"column:slot_id;not null;default:0;comment:slotId" json:"slot_id"`   // slotId
	GameID      int32  `gorm:"column:game_id;not null;comment:游戏ID(关联game_list)" json:"game_id"`  // 游戏ID(关联game_list)
	GameGroupID int32  `gorm:"column:game_group_id;not null;comment:游戏大类类型" json:"game_group_id"` // 游戏大类类型
	PlatID      int32  `gorm:"column:plat_id;not null;comment:游戏平台id" json:"plat_id"`             // 游戏平台id
	Sort        int32  `gorm:"column:sort;default:99;comment:排序" json:"sort"`                     // 排序
	CreatedAt   int64  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`         // 创建时间
	CreateUser  string `gorm:"column:create_user;not null;comment:创建人" json:"create_user"`        // 创建人
	UpdatedAt   int64  `gorm:"column:updated_at;comment:修改人" json:"updated_at"`                   // 修改人
	UpdateUser  string `gorm:"column:update_user;comment:修改人" json:"update_user"`                 // 修改人
}

// TableName WinFrontGameClassifyInfoCm's table name
func (*WinFrontGameClassifyInfoCm) TableName() string {
	return TableNameWinFrontGameClassifyInfoCm
}
