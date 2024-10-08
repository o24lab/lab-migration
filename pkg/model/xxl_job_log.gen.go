// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameXxlJobLog = "xxl_job_log"

// XxlJobLog mapped from table <xxl_job_log>
type XxlJobLog struct {
	ID                     int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	JobGroup               int32     `gorm:"column:job_group;not null;comment:执行器主键ID" json:"job_group"`                                // 执行器主键ID
	JobID                  int32     `gorm:"column:job_id;not null;comment:任务，主键ID" json:"job_id"`                                      // 任务，主键ID
	ExecutorAddress        string    `gorm:"column:executor_address;comment:执行器地址，本次执行的地址" json:"executor_address"`                     // 执行器地址，本次执行的地址
	ExecutorHandler        string    `gorm:"column:executor_handler;comment:执行器任务handler" json:"executor_handler"`                      // 执行器任务handler
	ExecutorParam          string    `gorm:"column:executor_param;comment:执行器任务参数" json:"executor_param"`                               // 执行器任务参数
	ExecutorShardingParam  string    `gorm:"column:executor_sharding_param;comment:执行器任务分片参数，格式如 1/2" json:"executor_sharding_param"`   // 执行器任务分片参数，格式如 1/2
	ExecutorFailRetryCount int32     `gorm:"column:executor_fail_retry_count;not null;comment:失败重试次数" json:"executor_fail_retry_count"` // 失败重试次数
	TriggerTime            time.Time `gorm:"column:trigger_time;comment:调度-时间" json:"trigger_time"`                                     // 调度-时间
	TriggerCode            int32     `gorm:"column:trigger_code;not null;comment:调度-结果" json:"trigger_code"`                            // 调度-结果
	TriggerMsg             string    `gorm:"column:trigger_msg;comment:调度-日志" json:"trigger_msg"`                                       // 调度-日志
	HandleTime             time.Time `gorm:"column:handle_time;comment:执行-时间" json:"handle_time"`                                       // 执行-时间
	HandleCode             int32     `gorm:"column:handle_code;not null;comment:执行-状态" json:"handle_code"`                              // 执行-状态
	HandleMsg              string    `gorm:"column:handle_msg;comment:执行-日志" json:"handle_msg"`                                         // 执行-日志
	AlarmStatus            int32     `gorm:"column:alarm_status;not null;comment:告警状态：0-默认、1-无需告警、2-告警成功、3-告警失败" json:"alarm_status"`   // 告警状态：0-默认、1-无需告警、2-告警成功、3-告警失败
}

// TableName XxlJobLog's table name
func (*XxlJobLog) TableName() string {
	return TableNameXxlJobLog
}
