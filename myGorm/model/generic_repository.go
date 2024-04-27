package model

import (
	"reflect"
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
	T          EntityType
}

func (r BaseRepository[T]) GetByID(id int64) (*T, error) {
	var entity T
	return &entity, db.First(&entity, id).Error
}

func (r BaseRepository[T]) GetByCondition(condition string, args ...interface{}) ([]*T, error) {
	var entities []*T
	return entities, db.Where(condition, args...).Find(&entities).Error
}

func (r BaseRepository[T]) GetByConditionWithPagination(cond string, pageSize int, page int, args ...interface{}) ([]T, error) {
	var entities []T
	offset := (page - 1) * pageSize
	return entities, db.Where(cond, args...).Offset(offset).Limit(pageSize).Find(&entities).Error
}

func (r BaseRepository[T]) Create(entity T) error {
	return db.Create(&entity).Error
}

func (r BaseRepository[T]) BatchCreate(entities []T) error {
	return db.CreateInBatches(&entities, 100).Error
}

func (r BaseRepository[T]) UpdateByID(id int64, update T) error {
	return db.Where("id = ?", id).Updates(update).Error
}

func (r BaseRepository[T]) UpdateByCondition(cond string, args []interface{}, update T) error {
	return db.Where(cond, args...).Updates(update).Error
}

func (r BaseRepository[T]) DeleteByID(id int64) error {
	var zero T
	return db.Delete(&zero, "id = ?", id).Error
}

func (r BaseRepository[T]) DeleteByCondition(cond string, args ...interface{}) error {
	var zero T
	return db.Where(cond, args...).Delete(&zero).Error
}

func RepositoryFactory[T EntityType](entityType reflect.Type) GenericRepository[T] {
	//repoType := reflect.TypeOf((*GenericRepository[T])(nil)).Elem()
	//repoValue := reflect.New(repoType)
	//
	////db := GetDB()
	//
	//// 根据实体类型生成Repository名称（如"UserRepository"）
	//repoName := entityType.Name() + "Repository"
	//repoFullName := "model." + repoName
	//
	//logrus.Info("我是当前的名字: " + repoFullName)
	//
	//resultRepoType := reflect.ValueOf(repoFullName).Type()
	//// 使用反射创建Repository实例
	//repository := reflect.New(resultRepoType).Interface().(GenericRepository[T])
	//
	//// 初始化Repository内部的BaseRepository
	//baseRepoType := reflect.TypeOf((*BaseRepository[T])(nil)).Elem()
	//baseRepoValue := reflect.New(baseRepoType)
	//baseRepo := baseRepoValue.Interface().(BaseRepository[T])
	//baseRepo.entityType = entityType
	//
	//// 将BaseRepository设置到Repository内部
	//repoValue.Elem().FieldByName("base").Set(baseRepoValue)
	//
	//// 返回初始化后的Repository实例
	//return repository
	baseRepo := BaseRepository[T]{entityType: entityType}

	// 返回BaseRepository[T]实例，它实现了GenericRepository[T]接口
	return baseRepo
}
