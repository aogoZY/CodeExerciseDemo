# golang知识点碎片

[TOC]



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



channel已关闭的读写:写报错,读会根据关闭前有无值做处理

```go l
package main

import (
   "fmt"
)

channel关闭之后 写出错 读取决于关闭之前是否有值 有值的话可读取出值 flag为true 否则返回默认值 flag为false
------------------------------------------------------------------------------
func main()  {
   //channel关闭之后，再写不了东西
   ch:=make(chan string,3)
   close(ch)
   ch<-"hello"
//panic: send on closed channel

------------------------------------------------------------------------------

//channel关闭之后，是否可以读取决于在关闭之前是否还有值，不存在的值为0，flag为false
func main() {
   //int型channel
   cn := make(chan int, 3)
   cn <- 1
   close(cn)
   num1, ok1 := <-cn
   num2, ok2 := <-cn
   num3, ok3 := <-cn
   fmt.Printf("读num channel: %v, %v\n", num1, ok1)
   fmt.Printf("再读num channel: %v, %v\n", num2, ok2)
   fmt.Printf("再再读num channel: %v, %v\n", num3, ok3)

   cs := make(chan string, 3)
   cs <- "a"
   cs <- "b"
   close(cs)
   str1, ok1 := <-cs
   str2, ok2 := <-cs
   str3, ok3 := <-cs
   fmt.Printf("读str channel: %v, %v\n", str1, ok1)
   fmt.Printf("再读str channel: %v, %v\n", str2, ok2)
   fmt.Printf("再再读str channel: %v, %v\n", str3, ok3)

   type Student struct {
      Name string
   }
   cst := make(chan Student, 3)
   cst <- Student{Name: "aogo"}
   close(cst)
   stu1, ok1 := <-cst
   stu2, ok2 := <-cst
   stu3, ok3 := <-cst
   fmt.Printf("读stu channel: %v, %v\n", stu1, ok1)
   fmt.Printf("再读stu channel: %v, %v\n", stu2, ok2)
   fmt.Printf("再再读stu channel: %v, %v\n", stu3, ok3)

//输出：
// 读num channel: 1, true
// 再读num channel: 0, false
// 再再读num channel: 0, false
// 读str channel: a, true
// 再读str channel: b, true
// 再再读str channel: , false
// 读stu channel: {aogo}, true
// 再读stu channel: {}, false
// 再再读stu channel: {}, false


}
```



### go mod vs go vendor

go的包引用的前世今生

(1)gopath

GOPATH 解决了第三方源码依赖的问题，用户比较容易的能够引用一个 Github 上的第三方代码库，进行打包构建。但这引入了一个问题，所有的团队依赖的只互联网上的一份源代码，如果这个代码库被人篡改，删除，你们的构建也会受到影响。

![image-20201031010812759](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010812759.png)

(2)为了解决这个问题，社区里提出了 Go Vendoring 的方法。

![image-20201031010849483](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010849483.png)

但这又引入了新的问题，随着项目的依赖增多，你的代码库可能会越来越大，例如上图，你依赖了 mypkg@v1.2.3， 这个库里有几百兆的图片，当你依赖了这个库，在 pull 代码的时候就大大增加的你的代码库大小。



(3)(Athens – Go 依赖管理工具

![image-20201031011059972](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031011059972.png)

所以你只需要标注你所有的依赖包的地址，go get 会安装先近后远的顺序帮你下载依赖，并且缓存在本地 mods 目录。

设置好之后，执行 Go build，就能够让 Athens 从 GoCenter 进行依赖包的下载了。当然 Athens 也能够从 Artifactory 进行 Go 依赖包的下载。



```bash
go mod init   生成go.sum和go.mod
go get -v
go build
1、若是需要翻墙的源码
replace (
	golang.org/x/net v0.0.0-20181106171534-e4dc69e5b2fd => github.com/golang/net latest
	//这表示系统依赖的 golang.org/x/net v0.0.0 这个版本应该从github.com/golang/net 这个地方下载latest也就是最新版本.
)
2、我依赖的不能是最新的代码怎么办
github.com/asdine/storm v2.1.2+incompatible 修改为v2.1.1

3、我依赖的某个项目是我修改过的,和官方版本不一样怎么办
直接clone一个官方的版本到自己的github上,然后修改. 待修改完毕以后,新建一个版本即可.
replace (
github.com/ethereum/go-ethereum v1.8.17 => github.com/nkbai/go-ethereum v1.9.1
)
```

go get命令拉包,eg:go get github.com/kardianos/govendor ，会将依赖包下载到`GOPATH`的路径下。

包管理工具:govendor 在go build时的应用路径搜索调整成为 `当前项目目录/vendor` 目录方式

cd /home/gopath/src/aogoWeb

#### go vendor

1. 初始化vendor目录      govendor  init(生成一个vendor目录)
2.  govendor add +external  将GOPATH中本工程使用到的依赖包自动移动到vendor目录中,若本地GOPATH没有依赖包，先go get相应的依赖包

#### go mod

(版本go1.13)

1. 设置环境

go env -w go111module=on

go env -w goproxy=yun.paic.com.cn

go env sumdb=off 不校验包的hash值

2. go mod init

   生成go.sum & go.mod

**go.sum** 是类似于比如 dep 的 Gopkg.lock 的一类文件，它详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。

**go.mod** 是启用了 Go moduels 的项目所必须的最重要的文件，它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:(包引用路径+版本号)

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

### make 和new的区别

new：为所有的类型分配内存，并初始化为零值，返回指针。

make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。





### **1.new和make的区别？**

- 相同点：

  - new和make都是用来开辟空间的

- 不同点：

  - new是初始化一个类型的指针,返回的是类型**指针**，而里面的值为默认初始值，只对值类型有效

  - make是针对**slice**切片，**map**字典，**chan管道**初始化，并且返回对应的初始值

    - 并非返回指针，而是对应的类型有效值

      

      

      make 只能为 slice、map或 channel 类型分配内存并初始化，同时返回一个有初始值的 slice、map 或 channel 类型引用，不是指针。内建函数 new 用来分配内存，它的第一个参数是一个类型，不是一个值，它的返回值是一个指向新分配类型零值的指针。

### 2.**数组和切片的区别**

- 相同点：
  - 都是一系列用来存放对应数据的集合
- 不同点：
  - 基本：
    - 数组不可改变，定义后只能修改，无法增删
    - 切片可以进行后续操作改变
  - 语法定义：
    - 数组的语法为： var arr  [10]int
    - 切片的语法为： var arr []int
  - 类型：
    - 数组：值类型，进行函数传递值时，通常是值传递，拷贝一份后进行操作
    - 切片：引用类型，函数操作时，针对传递指针进行操作
  - 空间大小：
    - 数组：数组大小为初始值时，默认的长度以及类型进行开辟空间
    - 切片：切片大小默认为24。这是因为切片的结构体只存放三个3个变量
      - 指针，长度，容量
      - 切片可以进行增删值，当超出现有容量后，会在1024容量内进行翻倍，超出后则每次增加1/4







1. 多态
2. 定时任务crontab linux 定时任务写法



1. docker k8s?
2. 算法:树、链表、冒泡排序



### map的key可以有哪些类型?可以用struct interface吗

可:bool, 数字，string, 指针, channel , interface, structs, arrays 

不可:slice， map , function ，没法用 == 来判断

key一定要是**可比较**的类型（可以理解为支持==的操作）：

| 可比较类型                                  | 不可比较类型 |      |
| ------------------------------------------- | ------------ | ---- |
| boolean                                     | slice        |      |
| numeric                                     | map          |      |
| string                                      | func         |      |
| pointer                                     |              |      |
| channel(类型是可比较的)                     |              |      |
| interface                                   |              |      |
| 包含前文类型的array和struct(类型是可比较的) |              |      |
| array(数值类型可比较)                       |              |      |

如果是非法的key类型，会报错：invalid map key type xxx。

value可以是**任意类型**。



### go的并发流程?

**Go语言通过系统的线程来多路派遣这些函数的执行，使得每个用go关键字执行的函数可以运行成为一个单位协程。当一个协程阻塞的时候，调度器就会自动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。而且调度的开销非常小，一颗CPU调度的规模不下于每秒百万次，**这使得我们能够创建大量的goroutine，从而可以很轻松地编写高并发程序，达到我们想要的目的。

 

select 是怎么选择的



进程、线程、协程



linux查看资源消耗



vim跳转回底部



redis是否时多线程



字符串起两个协程按顺序打印



http 版本区别



tcp三次握手四次挥手



mysql索引相关



联查两表中某用户创建的订单前十的用户名



sync.waitgroup的锁



map多线程操作是否线程安全



map的key可以是什么?不能是什么



