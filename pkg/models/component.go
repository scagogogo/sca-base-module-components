package models

import "time"

// Component 表示一个组件的一个版本信息
type Component[ComponentEcosystem any] struct {

	// 组件的名称
	Name string `mapstructure:"name" json:"name,omitempty" yaml:"name" db:"name" gorm:"column:name;index:component_name_value,unique"`

	// 组件的版本
	Version string `mapstructure:"version" json:"version,omitempty" yaml:"version" db:"version" gorm:"column:version;index:component_name_value,unique"`

	// 此组件此版本的license
	Licenses []string `mapstructure:"licenses" json:"licenses,omitempty" yaml:"licenses" db:"licenses" gorm:"column:licenses;serializer:json"`

	// 当前版本的sha1，反向关联时使用
	Sha1 string `mapstructure:"sha1" json:"sha1,omitempty" yaml:"sha1" db:"sha1" gorm:"column:sha1"`

	// 此版本的发布时间
	PublishTime time.Time `mapstructure:"publish_time" json:"publish_time,omitempty" yaml:"publish_time" db:"publish_time" gorm:"column:publish_time"`

	// 不同的包管理器可能会拥有的不同的字段，这里允许自行扩展
	ComponentEcosystem ComponentEcosystem `mapstructure:"component_ecosystem" json:"component_ecosystem,omitempty" yaml:"component_ecosystem" db:"component_ecosystem" gorm:"column:component_ecosystem;serializer:json"`

	// 几个通用的时间戳
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"column:create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"column:update_time"`
	ChangeTime *time.Time `mapstructure:"change_time" json:"change_time,omitempty" yaml:"change_time" db:"change_time" gorm:"column:change_time"`
}

func NewComponent[ComponentEcosystem any]() *Component[ComponentEcosystem] {
	return &Component[ComponentEcosystem]{}
}

func (x *Component[ComponentEcosystem]) TableName() string {
	return "components"
}

