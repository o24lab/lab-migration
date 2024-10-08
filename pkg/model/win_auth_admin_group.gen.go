// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinAuthAdminGroup = "win_auth_admin_group"

// WinAuthAdminGroup 用户组表
type WinAuthAdminGroup struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Title     string `gorm:"column:title;not null;comment:用户组名" json:"title"`                      // 用户组名
	Parent    int32  `gorm:"column:parent;not null;comment:创建人" json:"parent"`                     // 创建人
	Pid       int32  `gorm:"column:pid;not null;comment:父ID" json:"pid"`                           // 父ID
	Rules     string `gorm:"column:rules;comment:菜单ID集合" json:"rules"`                             // 菜单ID集合
	OperateID int32  `gorm:"column:operate_id;not null;comment:用户组ID" json:"operate_id"`           // 用户组ID
	Status    int32  `gorm:"column:status;not null;default:1;comment:状态: 1-启用 0-冻结" json:"status"` // 状态: 1-启用 0-冻结
	CreatedAt int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int32  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinAuthAdminGroup's table name
func (*WinAuthAdminGroup) TableName() string {
	return TableNameWinAuthAdminGroup
}
