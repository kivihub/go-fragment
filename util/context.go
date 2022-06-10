package utils

import (
	"flag"
	"fmt"
	"os"
)

// Version from go build -ldflags "-X main.Version=1.0.0" -x -o product main.go
var Version string = ""
var PsmFromFlag string
var PsmFromEnv string = os.Getenv("psm")

// 1. flag.Parse方法调用前，所声明的解析参数必须不少于实际传入的参数集合，否则会报错，所以不适合用于设置全局变量多次获取
// 2. ldflag也不适合全局变量定义，容易造成循环依赖
// 3. 全局变量传递推荐用环境变量，通过os.Getenv获取
func Init() {
	// 相同的Flag name只能解析一次
	flag.StringVar(&PsmFromFlag, "psm", "default", "PSM")
	//flag.Parse()

	flag.StringVar(&Version, "version", "default", "Version")
	flag.Parse()

	fmt.Printf("[context.go] Version is %s\n", Version)
	fmt.Printf("[context.go] psmFromFlag is %s\n", PsmFromFlag)
	fmt.Printf("[context.go] psmFromEnv is %s\n", PsmFromEnv)
}
