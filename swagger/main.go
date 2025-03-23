package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	_ "swagger/docs"
)

type Post struct {
	ID          int64
	Title       string
	Content     string
	Description string
}

func main() {
	r := gin.Default()
	r.GET("/post/:id", GetPostById)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}

// GetPostById
//
//	@Summary	获取文章信息
//	@Produce	json
//	@Param		id	path		string	true	"文章ID"
//	@Success	200	{object}	Post	"成功返回文章信息"
//	@Failure	400	{string}	string	"请求参数错误"
//	@Router		/post/{id} [get]
func GetPostById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.JSON(200, Post{
		ID:          id,
		Title:       "codepzj",
		Content:     "测试",
		Description: "测试",
	})
}
