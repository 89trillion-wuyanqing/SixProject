package utils

import (
	"log"
	"os"
)

var LogFile *os.File

func init() {
	// 按照所需读写权限创建文件
	var err error
	LogFile, err = os.OpenFile("config/clientLog.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//设置日志输出到 f
	log.SetOutput(LogFile)

}
