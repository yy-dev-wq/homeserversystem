package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "homeserversystem/models"
    "homeserversystem/utils"
)

// Login 处理用户登录
func Login(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
        return
    }

    // 在数据库中查询用户
    var foundUser models.User
    result := models.DB.Where("username =? AND password =?", user.Username, user.Password).First(&foundUser)
    if result.Error!= nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
        return
    }

    // 生成 JWT token
    token, err := utils.GenerateJWT(foundUser.Username)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "生成 token 失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// Register 处理用户注册
func Register(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err!= nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求参数"})
    return
    }

    // 对密码进行哈希处理
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
    if err!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "密码哈希失败"})
        return
    }
    newUser.Password = string(hashedPassword)

    // 创建新用户并保存到数据库
    result := models.DB.Create(&newUser)
    if result.Error!= nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}