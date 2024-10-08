// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameActivityConditionChristma = "activity_condition_christmas"

// ActivityConditionChristma 圣诞活动领取条件表
type ActivityConditionChristma struct {
	ID                        int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:规则ID，自增主键" json:"id"`                                                   // 规则ID，自增主键
	ActivityID                int64   `gorm:"column:activity_id;not null;comment:活动ID" json:"activity_id"`                                                           // 活动ID
	CompletionPeriodHours     int32   `gorm:"column:completion_period_hours;not null;comment:礼盒完成周期（小时）" json:"completion_period_hours"`                             // 礼盒完成周期（小时）
	MaxClaimPerDay            int32   `gorm:"column:max_claim_per_day;not null;default:1;comment:礼盒可领取次数（每天）" json:"max_claim_per_day"`                              // 礼盒可领取次数（每天）
	FirstStage                int32   `gorm:"column:first_stage;not null;comment:第一阶段彩金配置" json:"first_stage"`                                                       // 第一阶段彩金配置
	SecondStage               int32   `gorm:"column:second_stage;not null;comment:第二阶段彩金配置" json:"second_stage"`                                                     // 第二阶段彩金配置
	ThirdStage                int32   `gorm:"column:third_stage;not null;comment:第三阶段彩金配置" json:"third_stage"`                                                       // 第三阶段彩金配置
	FourthStage               int32   `gorm:"column:fourth_stage;not null;comment:第四阶段彩金配置" json:"fourth_stage"`                                                     // 第四阶段彩金配置
	StageThreshold            int32   `gorm:"column:stage_threshold;not null;comment:各阶段在距离领取达到提示阈值" json:"stage_threshold"`                                         // 各阶段在距离领取达到提示阈值
	StageReminderMessages     string  `gorm:"column:stage_reminder_messages;not null;comment:各阶段在距离领取达到提示阈值提示消息" json:"stage_reminder_messages"`                     // 各阶段在距离领取达到提示阈值提示消息
	RegisterAmount            float64 `gorm:"column:register_amount;not null;default:0.0000;comment:注册成本设定" json:"register_amount"`                                  // 注册成本设定
	NewSupportPointsRangeFrom float64 `gorm:"column:new_support_points_range_from;not null;default:0.0000;comment:新用户助力点数范围起" json:"new_support_points_range_from"`  // 新用户助力点数范围起
	NewSupportPointsRangeTo   float64 `gorm:"column:new_support_points_range_to;not null;default:0.0000;comment:新用户助力点数范围止" json:"new_support_points_range_to"`      // 新用户助力点数范围止
	OldSupportPointsRangeFrom float64 `gorm:"column:old_support_points_range_from;not null;default:0.0000;comment:老用户助力点数范围起" json:"old_support_points_range_from"`  // 老用户助力点数范围起
	OldSupportPointsRangeTo   float64 `gorm:"column:old_support_points_range_to;not null;default:0.0000;comment:老用户助力点数范围止" json:"old_support_points_range_to"`      // 老用户助力点数范围止
	RequireScreenshot         int32   `gorm:"column:require_screenshot;not null;default:1;comment:是否需上传分享截图，0表示否，1表示是" json:"require_screenshot"`                    // 是否需上传分享截图，0表示否，1表示是
	DailyJackpotThreshold     float64 `gorm:"column:daily_jackpot_threshold;not null;default:0.0000;comment:单日彩金上限告警阈值" json:"daily_jackpot_threshold"`              // 单日彩金上限告警阈值
	TotalBudgetThreshold      float64 `gorm:"column:total_budget_threshold;not null;default:0.0000;comment:活动总预算告警阈值" json:"total_budget_threshold"`                 // 活动总预算告警阈值
	KycRequired               int32   `gorm:"column:kyc_required;not null;default:1;comment:助力玩家是否KYC认证，0表示否，1表示是" json:"kyc_required"`                              // 助力玩家是否KYC认证，0表示否，1表示是
	NewUserFirstSupportReward float64 `gorm:"column:new_user_first_support_reward;not null;default:0.0000;comment:新用户首次助力成功奖励" json:"new_user_first_support_reward"` // 新用户首次助力成功奖励
	DepositChannels           string  `gorm:"column:deposit_channels;not null;default:0;comment:关联存款渠道：0-全部" json:"deposit_channels"`                                // 关联存款渠道：0-全部
	ClaimMethod               int32   `gorm:"column:claim_method;not null;comment:领取方式，0表示自动领取，1表示手动领取" json:"claim_method"`                                         // 领取方式，0表示自动领取，1表示手动领取
	IsAgentActive             int32   `gorm:"column:is_agent_active;not null;comment:是否代理线活动，0表示否，1表示是" json:"is_agent_active"`                                      // 是否代理线活动，0表示否，1表示是
	SpecificAgents            int32   `gorm:"column:specific_agents;not null;comment:指定代理线:0表示否，1表示是" json:"specific_agents"`                                        // 指定代理线:0表示否，1表示是
	InviteCodes               string  `gorm:"column:invite_codes;not null;comment:指定代理邀请码,使用,好分割" json:"invite_codes"`                                               // 指定代理邀请码,使用,好分割
	CreateAt                  int32   `gorm:"column:create_at;not null;comment:创建时间" json:"create_at"`                                                               // 创建时间
	UpdateAt                  int32   `gorm:"column:update_at;not null;comment:修改时间" json:"update_at"`                                                               // 修改时间
	OpUser                    string  `gorm:"column:op_user;not null;default:system;comment:操作人" json:"op_user"`                                                     // 操作人
	IsFour                    int32   `gorm:"column:is_four;not null;comment:圣诞礼盒第四阶段，0否，1是" json:"is_four"`                                                         // 圣诞礼盒第四阶段，0否，1是
}

// TableName ActivityConditionChristma's table name
func (*ActivityConditionChristma) TableName() string {
	return TableNameActivityConditionChristma
}
