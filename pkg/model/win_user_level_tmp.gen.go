// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinUserLevelTmp = "win_user_level_tmp"

// WinUserLevelTmp 会员等级
type WinUserLevelTmp struct {
	ID               int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Code             string `gorm:"column:code;not null;comment:会员等级" json:"code"`                           // 会员等级
	Name             string `gorm:"column:name;not null;comment:等级名称" json:"name"`                           // 等级名称
	Icon             string `gorm:"column:icon;not null;comment:ICON" json:"icon"`                           // ICON
	DepositChannel   string `gorm:"column:deposit_channel;comment:存款通道" json:"deposit_channel"`              // 存款通道
	ScoreUpgradeRate int32  `gorm:"column:score_upgrade_rate;comment:经验率" json:"score_upgrade_rate"`         // 经验率
	ScoreUpgradeMax  int32  `gorm:"column:score_upgrade_max;comment:最大经验值" json:"score_upgrade_max"`         // 最大经验值
	ScoreUpgradeMin  int32  `gorm:"column:score_upgrade_min;comment:最小经验值" json:"score_upgrade_min"`         // 最小经验值
	ScoreRelegation  int32  `gorm:"column:score_relegation;not null;comment:保级有效投注" json:"score_relegation"` // 保级有效投注
	RelegationDay    int32  `gorm:"column:relegation_day;not null;comment:保级有效天数" json:"relegation_day"`     // 保级有效天数
	BetSum           int32  `gorm:"column:bet_sum;not null;comment:累计投注" json:"bet_sum"`                     // 累计投注
	DepositSum       int32  `gorm:"column:deposit_sum;not null;comment:累计存款" json:"deposit_sum"`             // 累计存款
	WithdrawalCount  int32  `gorm:"column:withdrawal_count;not null;comment:每日提款次数" json:"withdrawal_count"` // 每日提款次数
	WithdrawalCoin   int32  `gorm:"column:withdrawal_coin;not null;comment:每日提款额度" json:"withdrawal_coin"`   // 每日提款额度
	RebateMax        int32  `gorm:"column:rebate_max;not null;comment:每日返水上限" json:"rebate_max"`             // 每日返水上限
	UpgradeReward    int32  `gorm:"column:upgrade_reward;not null;comment:升级礼金" json:"upgrade_reward"`       // 升级礼金
	BirthdayReward   int32  `gorm:"column:birthday_reward;not null;comment:生日礼金" json:"birthday_reward"`     // 生日礼金
	WeekReward       int32  `gorm:"column:week_reward;not null;comment:周礼金" json:"week_reward"`              // 周礼金
	MonthReward      int32  `gorm:"column:month_reward;not null;comment:月礼金" json:"month_reward"`            // 月礼金
	FlowClaim        int32  `gorm:"column:flow_claim;comment:打码倍数" json:"flow_claim"`                        // 打码倍数
	CreatedAt        int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt        int32  `gorm:"column:updated_at;not null" json:"updated_at"`
	UpdatedUser      string `gorm:"column:updated_user;comment:最后修改人" json:"updated_user"` // 最后修改人
}

// TableName WinUserLevelTmp's table name
func (*WinUserLevelTmp) TableName() string {
	return TableNameWinUserLevelTmp
}
