// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinAgentLink = "win_agent_links"

// WinAgentLink 代理专属域名
type WinAgentLink struct {
	ID        int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Link      string `gorm:"column:link;not null;comment:推广域名" json:"link"`                       // 推广域名
	UID       int32  `gorm:"column:uid;not null;comment:代理ID" json:"uid"`                         // 代理ID
	Status    int32  `gorm:"column:status;not null;default:1;comment:状态:1-启用 0-停用" json:"status"` // 状态:1-启用 0-停用
	CreatedAt int32  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt int32  `gorm:"column:updated_at;not null" json:"updated_at"`
}

// TableName WinAgentLink's table name
func (*WinAgentLink) TableName() string {
	return TableNameWinAgentLink
}
