# Go 语言flag库使用指南

`flag` 库是 Go 标准库中用于解析命令行参数的工具。它可以将命令行输入的参数提取并绑定到程序中的变量，适用于需要从终端接收用户输入的场景。本文将介绍 `flag` 库的基本用法、常用函数以及示例。


## `os.Args`：获取命令行参数

`os.Args` 是一个字符串切片（`[]string`），用于获取命令行传入的所有参数。它以空格分隔参数。

### 示例代码
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	res := os.Args
	fmt.Println(res)
}
```

### 执行结果
```
D:\Code\Program\fucking-golang\go-flag> go run main.go a b c d
[C:\Users\pzj\AppData\Local\Temp\go-build4165448248\b001\exe\main.exe a b c d]
```

## `flag.Parse`：解析命令行参数

`flag.Parse()` 是 `flag` 库的核心函数，用于解析命令行参数并将其绑定到通过 `flag` 定义的变量上。在调用 `flag.Parse()` 之前，需要先使用 `flag.Type` 函数定义参数。

## `flag` 库常用函数

`flag.Args()` 返回解析后剩余的非标志参数（即不以 `-` 或 `--` 开头的参数），类型为 `[]string`。  
`flag.NArg()` 返回 `flag.Args()` 中非标志参数的个数。  
`flag.NFlag()` 返回命令行中已解析的标志参数个数。

## `flag.Type`：定义命令行参数

`flag` 库提供了多种方法来定义不同类型的命令行参数，格式如下：

```go
flag.Type(name string, value Type, usage string) *Type
```

`name`：参数名称（在命令行中以 `-name` 或 `--name` 形式使用）。  
`value`：默认值。  
`usage`：参数的用途说明（用于生成帮助信息）。  
返回值：指向该类型变量的指针。

## 示例代码

以下是一个完整的示例，展示如何使用 `flag` 库解析不同类型的命令行参数：

```go
package main

import (
	"flag"
	"fmt"
)

var (
	name   = flag.String("name", "", "get name from terminal")
	age    = flag.Int64("age", 0, "get age from terminal")
	isMale = flag.Bool("is_male", true, "get isMale from terminal")
	score  = flag.Float64("score", 99.99, "get score from terminal")
)

func main() {
	flag.Parse()
	fmt.Println(*name, *age, *isMale, *score)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
```

### 执行命令
```
D:\Code\Program\fucking-golang\go-flag> go run main.go -name=codepzj -age=21 -is_male=true -score=100.00 a b c
```

### 输出结果
```
codepzj 21 true 100
[a b c]
3
4
```

## 总结

`flag` 库提供了一种简单而强大的方式来处理命令行参数。通过 `flag.Type` 定义参数、`flag.Parse()` 解析参数，以及辅助函数（如 `flag.Args()`、`flag.NArg()`、`flag.NFlag()`），开发者可以轻松实现灵活的命令行工具。结合实际需求，`flag` 库是 Go 语言开发中不可或缺的工具之一。
