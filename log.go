package main

import (
	"log"
	"os"
	"time"
)

//var loger *log.Logger

func initLogs() {
	file := "./log/nacos_agent." + time.Now().Format("2006-01") + ".log"
	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("打开日志文件失败:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
