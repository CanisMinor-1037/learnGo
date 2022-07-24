package main

import (
	"fmt"
	"log"
)

func main() {
	//log.Panic()函数终止程序，获得消息日志和错误堆栈跟踪
	log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")
}
