package main

import (
	"log"
	"os"
	"runtime"
)

// 创建一个全局的日志记录器
var logger = log.New(os.Stdout, "", log.Lshortfile)

// LogMessage 是一个自定义的日志函数，可以记录调用者的信息
func LogMessage(calldepth int, msg string) {
	// 获取调用者的信息
	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "unknown"
		line = 0
	}
	// 记录日志，包含调用者的文件名和行号
	logger.Printf("%s:%d: %s", file, line, msg)
}

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	LogMessage(1, "This is a log message with calldepth 1")
	LogMessage(2, "This is a log message with calldepth 2")
	LogMessage(3, "This is a log message with calldepth 3")
}
