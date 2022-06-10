package main

import (
	"github.com/kivihub/go-fragment/logger"
	"github.com/kivihub/go-fragment/util"
	utils "github.com/kivihub/go-project/util"
	"os"
)

var setEnv = os.Setenv("FromGlobalVar", "any")

func init() {
	logger.Infoln("[main.go#init] start")
	os.Setenv("FromInit", "any")
}

func main() {
	logger.Infoln("[main.go#main] start")
	os.Setenv("FromMain", "any")
	util.PrintAllEnv()

	utils.Echo("call go-project")
}
