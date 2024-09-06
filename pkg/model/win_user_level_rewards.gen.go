// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserLevelReward = "win_user_level_rewards"

// WinUserLevelReward vip奖金表
type WinUserLevelReward struct {
	ID          int64   `gorm:"column:id;primaryKey" json:"id"`
	UID         int32   `gorm:"column:uid;not null;comment:UID" json:"uid"`                                      // UID
	Username    string  `gorm:"column:username;not null;comment:用户名" json:"username"`                            // 用户名
	Coin        float64 `gorm:"column:coin;not null;default:0.0000;comment:彩金金额" json:"coin"`                    // 彩金金额
	CoinBefore  float64 `gorm:"column:coin_before;not null;default:0.0000;comment:即时金额" json:"coin_before"`      // 即时金额
	Category    int32   `gorm:"column:category;not null;comment:福利类型:0-晋级彩金 1-生日礼金 2-周彩金3 -月彩金" json:"category"` // 福利类型:0-晋级彩金 1-生日礼金 2-周彩金3 -月彩金
	ReceiveAt   int32   `gorm:"column:receive_at;comment:领取时间" json:"receive_at"`                                // 领取时间
	UserLevelID int32   `gorm:"column:user_level_id;comment:用户等级编号" json:"user_level_id"`                        // 用户等级编号
	FlowClaim   int32   `gorm:"column:flow_claim;comment:流水倍数" json:"flow_claim"`                                // 流水倍数
	Codes       float64 `gorm:"column:codes;comment:所需打码量" json:"codes"`                                         // 所需打码量
	EndedAt     int32   `gorm:"column:ended_at;comment:领取结束时间" json:"ended_at"`                                  // 领取结束时间
	Status      int32   `gorm:"column:status;not null;comment:状态:0-待领取 1-已领取 2-已过期" json:"status"`               // 状态:0-待领取 1-已领取 2-已过期
	CreatedAt   int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt   int32   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinUserLevelReward's table name
func (*WinUserLevelReward) TableName() string {
	return TableNameWinUserLevelReward
}
