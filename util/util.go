package util

import (
	"context"
	"flag"
	"github.com/kivihub/go-fragment/logger"

	"github.com/shirou/gopsutil/process"

	"os"
)

// Version from go build -ldflags "-X main.Version=1.0.0" -x -o product main.go
var Version = ""
var PsmFromFlag string
var PsmFromEnv = os.Getenv("psm")

//func init() {
//	logger.Infoln("[util.go#init] start")
//	PrintOsEnv()
//}

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

	logger.Infof("[util.go#ParseFlag] Version is %s\n", Version)
	logger.Infof("[util.go#ParseFlag] psmFromFlag is %s\n", PsmFromFlag)
	logger.Infof("[util.go#ParseFlag] psmFromEnv is %s\n", PsmFromEnv)
	logger.Infof("[util.go#ParseFlag] flag.args", flag.Args())
}

func ParseFlagAgain() {
	if !flag.Parsed() {
		ParseFlag()
	}
	ret := flag.Lookup("psm")
	logger.Infof("[util.go#ParseFlagAgain] psmFromFlag is %s\n", ret.Value)
}

func PrintOsArgs() {
	logger.Infoln("os.Args", os.Args)
	for idx, arg := range os.Args {
		logger.Infoln(idx, arg)
	}
}

func PrintOsEnv() {
	logger.Infoln("os.Env", os.Environ())
}

func PrintCtx() {
	ctx := context.Background()
	logger.Infof("[util.go#PrintCtx] %v\n", ctx)
}

func CycleImport() {
	//util.GetVersion()
}

func SetOSEnv(key, value string) (ret string) {
	os.Setenv(key, value)
	return value
}

func PrintProcessInfo() {
	process.Pids()
}
