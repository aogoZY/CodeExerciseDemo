package main

import "fmt"

//
//import (
//	"fmt"
//	"time"
//)

//channel相关
//----------------------------------------------------------------------
//1、 goroutine间通信可通过时间定时器管理 旨在避免程序在goroutine返回之前退出。
//    但对于复杂并发场景不适宜，故引入channel
//
//func main() {
//	go slowFunc()
//	fmt.Println("i am waiting for ")
//	time.Sleep(8 * time.Second)
//}
//
//func slowFunc() {
//	time.Sleep(1 * time.Second)
//	fmt.Println("i am coming")
//}
//
//----------------------------------------------------------------------
//
//2 使用channel通道做通信
//func main() {
//	//channel的创建
//	c := make(chan string)
//	go slowFunc2(c)
//	msg := <-c
//	//fmt.Println(msg)
//}
//
//func slowFunc2(ch chan string) {
//	time.Sleep(1 * time.Second)
//	ch <- "slowFunc2 finished"
//}
//
//tips:无缓存channel读写值
func main() {
	//ch := make(chan string)
	//ch <- "hello"
	//fatal error: all goroutines are asleep - deadlock!  无缓存直接写值会出错

	//ch2 := make(chan string)
	//msg, ok := <-ch2
	//fmt.Println(msg, ok)   //向一个无缓存通道取值 报错

	//ch3 := make(chan string, 2)
	//msg2, ok := <-ch3
	//fmt.Println(msg2, ok) //向一个有缓存通道取值 报错

	ch4 := make(chan string, 2)
	ch4 <- "hello"
	msg3, ok := <-ch4
	fmt.Println(msg3, ok) //向一个有缓存有值的channel取值 ok
	msg4, ok := <-ch4
	fmt.Println(msg4, ok) //向一个有缓存有值的channel取值 ok


}

//----------------------------------------------------------------------
//
//3 使用缓冲通道
//func main() {
//	//创建缓冲大小长度为2、类型为string的通道
//	messages := make(chan string, 2)
//	messages <- "hello"
//	messages <- "world"
//	//关闭通道 不能再向通道发消息
//	close(messages)
//	fmt.Println("push 2 msgs to channel")
//	fmtMessage(messages)
//
//	//push 2 msgs to channel
//	//hello
//	//world
//}
//
//func fmtMessage(c chan string) {
//	for msg := range c {
//		fmt.Println(msg)
//	}
//}
//
//----------------------------------------------------------------------
//
//4 使用select语句
//执行最先返回消息的接受者，其余将丢弃
//func main() {
//	//创建两个channel
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//
//	//启动两个goroutine 分别执行两个函数
//	go pingCh1(ch1)
//	go pingCh2(ch2)
//
//	// select语句创建两个接受者，分别接受两个channel的消息。
//	// 3秒后函数pingCh1返回，向channel1发送消息
//	// 收到来自channel1的消息后，执行第一条case语句
//	select {
//	case msg1 := <-ch1:
//		fmt.Println(msg1)
//	case msg2 := <-ch2:
//		fmt.Println(msg2)
//	}
//}
////返回：ping 1
//
//func pingCh1(c chan string) {
//	time.Sleep(time.Second * 3)
//	c <- "ping 1"
//}
//
//func pingCh2(c chan string) {
//	time.Sleep(time.Second * 9)
//	c <- "ping 2"
//}
//
//----------------------------------------------------------------------
//
//5、若channel没有接受返回值，可设置指定时间，使其不再阻塞，以便接着往下执行
//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan string)
////
//	//启动两个goroutine 分别执行两个函数
//	go pingCh1(ch1)
//	go pingCh2(ch2)
//
//	// 设置超时时间，使goroutine不再等待
//	select {
//	case msg1 := <-ch1:
//		fmt.Println(msg1)
//	case msg2 := <-ch2:
//		fmt.Println(msg2)
//	case <-time.After(time.Millisecond * 500):
//		fmt.Println("do other thing,not wait anymore")
//	}
//	//do other thing,not wait anymore
//}
//
//
//func pingCh1(c chan string) {
//	time.Sleep(time.Second * 3)
//	c <- "ping 1"
//}
//
//func pingCh2(c chan string) {
//	time.Sleep(time.Second * 9)
//	c <- "ping 2"
//}
//
//----------------------------------------------------------------------
//
////6 退出通道
////需要无限制阻塞 有需要随时返回
////可定义一个quit channel，用于退出通道
//func main() {
//	//新建两个channel 一个是string类型的用于接受数据 一个是bool类型的用于退出
//	channel := make(chan string)
//	quit := make(chan bool)
//
//	go sender(channel)
//	go func() {
//		time.Sleep(time.Second * 3)
//		fmt.Println("time is up!")
//		quit <- true
//	}()
//	//for循环使用select语句，可在接受消息后一直打印
//	//由于这是一个阻塞操作，会一直打印消息，直到从quit中接收信息，结束for循环
//	for {
//		select {
//		case msg1 := <-channel:
//			fmt.Println(msg1)
//		case <-quit:
//			return
//		}
//	}
//}
//
////返回：
////sending message ing...
////sending message ing...
////sending message ing...
////sending message ing...
////time is up!
//
//func sender(c chan string) {
//	t := time.NewTicker(1 * time.Second)
//	for {
//		c <- "sending message ing..."
//		<-t.C
//	}
//}
