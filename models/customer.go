package main

import (
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/mysql"
)
type Models_customer struct {
    Name    string
    Gender  string
    Age     int
    Phone   string `gorm:"unique;not null"`
    Address string
    UserID  uint
}
func main() {
    dsn:="root:000000@tcp(localhost:3306)/dev?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("数据库连接失败",err.Error())
    }
	if db.Migrator().HasTable(&Models_customer{}) == false {
	   db.AutoMigrate(&Models_customer{})
	   fmt.Println("创建表成功")
	}
	if db.Migrator().HasTable(&Models_customer{}) == true{
		fmt.Println("表已存在,可以插入数据")
		u := Models_customer{
			Name:"张三",
			Gender:"男",
			Age:18,
			Phone:"12345678901",
			Address:"广州市天河区",
			UserID: 1,
		}
		if err := db.Create(u).Error; err != nil {
            fmt.Println("数据插入失败", err)
        }else{
            fmt.Println("数据插入成功")
        }
	}
}

