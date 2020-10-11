package main

import (
	"fmt"
	"time"
)

// 1、 goroutine间通信可通过时间定时器管理 旨在避免程序在goroutine返回之前退出。
//     但对于复杂并发场景不适宜，故引入channel

func main() {
	go slowFunc()
	fmt.Println("i am waiting for ")
	time.Sleep(8 * time.Second)
}

func slowFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("i am coming")
}

