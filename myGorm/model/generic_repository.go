package model

import (
	"reflect"
	"strings"
)

type GenericRepository[T EntityType] interface {
	GetByID(id int64) (*T, error)
	GetByCondition(condition string, args ...interface{}) ([]*T, error)
	GetByConditionWithPagination(cond string, pageSize int, page int, args ...interface{}) ([]T, error)

	Create(entity T) error
	BatchCreate(entities []T) error

	UpdateByID(id int64, update T) error
	UpdateByCondition(cond string, args []interface{}, update T) error

	DeleteByID(id int64) error
	DeleteByCondition(cond string, args ...interface{}) error
}

type BaseRepository[T EntityType] struct {
	entityType reflect.Type
}

func (r BaseRepository[T]) GetByID(id int64) (*T, error) {
	var entity T
	return &entity, db.Table(r.entityType.Name()).First(&entity, id).Error
}

func (r BaseRepository[T]) GetByCondition(condition string, args ...interface{}) ([]*T, error) {
	var entities []*T
	return entities, db.Table(r.entityType.Name()).Where(condition, args...).Find(&entities).Error
}

func (r BaseRepository[T]) GetByConditionWithPagination(cond string, pageSize int, page int, args ...interface{}) ([]T, error) {
	var entities []T
	offset := (page - 1) * pageSize
	return entities, db.Table(r.entityType.Name()).Where(cond, args...).Offset(offset).Limit(pageSize).Find(&entities).Error
}

func (r BaseRepository[T]) Create(entity T) error {
	return db.Table(r.entityType.Name()).Create(&entity).Error
}

func (r BaseRepository[T]) BatchCreate(entities []T) error {
	return db.Table(r.entityType.Name()).Create(&entities).Error
}

func (r BaseRepository[T]) UpdateByID(id int64, update T) error {
	return db.Table(r.entityType.Name()).Where("id = ?", id).Updates(update).Error
}

func (r BaseRepository[T]) UpdateByCondition(cond string, args []interface{}, update T) error {
	return db.Table(r.entityType.Name()).Where(cond, args...).Updates(update).Error
}

func (r BaseRepository[T]) DeleteByID(id int64) error {
	var zero T
	return db.Table(r.entityType.Name()).Delete(&zero, "id = ?", id).Error
}

func (r BaseRepository[T]) DeleteByCondition(cond string, args ...interface{}) error {
	var zero T
	return db.Table(r.entityType.Name()).Where(cond, args...).Delete(&zero).Error
}

func RepositoryFactory[T any](entityType reflect.Type) GenericRepository[T] {
	repoType := reflect.TypeOf((*GenericRepository[T])(nil)).Elem()
	repoValue := reflect.New(repoType)

	// 根据实体类型生成Repository名称（如"UserRepository"）
	repoName := strings.ToTitle(entityType.Name()) + "Repository"
	repoFullName := "model." + repoName

	// 使用反射创建Repository实例
	repository := reflect.New(reflect.TypeOf(repoFullName)).Interface().(GenericRepository[T])

	// 初始化Repository内部的BaseRepository
	baseRepoType := reflect.TypeOf((*BaseRepository[T])(nil)).Elem()
	baseRepoValue := reflect.New(baseRepoType)
	baseRepo := baseRepoValue.Interface().(BaseRepository[T])
	baseRepo.entityType = entityType

	// 将BaseRepository设置到Repository内部
	repoValue.Elem().FieldByName("base").Set(baseRepoValue)

	// 返回初始化后的Repository实例
	return repository
}
