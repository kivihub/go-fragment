package main

import (
	go_project "github.com/kivihub/go-project"
	utils "github.com/kivihub/go-project/util"
	"go-fragment/util"
)


func main() {
	util.Init()
	go_project.SayHello()

	// import时指定的是目录，跟package没关系
	utils.Echo("aa")
}
