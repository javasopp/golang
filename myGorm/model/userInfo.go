package model

type BigIntID struct {
	ID int64 `gorm:"primary_key"`
}

type User struct {
	BigIntID
	//gorm.Model
	Name     string
	Email    string
	Password string
}

func (User) TableName() string {
	return "user_info"
}
