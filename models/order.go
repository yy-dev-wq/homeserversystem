package main
import (
    "gorm.io/gorm"
    "fmt"
    "gorm.io/driver/mysql"
)
type Models_order struct {
    CustomerID       uint
    WorkerID         uint
    Order_service    string
    Order_price      float64
    Order_status     string
    Order_time       string
}

func main() {
    dsn := "root:000000@tcp(localhost:3306)/dev?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err!= nil {
        fmt.Println("数据库连接失败", err.Error())
    }
    if db.Migrator().HasTable(&Models_order{}) == false {
        db.AutoMigrate(&Models_order{})
        fmt.Println("创建表成功")
    } else {
        fmt.Println("表已存在,可以插入数据")
    }
    u := Models_order{CustomerID: 1, WorkerID: 1, Order_service: "网站建设", Order_price: 5000, Order_status: "pending", Order_time: "2023-05-05"}
    if err := db.Create(&u).Error; err!= nil {
        fmt.Println("数据插入失败", err)
    } else {
        fmt.Println("数据插入成功")
    }
}