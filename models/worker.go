package main

import (
    "gorm.io/gorm"
    "fmt"
    "gorm.io/driver/mysql"
)

type Models_worker struct {
    WorkerID          uint `gorm:"primaryKey;autoIncrement"`
    Workername        string `gorm:"not null"`
    WorkerGender      string
    Workerage         int
    Workerexperience  string
    WorkerphoneNumber string `gorm:"not null"`
    WorkerAddress     string
    Workerdescription string
    Rate        float64
    UserID      uint
}

func main() {
    dsn := "root:000000@tcp(localhost:3306)/dev?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err!= nil {
        fmt.Println("数据库连接失败", err.Error())
    }

    if db.Migrator().HasTable(&Models_worker{}) == false {
        db.AutoMigrate(&Models_worker{})
        fmt.Println("创建表成功")
    } else {
        fmt.Println("表已存在,可以插入数据")
    }
    u := Models_worker{
        Workername: "小王",
        WorkerGender: "男",
        Workerage: 30,
        Workerexperience: "5年",
        WorkerphoneNumber: "13888888888",
        WorkerAddress: "广州",
        Workerdescription: "5年搬运工作经验，服务态度很好",
        Rate: 4500,
        UserID: 6,
    }
    if err := db.Create(&u).Error; err!= nil {
        fmt.Println("数据插入失败", err)
    } else {
        fmt.Println("数据插入成功")
    }
}