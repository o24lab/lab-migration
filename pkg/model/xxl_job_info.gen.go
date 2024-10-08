// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameXxlJobInfo = "xxl_job_info"

// XxlJobInfo mapped from table <xxl_job_info>
type XxlJobInfo struct {
	ID                     int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	JobGroup               int32     `gorm:"column:job_group;not null;comment:执行器主键ID" json:"job_group"` // 执行器主键ID
	JobDesc                string    `gorm:"column:job_desc;not null" json:"job_desc"`
	AddTime                time.Time `gorm:"column:add_time" json:"add_time"`
	UpdateTime             time.Time `gorm:"column:update_time" json:"update_time"`
	Author                 string    `gorm:"column:author;comment:作者" json:"author"`                                                     // 作者
	AlarmEmail             string    `gorm:"column:alarm_email;comment:报警邮件" json:"alarm_email"`                                         // 报警邮件
	ScheduleType           string    `gorm:"column:schedule_type;not null;default:NONE;comment:调度类型" json:"schedule_type"`               // 调度类型
	ScheduleConf           string    `gorm:"column:schedule_conf;comment:调度配置，值含义取决于调度类型" json:"schedule_conf"`                          // 调度配置，值含义取决于调度类型
	MisfireStrategy        string    `gorm:"column:misfire_strategy;not null;default:DO_NOTHING;comment:调度过期策略" json:"misfire_strategy"` // 调度过期策略
	ExecutorRouteStrategy  string    `gorm:"column:executor_route_strategy;comment:执行器路由策略" json:"executor_route_strategy"`              // 执行器路由策略
	ExecutorHandler        string    `gorm:"column:executor_handler;comment:执行器任务handler" json:"executor_handler"`                       // 执行器任务handler
	ExecutorParam          string    `gorm:"column:executor_param;comment:执行器任务参数" json:"executor_param"`                                // 执行器任务参数
	ExecutorBlockStrategy  string    `gorm:"column:executor_block_strategy;comment:阻塞处理策略" json:"executor_block_strategy"`               // 阻塞处理策略
	ExecutorTimeout        int32     `gorm:"column:executor_timeout;not null;comment:任务执行超时时间，单位秒" json:"executor_timeout"`              // 任务执行超时时间，单位秒
	ExecutorFailRetryCount int32     `gorm:"column:executor_fail_retry_count;not null;comment:失败重试次数" json:"executor_fail_retry_count"`  // 失败重试次数
	GlueType               string    `gorm:"column:glue_type;not null;comment:GLUE类型" json:"glue_type"`                                  // GLUE类型
	GlueSource             string    `gorm:"column:glue_source;comment:GLUE源代码" json:"glue_source"`                                      // GLUE源代码
	GlueRemark             string    `gorm:"column:glue_remark;comment:GLUE备注" json:"glue_remark"`                                       // GLUE备注
	GlueUpdatetime         time.Time `gorm:"column:glue_updatetime;comment:GLUE更新时间" json:"glue_updatetime"`                             // GLUE更新时间
	ChildJobid             string    `gorm:"column:child_jobid;comment:子任务ID，多个逗号分隔" json:"child_jobid"`                                 // 子任务ID，多个逗号分隔
	TriggerStatus          int32     `gorm:"column:trigger_status;not null;comment:调度状态：0-停止，1-运行" json:"trigger_status"`                // 调度状态：0-停止，1-运行
	TriggerLastTime        int64     `gorm:"column:trigger_last_time;not null;comment:上次调度时间" json:"trigger_last_time"`                  // 上次调度时间
	TriggerNextTime        int64     `gorm:"column:trigger_next_time;not null;comment:下次调度时间" json:"trigger_next_time"`                  // 下次调度时间
}

// TableName XxlJobInfo's table name
func (*XxlJobInfo) TableName() string {
	return TableNameXxlJobInfo
}
