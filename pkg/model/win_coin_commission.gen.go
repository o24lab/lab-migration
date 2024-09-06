// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinCoinCommission = "win_coin_commission"

// WinCoinCommission 佣金表
type WinCoinCommission struct {
	ID             int64   `gorm:"column:id;primaryKey" json:"id"`
	UID            int32   `gorm:"column:uid;not null;comment:代理UID" json:"uid"`                            // 代理UID
	Username       string  `gorm:"column:username;not null;comment:用户名" json:"username"`                    // 用户名
	AgentLevel     int32   `gorm:"column:agent_level;not null;comment:代理层级" json:"agent_level"`             // 代理层级
	Riqi           int32   `gorm:"column:riqi;not null;comment:佣金时间" json:"riqi"`                           // 佣金时间
	Coin           float64 `gorm:"column:coin;not null;comment:佣金金额" json:"coin"`                           // 佣金金额
	SubUID         int32   `gorm:"column:sub_uid;not null;comment:下级UID" json:"sub_uid"`                    // 下级UID
	SubUsername    string  `gorm:"column:sub_username;not null;comment:下级用户名" json:"sub_username"`          // 下级用户名
	SubBetTrunover float64 `gorm:"column:sub_bet_trunover;not null;comment:下级流水总额" json:"sub_bet_trunover"` // 下级流水总额
	Rate           float64 `gorm:"column:rate;not null;comment:佣金比例" json:"rate"`                           // 佣金比例
	CoinBefore     float64 `gorm:"column:coin_before;comment:即时余额" json:"coin_before"`                      // 即时余额
	Status         int32   `gorm:"column:status;not null;default:1;comment:状态:0-未发放 1-已发放" json:"status"`   // 状态:0-未发放 1-已发放
	CreatedAt      int32   `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt      int32   `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinCoinCommission's table name
func (*WinCoinCommission) TableName() string {
	return TableNameWinCoinCommission
}
