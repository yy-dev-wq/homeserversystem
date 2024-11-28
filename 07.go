/*package main

import("fmt"
       "database/sql"
	   _"github.com/go-sql-driver/mysql")
func main()  {
	    db,err := sql.Open("mysql","root:000000@tcp(localhost:3306)/dev?charset=utf8")
        if err != nil{
			    fmt.Println("连接数据库失败")

		}else {
			    fmt.Println("连接数据库成功")
		}	
		defer db.Close()
}*/
/*package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

// User结构体对应数据库中的users表
type User struct {
    ID       int    `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"type:varchar(50);not null"`
    Password string `gorm:"type:varchar(255);not null"`
    Role     string `gorm:"type:enum('customer', 'worker', 'admin');not null"`
}

func main() {
    // 连接数据库
    dsn := "root:000000@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err!= nil {
        panic(err)
    }

    // 自动迁移创建users表（如果不存在）
    err = db.AutoMigrate(&User{})
    if err!= nil {
        panic(err)
    }

    // 示例：插入一条用户数据
    newUser := User{
        Username: "lxy",
        Password: "12345678", // 这里应该是实际加密后的密码
        Role:     "customer",
    }
    result := db.Create(&newUser)
    if result.Error!= nil {
        panic(result.Error)
    }

    // 示例：查询所有用户数据
    var users []User
    db.Find(&users)
    for _, user := range users {
        println(user.Username, user.Role)
    }
}*/
package main

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"
)

// Service结构体对应服务项目表（services）
type Service struct {
    ID          int     `gorm:"primaryKey;autoIncrement"`
    Name        string  `gorm:"type:varchar(100);not null"`
    Description string  `gorm:"type:text"`
    Price       float64 `gorm:"type:decimal(10, 2)"`
}

// Order结构体对应订单表（orders）
type Order struct {
    ID         int       `gorm:"primaryKey;autoIncrement"`
    CustomerID int       `gorm:"not null"`
    ServiceID  int       `gorm:"not null"`
    WorkerID   int       `gorm:""`
    OrderTime  time.Time `gorm:"type:datetime"`
    Status     string    `gorm:"type:enum('pending', 'assigned', 'in_progress', 'completed');not null"`
}

// Review结构体对应评价表（reviews）
type Review struct {
    ID         int    `gorm:"primaryKey;autoIncrement"`
    OrderID    int    `gorm:"not null"`
    CustomerID int    `gorm:"not null"`
    WorkerID   int    `gorm:"not null"`
    Rating     int    `gorm:"not null"`
    Comment    string `gorm:"type:text"`
}

func main() {
    // 连接数据库
    dsn := "root:000000@tcp(localhost:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err!= nil {
        panic(err)
    }

    // 自动迁移创建各个表（如果不存在）
    err = db.AutoMigrate(&Service{}, &Order{}, &Review{})
    if err!= nil {
        panic(err)
    }

    // 以下是一些示例操作，可根据实际需求扩展

    // 插入一条服务项目数据示例
    newService := Service{
        Name:        "家庭清洁服务",
        Description: "全面清洁家庭各个区域",
        Price:       150.00,
    }
    result := db.Create(&newService)
    if result.Error!= nil {
        panic(result.Error)
    }

    // 插入一条订单数据示例
    newOrder := Order{
        CustomerID: 1,
        ServiceID:  1,
        Status:     "pending",
    }
    result = db.Create(&newOrder)
    if result.Error!= nil {
        panic(result.Error)
    }

    // 插入一条评价数据示例
    newReview := Review{
        OrderID:    1,
        CustomerID: 1,
        WorkerID:   1,
        Rating:     4,
        Comment:    "服务很不错，很满意。",
    }
    result = db.Create(&newReview)
    if result.Error!= nil {
        panic(result.Error)
    }

    // 查询所有服务项目数据示例
    var services []Service
    db.Find(&services)
    for _, service := range services {
        println(service.Name, service.Price)
    }

    // 查询所有订单数据示例
    var orders []Order
    db.Find(&orders)
    for _, order := range orders {
        println(order.Status)
    }

    // 查询所有评价数据示例
    var reviews []Review
    db.Find(&reviews)
    for _, review := range reviews {
        println(review.Rating, review.Comment)
    }
}