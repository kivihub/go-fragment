package util

import (
	"context"
	"flag"
	"github.com/kivihub/go-fragment/logger"
	"os"
)

// Version from go build -ldflags "-X main.Version=1.0.0" -x -o product main.go
var Version = ""
var PsmFromFlag string
var PsmFromEnv = os.Getenv("psm")

func init() {
	logger.Infoln("[context.go#init] start")
	PrintOsEnv()
}

func PrintAllEnv() {
	ParseFlag()
	ParseFlagAgain()
	PrintOsEnv()
	PrintCtx()
}

// ParseFlag
// 1. flag.Parse方法调用前，所声明的解析参数必须不少于实际传入的参数集合，否则会报错，所以不适合用于设置全局变量多次获取
// 2. ldflag也不适合全局变量定义，容易造成循环依赖
// 3. 全局变量传递推荐用环境变量，通过os.Getenv获取
func ParseFlag() {
	// 相同的Flag name只能解析一次
	flag.StringVar(&PsmFromFlag, "psm", "default", "PSM")
	//flag.Parse()

	flag.StringVar(&Version, "version", "default", "Version")
	if !flag.Parsed() {
		flag.Parse()
	}

	logger.Infof("[context.go#ParseFlag] Version is %s\n", Version)
	logger.Infof("[context.go#ParseFlag] psmFromFlag is %s\n", PsmFromFlag)
	logger.Infof("[context.go#ParseFlag] psmFromEnv is %s\n", PsmFromEnv)
}

func ParseFlagAgain() {
	if !flag.Parsed() {
		ParseFlag()
	}
	ret := flag.Lookup("psm")
	logger.Infof("[context.go#ParseFlagAgain] psmFromFlag is %s\n", ret.Value)
}

func PrintOsEnv() {
	logger.Infoln(os.Environ())
}

func PrintCtx() {
	ctx := context.Background()
	logger.Infof("[context.go#PrintCtx] %v\n", ctx)
}
