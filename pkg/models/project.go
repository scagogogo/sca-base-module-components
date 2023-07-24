package models

// Project 表示一个项目，一个项目可能会有多个模块组成
type Project[ProjectEcosystem, ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem any] struct {

	// 项目本身也可以认为是一个组件的扩展
	*Component[ComponentEcosystem]

	// 此项目都包含哪些模块，不支持模块系统的项目数组长度就恒为1
	Modules map[string]*Module[ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem]

	// 不同的包管理器可能会有不同的项目属性，允许自行扩展
	ProjectEcosystem ProjectEcosystem
}

// TakeFirstModule 获取项目的第一个Module，方便在不支持模块的项目级别贯通project和component
func (x *Project[ProjectEcosystem, ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem]) TakeFirstModule() *Module[ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem] {
	for _, module := range x.Modules {
		return module
	}
	return nil
}
