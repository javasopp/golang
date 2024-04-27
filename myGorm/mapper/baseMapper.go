package mapper

import "gorm.io/gorm"

// GenericRepository 通用 Repository 接口
type GenericRepository[T any] interface {
	// GetByID 根据ID获取实体
	GetByID(db *gorm.DB, id int64) (*T, error)

	// GetByCondition 根据条件获取实体列表
	GetByCondition(db *gorm.DB, cond string, args ...interface{}) ([]T, error)

	// GetByConditionWithPagination 根据条件获取分页实体列表
	GetByConditionWithPagination(db *gorm.DB, cond string, pageSize int, page int, args ...interface{}) ([]T, error)

	// Create 创建实体
	Create(db *gorm.DB, entity T) error

	// BatchCreate 批量创建实体
	BatchCreate(db *gorm.DB, entities []T) error

	// UpdateByID 根据ID更新实体
	UpdateByID(db *gorm.DB, id int64, update T) error

	// UpdateByCondition 根据条件更新实体
	UpdateByCondition(db *gorm.DB, cond string, args []interface{}, update T) error

	// DeleteByID 根据ID删除实体
	DeleteByID(db *gorm.DB, id int64) error

	// DeleteByCondition 根据条件删除实体
	DeleteByCondition(db *gorm.DB, cond string, args ...interface{}) error
}
