// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinHelpType = "win_help_type"

// WinHelpType 帮助类型
type WinHelpType struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键编号" json:"id"` // 主键编号
	Language  string `gorm:"column:language;not null;comment:语言" json:"language"`            // 语言
	TypeName  string `gorm:"column:type_name;not null;comment:类型名称" json:"type_name"`        // 类型名称
	ImageURL  string `gorm:"column:image_url;comment:图片地址" json:"image_url"`                 // 图片地址
	Sort      int32  `gorm:"column:sort;comment:排序" json:"sort"`                             // 排序
	Status    int32  `gorm:"column:status;comment:状态:1-启用 0-停用" json:"status"`               // 状态:1-启用 0-停用
	CreateBy  string `gorm:"column:create_by;comment:创建者" json:"create_by"`                  // 创建者
	UpdateBy  string `gorm:"column:update_by;comment:更新人" json:"update_by"`                  // 更新人
	CreatedAt int32  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt int32  `gorm:"column:updated_at" json:"updated_at"`
}

// TableName WinHelpType's table name
func (*WinHelpType) TableName() string {
	return TableNameWinHelpType
}
