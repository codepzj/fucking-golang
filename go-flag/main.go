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
	// 解析命令行参数
	flag.Parse()
	fmt.Println(*name, *age, *isMale, *score)
	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
