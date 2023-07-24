package component_dao

import (
	"context"
	"github.com/scagogogo/sca-base-module-components/pkg/models"
)

// ComponentDao 组件的dao
type ComponentDao[ComponentEcosystem any] interface {

	// Create 新增组件
	Create(ctx context.Context, c *models.Component[ComponentEcosystem]) error

	// Delete 删除组件
	Delete(ctx context.Context, name, version string) error

	// Upsert 存在则更新，不存在则插入
	Upsert(ctx context.Context, c *models.Component[ComponentEcosystem]) error

	// Find 查找组件
	Find(ctx context.Context, name, version string) (*models.Component[ComponentEcosystem], error)

	// FindByNamePrefix 根据name前缀查询
	FindByNamePrefix(ctx context.Context, namePrefix string) ([]*models.Component[ComponentEcosystem], error)

	// LoadAll 加载当期的所有组件
	LoadAll(ctx context.Context) ([]*models.Component[ComponentEcosystem], error)
}
