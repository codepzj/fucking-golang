package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

func main() {
	r := gin.Default()
	c := cache.New(30*time.Second, 1*time.Minute)
	r.GET("/", func(ctx *gin.Context) {
		c.Set("name", "codepzj", cache.DefaultExpiration)
		ctx.String(200, "set name successfully")
	})
	r.GET("/cache", func(ctx *gin.Context) {
		fmt.Println("cache items", c.Items())
		val, exp, found := c.GetWithExpiration("name")
		if found {
			ctx.String(200, "get cache successfully, value %s, expired at %s", val, exp.Format("2006-01-02 15:04:05"))
			return
		}
		ctx.String(200, "get cache failed, expired")
	})
	r.GET("/delete", func(ctx *gin.Context) {
		c.Delete("name")
		ctx.String(200, "delete cache successfully")
	})
	r.Run(":8080")
}
