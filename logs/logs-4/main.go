package main

import "log"

func main() {
	// log.SetProfix()向程序的日志消息添加前缀
	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")
	log.Fatal("Hey, I'm an error log!")
}
