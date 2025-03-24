package main

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var sessionStore sync.Map

type User struct {
	Name     string `json:"name,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// 生成 Session ID 并存储用户数据
func generateSessionID(user User) (string, error) {
	sessionID := uuid.NewString()
	userData, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	sessionStore.Store(sessionID, userData)
	return sessionID, nil
}

// 从 sessionStore 获取用户信息
func getUserFromSession(sessionID string) (*User, bool) {
	if data, exists := sessionStore.Load(sessionID); exists {
		var user User
		if err := json.Unmarshal(data.([]byte), &user); err == nil {
			return &user, true
		}
	}
	return nil, false
}

func loginHandler(ctx *gin.Context) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 从 Cookie 获取 sessionID，并尝试获取用户信息
	if sessionID, err := ctx.Cookie("auth"); err == nil {
		if cachedUser, found := getUserFromSession(sessionID); found {
			fmt.Println("命中缓存")
			ctx.JSON(200, cachedUser)
			return
		}
	}

	fmt.Println("未命中缓存")
	sessionID, err := generateSessionID(user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "服务器错误"})
		return
	}

	ctx.SetCookie("auth", sessionID, 3600, "/", "localhost", false, false)
	ctx.JSON(200, gin.H{"message": "登录成功"})
}

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.Run()
}
