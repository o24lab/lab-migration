// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinConfig = "win_config"

// WinConfig 系统配置
type WinConfig struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string `gorm:"column:title;not null;comment:名字" json:"title"`                                       // 名字
	TitleZh   string `gorm:"column:title_zh;not null;comment:名字中文" json:"title_zh"`                               // 名字中文
	Value     string `gorm:"column:value;not null;comment:字段值" json:"value"`                                      // 字段值
	ShowApp   int32  `gorm:"column:show_app;not null;default:1;comment:类型:0-全部 1-WEB 2-后台 3-不显示" json:"show_app"` // 类型:0-全部 1-WEB 2-后台 3-不显示
	CanModify int32  `gorm:"column:can_modify;not null;default:1;comment:支持修改:1-支持 0-不支持" json:"can_modify"`      // 支持修改:1-支持 0-不支持
	Status    int32  `gorm:"column:status;not null;default:1;comment:是否启用:1-启用 0-不启用" json:"status"`              // 是否启用:1-启用 0-不启用
	CreatedAt int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int32  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinConfig's table name
func (*WinConfig) TableName() string {
	return TableNameWinConfig
}
