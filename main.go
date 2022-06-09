package main

import (
	go_project "github.com/kivihub/go-project"
	utilsaa "github.com/kivihub/go-project/util"
	aaaaa "go-fragment/util"
)

// import时指定的是目录，而不是package
// 目录下go声明的package必须相同，但不强制和文件夹名一致
func main() {
	aaaaa.Init()
	go_project.SayHello()

	utilsaa.Echo("aa")
}
