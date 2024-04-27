package model

type BigIntID struct {
	ID int64 `gorm:"primary_key"`
}

type UserInfo struct {
	BigIntID
	//gorm.Model
	Name     string
	Email    string
	Password string
}

func (u UserInfo) TableName() string {
	return "user_info"
}
