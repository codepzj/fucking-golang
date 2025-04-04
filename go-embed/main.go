package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"os"
	"path"
)

//go:embed demo.tmpl
var message embed.FS

var (
	outputDir = flag.String("output", "", "the output directory;eg: internal/user")
)

func main() {
	// 解析命令行参数
	flag.Parse()
	tpl, err := template.ParseFS(message, "demo.tmpl")
	if err != nil {
		return
	}
	_ = os.MkdirAll(*outputDir, os.ModePerm)
	fmt.Println(*outputDir)
	file, _ := os.Create(path.Join(*outputDir, "demo.go"))
	err = tpl.Execute(file, map[string]string{
		"packageName": "mkg",
		"info":        "hello world!",
	})

	if err == nil {
		fmt.Println("解析成功")
	} else {
		fmt.Println("解析失败")
	}
}
