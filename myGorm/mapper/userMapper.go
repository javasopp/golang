package mapper

import (
	"gorm.io/gorm"
	"myGorm/model"
)

// UserRepository 用户Repository实现
type UserRepository struct{}

func (u UserRepository) GetByID(db *gorm.DB, id int64) (*model.User, error) {
	var user model.User
	return &user, db.First(&user, id).Error
}

func (u UserRepository) GetByCondition(db *gorm.DB, cond string, args ...interface{}) ([]model.User, error) {
	var users []model.User
	return users, db.Where(cond, args...).Find(&users).Error
}

func (u UserRepository) GetByConditionWithPagination(db *gorm.DB, cond string, pageSize int, page int, args ...interface{}) ([]model.User, error) {
	var users []model.User
	offset := (page - 1) * pageSize
	return users, db.Where(cond, args...).Limit(pageSize).Offset(offset).Find(&users).Error
}

func (u UserRepository) Create(db *gorm.DB, user model.User) error {
	return db.Create(&user).Error
}

func (u UserRepository) BatchCreate(db *gorm.DB, users []model.User) error {
	return db.Create(users).Error
}

func (u UserRepository) UpdateByID(db *gorm.DB, id int64, updateUser model.User) error {
	return db.Model(&model.User{}).Where("id = ?", id).UpdateColumns(updateUser).Error
}

func (u UserRepository) UpdateByCondition(db *gorm.DB, cond string, args []interface{}, updateUser model.User) error {
	return db.Model(&model.User{}).Where(cond, args...).UpdateColumns(updateUser).Error
}

func (u UserRepository) DeleteByID(db *gorm.DB, id int64) error {
	return db.Delete(&model.User{}, id).Error
}

func (u UserRepository) DeleteByCondition(db *gorm.DB, cond string, args ...interface{}) error {
	return db.Where(cond, args...).Delete(&model.User{}).Error
}
