// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWinCityItem = "win_city_item"

// WinCityItem 省级城市字典表
type WinCityItem struct {
	ID           int32  `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`      // id
	ProvinceCode string `gorm:"column:province_code;not null;comment:省份字典码" json:"province_code"`  // 省份字典码
	Nation       string `gorm:"column:nation;not null;comment:国家" json:"nation"`                   // 国家
	ProvinceName string `gorm:"column:province_name;not null;comment:省份name" json:"province_name"` // 省份name
	City         string `gorm:"column:city;not null;comment:城市" json:"city"`                       // 城市
	Remark       string `gorm:"column:remark;comment:备注" json:"remark"`                            // 备注
	Sort         int32  `gorm:"column:sort;not null;comment:排序:从高到低" json:"sort"`                  // 排序:从高到低
	ReferID      int32  `gorm:"column:refer_id;not null;comment:预设字段" json:"refer_id"`             // 预设字段
	CreatedAt    int32  `gorm:"column:created_at;not null;comment:创建时间" json:"created_at"`         // 创建时间
	UpdatedAt    int32  `gorm:"column:updated_at;not null;comment:更新时间" json:"updated_at"`         // 更新时间
}

// TableName WinCityItem's table name
func (*WinCityItem) TableName() string {
	return TableNameWinCityItem
}
