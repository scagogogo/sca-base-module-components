package models

import "time"

// ComponentDependency 用于表示组件之间的依赖关系
type ComponentDependency[ComponentDependencyEcosystem any] struct {

	// 组件的名字
	Name string `mapstructure:"name" json:"name,omitempty" yaml:"name" db:"name" gorm:"column:name;index:component_dependency,unique"`

	// 组件的版本
	Version string `mapstructure:"version" json:"version,omitempty" yaml:"version" db:"version" gorm:"column:version;index:component_dependency,unique"`

	// 被依赖的组件的名字
	DependencyName string `mapstructure:"dependency_name" json:"dependency_name,omitempty" yaml:"dependency_name" db:"dependency_name" gorm:"column:dependency_name;index:component_dependency,unique"`

	// 被依赖的组件的版本
	DependencyVersion string `mapstructure:"dependency_version" json:"dependency_version,omitempty" yaml:"dependency_version" db:"dependency_version" gorm:"column:dependency_version;index:component_dependency,unique"`

	// 依赖的scope，不同的语言可能会有不同的解释，这里暂时先不做统一抽象
	Scope string `mapstructure:"scope" json:"scope,omitempty" yaml:"scope" db:"scope" gorm:"column:scope;index:component_dependency,unique"`

	// 不同的包管理器可能有自己独特的组件依赖字段，都可以放到这里
	ComponentDependencyEcosystem ComponentDependencyEcosystem `mapstructure:"component_dependency_ecosystem" json:"component_dependency_ecosystem,omitempty" yaml:"component_dependency_ecosystem" db:"component_dependency_ecosystem" gorm:"column:component_dependency_ecosystem;serializer:json"`

	// 几个通用的时间戳
	CreateTime *time.Time `mapstructure:"create_time" json:"create_time,omitempty" yaml:"create_time" db:"create_time" gorm:"column:create_time"`
	UpdateTime *time.Time `mapstructure:"update_time" json:"update_time,omitempty" yaml:"update_time" db:"update_time" gorm:"column:update_time"`
	ChangeTime *time.Time `mapstructure:"change_time" json:"change_time,omitempty" yaml:"change_time" db:"change_time" gorm:"column:change_time"`
}

func (x *ComponentDependency[ComponentDependencyEcosystem]) TableName() string {
	return "component_dependencies"
}
