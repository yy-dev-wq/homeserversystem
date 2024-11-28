package main

import (
	"gorm.io/gorm"
	"fmt"
	"gorm.io/driver/mysql"
)

type Models_user struct {
    Username    string 
    Password    string 
    Role        string 
}
func main() {
    dsn:="root:000000@tcp(localhost:3306)/dev?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("数据库连接失败",err.Error())
    }
    if db.Migrator().HasTable(&Models_user{}) == true{
		fmt.Println("表已存在") 
		u := Models_user{
		    Username:"admin",
		    Password:"123456",
		    Role:"客户",
	
		}
		if err := db.Create(u).Error; err != nil {
            fmt.Println("数据插入失败", err)
        }else{
            fmt.Println("数据插入成功")
        }
    }
}

    /*}
       db.AutoMigrate(&Models_user{})
	   fmt.Println("创建表成功")
        }
    }*/

