package main

import (
	"SixProject/internal/utils"
	"SixProject/view"
)

func main() {

	//启动UI
	view.InitView()
	//关闭日志文件
	defer utils.LogFile.Close()
}
