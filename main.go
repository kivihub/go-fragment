package main

import (
	"github.com/kivihub/go-fragment/logger"
	"github.com/kivihub/go-fragment/util"
)

var setEnv = util.SetOSEnv("FromGlobalVar", "any")

func init() {
	logger.Infoln("[main.go#init] start")
	util.SetOSEnv("FromMainInit", "any")
}

func main() {
	logger.Infoln("[main.go#main] start")
	util.SetOSEnv("FromMain", "any")
}
