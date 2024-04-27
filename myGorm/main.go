package main

import (
	"fmt"
	"myGorm/config"
	"myGorm/model"
	"reflect"
)

func init() {
	config.ReadConfig()
}

func main() {
	//
	//db, err := gorm.Open(mysql.New(mysql.Config{
	//	DSN:                       "root:Zxcv123.00@tcp(127.0.0.1:3307)/gorm_test?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	//	DefaultStringSize:         256,                                                                                   // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                                                                                  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                                                                                  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                                                                                  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                                                                                 // 根据当前 MySQL 版本自动配置
	//}), &gorm.Config{})
	//
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//
	//userRepo := mapper.UserRepository{}
	// 连接数据库
	//db := model.GetDB()

	// 创建User类型的GenericRepository实例
	userRepository := model.RepositoryFactory[model.UserInfo](reflect.TypeOf(model.UserInfo{}))

	// 示例操作：查询、创建、更新和删除用户
	// 查询用户
	user, err := userRepository.GetByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User fetched: %+v\n", user)

	// 创建用户
	newUser := model.UserInfo{
		// ... 设置新用户的字段 ...
	}
	err = userRepository.Create(newUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %+v\n", newUser)

	// 更新用户
	updatedUser := model.UserInfo{
		// ... 设置要更新的用户字段 ...
	}
	err = userRepository.UpdateByID(1, updatedUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User updated: %+v\n", updatedUser)

	// 删除用户
	err = userRepository.DeleteByID(1)
	if err != nil {
		panic(err)
	}
	fmt.Println("User deleted")

}
