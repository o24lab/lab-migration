// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinCoinRebateSet = "win_coin_rebate_set"

// WinCoinRebateSet 天将红包反水设置表
type WinCoinRebateSet struct {
	ID                    int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                         // 自增主键
	ValidAmountStart      float64 `gorm:"column:valid_amount_start;not null;default:0.0000;comment:有效流水范围开始" json:"valid_amount_start"`           // 有效流水范围开始
	ValidAmountEnd        float64 `gorm:"column:valid_amount_end;not null;default:0.0000;comment:有效流水范围结束" json:"valid_amount_end"`               // 有效流水范围结束
	RebatePercentageStart float64 `gorm:"column:rebate_percentage_start;not null;default:0.0000;comment:反水比例范围开始" json:"rebate_percentage_start"` // 反水比例范围开始
	RebatePercentageEnd   float64 `gorm:"column:rebate_percentage_end;not null;default:0.0000;comment:反水比例范围结束" json:"rebate_percentage_end"`     // 反水比例范围结束
	CreatedAt             int32   `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`                                              // 创建时间
	UpdatedAt             int32   `gorm:"column:updated_at;not null;comment:修改时间" json:"updated_at"`                                              // 修改时间
	UpdatedUser           string  `gorm:"column:updated_user;comment:最后修改人" json:"updated_user"`                                                  // 最后修改人
}

// TableName WinCoinRebateSet's table name
func (*WinCoinRebateSet) TableName() string {
	return TableNameWinCoinRebateSet
}
