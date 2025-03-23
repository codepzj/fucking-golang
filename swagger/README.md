# 使用 Swagger 为 Go 项目生成 API 文档

Swagger 是一个基于 OpenAPI 规范设计的工具，用于为 RESTful API 生成交互式文档。本文将介绍如何在 Go 项目中集成 Swagger，特别是结合 Gin 框架生成 API 文档。

## 安装 Swagger

### 全局安装 `swag` CLI

`swag` 是 Swagger 的命令行工具，用于生成 API 文档。可以通过以下命令全局安装：

```bash
go get github.com/swaggo/swag/cmd/swag@latest
go install github.com/swaggo/swag/cmd/swag@latest
```

### 项目依赖安装

在项目中需要安装以下依赖以支持 Gin 和 Swagger 的集成：

```bash
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files
go get github.com/alecthomas/template
```

### 格式化 Swagger 注释

使用 `swag fmt` 命令可以格式化项目中的 Swagger 注释，确保注释符合规范：

```bash
swag fmt
```

## 使用 `swag` CLI 生成文档

运行以下命令生成 Swagger 文档（默认生成 `docs.go`、`swagger.json` 和 `swagger.yaml` 文件）：

```bash
swag init
```

### `swag init` 常用选项

| 选项                        | 说明                                                                                 | 默认值         |
|-----------------------------|--------------------------------------------------------------------------------------|----------------|
| `--generalInfo, -g`         | 指定包含通用 API 信息的 Go 文件路径                                                  | `main.go`      |
| `--dir, -d`                 | 指定解析的目录                                                                      | `./`           |
| `--exclude`                 | 排除解析的目录（多个目录用逗号分隔）                                                | 空             |
| `--propertyStrategy, -p`    | 结构体字段命名规则（`snakecase`、`camelcase`、`pascalcase`）                        | `camelcase`    |
| `--output, -o`              | 输出文件目录（`swagger.json`、`swagger.yaml` 和 `docs.go`）                         | `./docs`       |
| `--parseVendor`             | 是否解析 `vendor` 目录中的 Go 文件                                                  | 否             |
| `--parseDependency`         | 是否解析依赖目录中的 Go 文件                                                        | 否             |
| `--parseInternal`           | 是否解析 `internal` 包中的 Go 文件                                                  | 否             |
| `--instanceName`            | 设置文档实例名称                                                                    | `swagger`      |

示例：
```bash
swag init --dir ./ --output ./docs --propertyStrategy snakecase
```

## Swagger 注释格式

Swagger 使用声明式注释来定义 API 的元信息。以下是常用注释及其说明：

### 通用 API 信息

通常在 `main.go` 中定义，用于描述整个 API 的基本信息：

```go
// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
```

### API 路由注释

在具体路由处理函数上方添加注释，定义该接口的行为：

```go
// GetPostById
// @Summary 获取文章信息
// @Produce json
// @Param id path string true "文章ID"
// @Success 200 {object} Post "成功返回文章信息"
// @Failure 400 {string} string "请求参数错误"
// @Router /post/{id} [get]
func GetPostById(c *gin.Context) {
    // 函数实现
}
```

- `@Summary`：接口简述
- `@Produce`：返回的 MIME 类型
- `@Param`：参数定义（格式：`名称 位置 类型 是否必填 描述`）
- `@Success`：成功响应（格式：`状态码 {类型} 数据结构 描述`）
- `@Failure`：失败响应
- `@Router`：路由路径和方法

## 示例项目代码

以下是一个完整的示例，展示如何在 Gin 项目中集成 Swagger：

```go
package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"
	_ "swagger/docs" // 导入生成的 Swagger 文档
)

// Post 文章结构体
type Post struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Description string `json:"description"`
}

func main() {
	r := gin.Default()
	r.GET("/post/:id", GetPostById)
	// 配置 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

// GetPostById 获取文章信息
// @Summary 获取文章信息
// @Produce json
// @Param id path string true "文章ID"
// @Success 200 {object} Post "成功返回文章信息"
// @Failure 400 {string} string "请求参数错误"
// @Router /post/{id} [get]
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
```

## 生成并访问文档

1. 运行 `swag init` 生成文档。

2. 启动项目：`go run main.go`。

3. 在浏览器中访问 `http://localhost:8080/swagger/index.html`，即可查看交互式 API 文档。

   ![image-20250323141131759](https://image.codepzj.cn/image/202503231411629.png)

## 总结

通过 `swag` 和 `gin-swagger`，我们可以轻松为 Go 项目生成规范的 API 文档。只需要编写简单的注释，Swagger 就能自动生成交互式的文档页面，方便开发和调试。
