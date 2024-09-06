// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActivityCondition = "activity_condition"

// ActivityCondition 活动领取条件-谁能领
type ActivityCondition struct {
	ID                 int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:规则ID，自增主键" json:"id"`                                // 规则ID，自增主键
	ActivityID         int64   `gorm:"column:activity_id;not null;comment:活动ID" json:"activity_id"`                                        // 活动ID
	ApplicableVipLevel int32   `gorm:"column:applicable_vip_level;not null;comment:适用VIP等级：0-全部" json:"applicable_vip_level"`              // 适用VIP等级：0-全部
	MaxClaimPerPeriod  int32   `gorm:"column:max_claim_per_period;not null;comment:可领取次数，0表示不限制" json:"max_claim_per_period"`              // 可领取次数，0表示不限制
	ClaimIntervalUnit  int32   `gorm:"column:claim_interval_unit;not null;comment:领取次数间隔单位：1-天，2-周，3-月，4-年" json:"claim_interval_unit"`    // 领取次数间隔单位：1-天，2-周，3-月，4-年
	DepositChannels    string  `gorm:"column:deposit_channels;not null;default:0;comment:关联存款渠道：0-全部" json:"deposit_channels"`             // 关联存款渠道：0-全部
	KycRequired        int32   `gorm:"column:kyc_required;not null;default:1;comment:领取彩金是否KYC认证，0表示否，1表示是" json:"kyc_required"`           // 领取彩金是否KYC认证，0表示否，1表示是
	RequireScreenshot  int32   `gorm:"column:require_screenshot;not null;default:1;comment:是否需上传分享截图，0表示否，1表示是" json:"require_screenshot"` // 是否需上传分享截图，0表示否，1表示是
	ProfitAndLoss      float64 `gorm:"column:profit_and_loss;not null;default:0.0000;comment:盈亏设定：盈亏金额大于等于" json:"profit_and_loss"`        // 盈亏设定：盈亏金额大于等于
	ClaimMethod        int32   `gorm:"column:claim_method;not null;comment:领取方式，0表示自动领取，1表示手动领取" json:"claim_method"`                      // 领取方式，0表示自动领取，1表示手动领取
	IsAgentActive      int32   `gorm:"column:is_agent_active;not null;comment:是否代理线活动，0表示否，1表示是" json:"is_agent_active"`                   // 是否代理线活动，0表示否，1表示是
	SpecificAgents     int32   `gorm:"column:specific_agents;not null;comment:指定代理线:0表示否，1表示是" json:"specific_agents"`                     // 指定代理线:0表示否，1表示是
	InviteCodes        string  `gorm:"column:invite_codes;not null;comment:指定代理邀请码,使用,好分割" json:"invite_codes"`                            // 指定代理邀请码,使用,好分割
	CreateAt           int32   `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`                                            // 创建时间
	UpdateAt           int32   `gorm:"column:update_at;not null;comment:修改时间" json:"update_at"`                                            // 修改时间
	OpUser             string  `gorm:"column:op_user;not null;default:system;comment:操作人" json:"op_user"`                                  // 操作人
}

// TableName ActivityCondition's table name
func (*ActivityCondition) TableName() string {
	return TableNameActivityCondition
}
