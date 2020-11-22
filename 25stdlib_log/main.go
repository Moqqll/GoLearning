package main

import (
	"fmt"
	"log"
	"os"
)

func lg() {
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志\n", v)
	log.Printf("这是一条%s狗\n", v)
	// log.Fatalln("这是一条会触发Fatal的日志")
	log.Panicln("这是一条会触发Panic的日志")
}

func filelog() {
	logFile, err := os.OpenFile("xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[moqqll]")
	log.Println("这是一条很普通的日志")
}

func mylogger() {
	logger := log.New(os.Stdout, "<NEW>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志")
}

func main() {
	// lg()
	// nint := log.Flags()
	// fmt.Println(nint)
	// log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条很普通的日志。")
	// // fmt.Println(log.Prefix())
	// log.SetPrefix("[moqqll]")
	// log.Println("这是一条很普通的日志。")
	// filelog()
	mylogger()
}
