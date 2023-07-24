package models

// Module 表示一个模块信息
type Module[ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem any] struct {

	// 模块可以认为是一堆组件的集合，它的本质其实就是一个组件，在这个组件的基础上添加了一些属性就成了模块
	*Component[ComponentEcosystem]

	// 父模块
	ParentModule *Module[ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem]

	// 子模块
	SubModule map[string]*Module[ModuleEcosystem, ComponentEcosystem, ComponentDependencyEcosystem]

	// 当前模块的依赖
	Dependencies []*ComponentDependency[ComponentDependencyEcosystem]

	// 不同的包管理器可以有自己独有的模块级别的特性，这里允许自行扩展
	ModuleEcosystem ModuleEcosystem
}
