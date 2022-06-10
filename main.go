package main

import (
	"go-fragment/logger"
	"go-fragment/util"
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
}
