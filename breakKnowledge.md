# golang知识点碎片

### go的数组、切片,如何扩容?指针引用还是值引用?

- 定义方式不一样
- 初始化方式不一样:数组的长度是固定的(需指定大小)，而切片不是（切片是动态的数组）
- 传递方式不一样: 切片是指针类型，数组是值类型.

切片比数组多一个属性：容量（cap)

切片的底层是数组

```go
package main

import "fmt"

//slice和array的区别

func main()  {
   //定义slice切片 指针传递 改变原值
   slice:= []int{1,2,3}
   sliceCopy := slice
   sliceCopy[0]=9
   fmt.Println(slice)
   //9 2 3]    切片是指针 将原有值也改变了
   fmt.Println(sliceCopy)
   //[9 2 3]

   //定义array数组 值传递
   array:=[3]int{1,2,3}
   arrayCopy:=array    //数组是将整个复制的值传递
   arrayCopy[0]=9
   fmt.Println(array)
   //[1 2 3]
   fmt.Println(arrayCopy)
   //[9 2 3]
}
```



### go的协程原理?为啥支持高并发?为啥轻量级?

**进程、线程、协程的关系和区别：**

- 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。
- 线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是的)。
- 协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

**为什么协程比线程轻量？**

- go协程调用跟切换比线程效率
- go协程占用内存少



### channel 向channel里写、读、已关闭 会发生什么情况

不同于传统的多线程并发模型使用共享内存来实现线程间通信的方式，golang 的哲学是通过 channel 进行协程(goroutine)之间的通信来实现数据共享：

> Do not communicate by sharing memory; instead, share memory by communicating.  --《effective go》

```go
package main

import (
	"fmt"
	"time"
)
----------------------------------------------------------------------
1、 goroutine间通信可通过时间定时器管理 旨在避免程序在goroutine返回之前退出。
    但对于复杂并发场景不适宜，故引入channel

func main() {
	go slowFunc()
	fmt.Println("i am waiting for ")
	time.Sleep(8 * time.Second)
}

func slowFunc() {
	time.Sleep(1 * time.Second)
	fmt.Println("i am coming")
}

----------------------------------------------------------------------

2 使用channel通道做通信
func main() {
	//channel的创建
	c := make(chan string)
	go slowFunc2(c)
	msg := <-c
	//fmt.Println(msg)
}

func slowFunc2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "slowFunc2 finished"
}

----------------------------------------------------------------------

3 使用缓冲通道
func main() {
	//创建缓冲大小长度为2、类型为string的通道
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	//关闭通道 不能再向通道发消息
	close(messages)
	fmt.Println("push 2 msgs to channel")
	fmtMessage(messages)

	//push 2 msgs to channel
	//hello
	//world
}

func fmtMessage(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

----------------------------------------------------------------------

4 使用select语句
执行最先返回消息的接受者，其余将丢弃
func main() {
	//创建两个channel
	ch1 := make(chan string)
	ch2 := make(chan string)

	//启动两个goroutine 分别执行两个函数
	go pingCh1(ch1)
	go pingCh2(ch2)

	// select语句创建两个接受者，分别接受两个channel的消息。
	// 3秒后函数pingCh1返回，向channel1发送消息
	// 收到来自channel1的消息后，执行第一条case语句
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}
//返回：ping 1

func pingCh1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "ping 1"
}

func pingCh2(c chan string) {
	time.Sleep(time.Second * 9)
	c <- "ping 2"
}

----------------------------------------------------------------------

5、若channel没有接受返回值，可设置指定时间，使其不再阻塞，以便接着往下执行
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
//
	//启动两个goroutine 分别执行两个函数
	go pingCh1(ch1)
	go pingCh2(ch2)

	// 设置超时时间，使goroutine不再等待
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(time.Millisecond * 500):
		fmt.Println("do other thing,not wait anymore")
	}
	//do other thing,not wait anymore
}


func pingCh1(c chan string) {
	time.Sleep(time.Second * 3)
	c <- "ping 1"
}

func pingCh2(c chan string) {
	time.Sleep(time.Second * 9)
	c <- "ping 2"
}

----------------------------------------------------------------------

//6 退出通道
//需要无限制阻塞 有需要随时返回
//可定义一个quit channel，用于退出通道
func main() {
	//新建两个channel 一个是string类型的用于接受数据 一个是bool类型的用于退出
	channel := make(chan string)
	quit := make(chan bool)

	go sender(channel)
	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("time is up!")
		quit <- true
	}()
	//for循环使用select语句，可在接受消息后一直打印
	//由于这是一个阻塞操作，会一直打印消息，直到从quit中接收信息，结束for循环
	for {
		select {
		case msg1 := <-channel:
			fmt.Println(msg1)
		case <-quit:
			return
		}
	}
}

//返回：
//sending message ing...
//sending message ing...
//sending message ing...
//sending message ing...
//time is up!

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)
	for {
		c <- "sending message ing..."
		<-t.C
	}
}

```



1. docker k8s?

2. 算法:树、链表、冒泡排序

   

   

   

   ### 



1. 多态
2. 定时任务crontab linux 定时任务写法
3. go mod和go vendor