package main

import (
	"SixProject/internal/utils"
	"SixProject/view"
)

func main() {
	//初始化log
	utils.LogInit()
	//启动UI
	view.InitView()

	defer utils.LogFile.Close()
}
