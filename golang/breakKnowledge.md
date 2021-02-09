# golang知识点碎片

[TOC]

# 1、go的数组、切片,如何扩容?指针引用还是值引用?

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



# 2、go的协程原理?为啥支持高并发?为啥轻量级?

**进程、线程、协程的关系和区别：**

- 进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。
- 线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是的)。
- 协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

**为什么协程比线程轻量？**

- go协程调用跟切换比线程效率高
- go协程占用内存少



# 3、channel 向channel里写、读、已关闭 会发生什么情况

#### 什么是channel，为什么它可以做到线程安全？

Channel是Go中的一个核心类型，可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯(communication),Channel也可以理解是一个先进先出的队列，通过管道进行通信。

Golang的Channel,发送一个数据到Channel 和 从Channel接收一个数据 都是 原子性的。而且Go的设计思想就是:不要通过共享内存来通信，而是通过通信来共享内存，前者就是传统的加锁，后者就是Channel。也就是说，设计Channel的主要目的就是在多任务间传递数据的，这当然是安全的。

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



# 4、go mod vs go vendor

go的包引用的前世今生

(1)gopath

GOPATH 解决了第三方源码依赖的问题，用户比较容易的能够引用一个 Github 上的第三方代码库，进行打包构建。但这引入了一个问题，所有的团队依赖的只互联网上的一份源代码，如果这个代码库被人篡改，删除，你们的构建也会受到影响。

![image-20201031010812759](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010812759.png)

(2)为了解决这个问题，社区里提出了 Go Vendoring 的方法。

![image-20201031010849483](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010849483.png)

但这又引入了新的问题，随着项目的依赖增多，你的代码库可能会越来越大，例如上图，你依赖了 mypkg@v1.2.3， 这个库里有几百兆的图片，当你依赖了这个库，在 pull 代码的时候就大大增加的你的代码库大小。



(3)Athens – Go 依赖管理工具

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

##### go vendor

1. 初始化vendor目录      govendor  init(生成一个vendor目录)
2.  govendor add +external  将GOPATH中本工程使用到的依赖包自动移动到vendor目录中,若本地GOPATH没有依赖包，先go get相应的依赖包

##### go mod

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
- 

# 5、make vs new

new：为所有的类型分配内存，并初始化为零值，返回指针。

make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。

**1.new和make的区别？**

- 相同点：

  - new和make都是用来开辟空间的

- 不同点：

  - new是初始化一个类型的指针,返回的是类型**指针**，而里面的值为默认初始值，只对值类型有效

    返回值是一个指向新分配类型零值的指针

  - make是针对**slice**切片，**map**字典，**chan管道**初始化，并非返回指针，而是对应的类型有效

    

# 6、数组和切片的区别

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



# 7、map的key可以有哪些类型?可以用struct interface吗

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



# 8、go的并发流程?

**Go语言通过系统的线程来多路派遣这些函数的执行，使得每个用go关键字执行的函数可以运行成为一个单位协程。当一个协程阻塞的时候，调度器就会自动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。而且调度的开销非常小，一颗CPU调度的规模不下于每秒百万次，**这使得我们能够创建大量的goroutine，从而可以很轻松地编写高并发程序，达到我们想要的目的。

 

select 是怎么选择的



进程、线程、协程



linux查看资源消耗



vim跳转回底部



redis是否时多线程



字符串起两个协程按顺序打印



# 9、http 版本区别

- H T T P是一种简单无状态，纯文本的客户端服务器协议，他用于在客户端和服务器之间进行数据交换。
- 请求方法是请求行中的第一个单词，他指明了客户端想要对资源执行的操作。

​      http 0.9只有get方法

​      http 1.0添加了post和head方法

​      http 1.1则添加了put、delete、options、 trace、connect五个方法。

- 幂等的请求方法：一个http方法在使用相同的数据进行第二次调用的时候不会对服务器的状态造成任何改变，这个方法称为幂等。eg：put、delete。



# 10、tcp三次握手四次挥手

#### **为什么需要三次握手？**

**为了防止已失效的连接请求报文段突然又传送到了服务端**，造成资源浪费

eg:client发出的第一个连接请求报文段并没有丢失，而是在某个网络结点长时间的滞留了，以致延误到连接释放以后的某个时间才到达server。本来这是一个早已失效的报文段。但server收到此失效的连接请求报文段后，就误认为是client再次发出的一个新的连接请求。于是就向client发出确认报文段，同意建立连接

假设不采用“三次握手”，那么只要server发出确认，新的连接就建立了。由于现在client并没有发出建立连接的请求，因此不会理睬server的确认，也不会向server发送数据。但server却以为新的运输连接已经建立，并一直等待client发来数据。这样，server的很多资源就白白浪费掉了。采用“三次握手”的办法可以防止上述现象发生。例如刚才那种情况，client不会向server的确认发出确认。server由于收不到确认，就知道client并没有要求建立连接。主要目的防止server端一直等待，浪费资源。

#### **为什么关闭连接需要四次挥手呢？**

关闭连接时，当收到对方的FIN报文时，仅仅表示对方不再发送数据了但是还能接收数据，己方也未必全部数据都发送给对方了，所以己方可以立即close，也可以发送一些数据给对方后，再发送FIN报文给对方来表示同意现在关闭连接



# 11、TCP,UDP通信使用场景及区别比较

TCP通信协议

- TCP是面向连接的；
- 每条TCP连接只能由于两个端点，一对一通信；
- TCP提供可靠的交付服务，传输数据无差错，不丢失，不重复，且按时序到达；
- TCP提供全双工通信；
- 面向字节流，TCP根据对方给出的窗口和当前的网络拥塞程度决定一个报文应该包含多少个字节。



UDP通信协议

- 无连接；
- UDP使用尽最大努力交付，不保证可靠性UDP;是面向报文的，UDP对应用层交付下来的报文，既不合并，也不拆分，而是保留报文的边界；
- 应用层交给UDP多长的报文，UDP就照样发送，即一次发送一个报文；
- UDP没有拥塞控制；
- UDP支持一对一，一对多，多对一和多对多的交互通信。
- UDP的首部开销小，只有8字节。

|            | TCP                                                          | UDP                               |
| ---------- | ------------------------------------------------------------ | --------------------------------- |
| 连接性     | 面向连接                                                     | 不需建立连接                      |
| 可靠性     | 可靠性的服务;无差错、不丢失、不重复、按顺序到达(校验和，重传控制，序号标识，滑动窗口、确认应答实现可靠传输) | 尽最大努力交付,不保证可靠性       |
| 工作效率   | 低                                                           | 高,适用高速传播和实时性较高的通信 |
| 连接方式   | 点到点                                                       | 一对一，一对多，多对一和多对多    |
| 对资源要求 | 多                                                           | 少                                |
| 适用场景   | 传输文件                                                     | 会议视频图像                      |



# 12、从输入url到页面加载完成发生了什么

​	   1、浏览器的地址栏输入URL并按下回车。

　　2、浏览器查找当前URL是否存在缓存，并比较缓存是否过期。

　　3、DNS解析URL对应的IP。

　　4、根据IP建立TCP连接（三次握手）。

![img](https://images2015.cnblogs.com/blog/1034346/201703/1034346-20170329145607592-1103856922.png)

　　5、HTTP Reques发起请求。　　

​        完整的HTTP请求包含请求起始行、请求头部、请求主体三部分。

　　6、服务器处理请求，浏览器接收HTTP Response响应。

```
   1xx：指示信息–表示请求已接收，继续处理。

　　2xx：成功–表示请求已被成功接收、理解、接受。

　　3xx：重定向–要完成请求必须进行更进一步的操作。

　　4xx：客户端错误–请求有语法错误或请求无法实现。

　　5xx：服务器端错误–服务器未能实现合法的请求。
```

　　7、渲染页面，构建DOM树。

　　8、关闭TCP连接（四次挥手）。



从网络协议层面来看:

![img](https://images2015.cnblogs.com/blog/1034346/201703/1034346-20170329153945389-2019926409.png)

1、应用层进行DNS解析

　　通过DNS将域名解析成IP地址。在解析过程中，按照`浏览器缓存`、`系统缓存`、`路由器缓存`、`ISP(运营商)DNS缓存`、`根域名服务器`、`顶级域名服务器`、`主域名服务器`的顺序，逐步读取缓存，直到拿到IP地址

2、应用层生成HTTP请求报文

　　接着，应用层生成针对目标WEB服务器的HTTP请求报文，HTTP请求报文包括起始行、首部和主体部分

3、传输层建立TCP连接

​        传输层传输协议分为UDP和TCP两种

　　UDP是无连接的协议，而TCP是可靠的有连接的协议，主要表现在：接收方会对收到的数据进行确认、发送方会重传接收方未确认的数据、接收方会将接收到数据按正确的顺序重新排序，并删除重复的数据、提供了控制拥挤的机制

　　HTTP协议使用的是TCP协议，为了方便通信

(1)HTTP请求报文按序号分为多个报文段**(segment)**，并对每个报文段进行封装。

(2)使用本地一个大于1024以上的随机TCP源端口建立到目的服务器TCP80号端口(HTTPS协议对应的端口号是443)的连接，**TCP源端口和目的端口**加入到报文段中，学名叫协议数据单元(Protocol Data Unit, PDU)。

(3)TCP是一个可靠的传输控制协议，传输层还会加入序列号、确认号、窗口大小、校验和等参数，共添加20字节的头部信息

![http](https://pic.xiaohuochai.site/blog/HTTP_network15.jpg)

构建TCP请求会增加大量的网络时延，常用的优化方式如下所示

　　（1）资源打包，合并请求

　　（2）多使用缓存，减少网络传输

　　（3）使用keep-alive建立持久连接

　　4、网络层使用IP协议来选择路线

​			处理上层传输层的数据段segment，将数据段segment装入数据包**packet**，主要就是添加**源和目的IP地址**，然后发送数据。在数据传输的过程中，IP协议负责选择传送的路线，称为路由功能

![http](https://pic.xiaohuochai.site/blog/HTTP_network12.jpg)

5、数据链路层实现网络相邻结点间可靠的数据通信

　　为了保证数据的可靠传输，把数据包packet封装成帧(**Frame**)，并按顺序传送各帧。由于物理线路的不可靠，发出的数据帧有可能在线路上出错或丢失，于是为每个数据分块计算出CRC(循环冗余检验)，并把CRC添加到帧中，这样接收方就可以通过重新计算CRC来判断数据接收的正确性。一旦出错就重传

　　将数据包packet封装成帧(**Frame**)，包括帧头和帧尾。帧尾是添加被称做CRC的循环冗余校验部分。帧头主要是添加数据链路层的地址，即数据链路层的源地址和目的地址，即**网络相邻结点间的源MAC地址和目的MAC地址**

　　6、物理层传输数据

　　数据链路层的帧(Frame)转换成二进制形式的比特(Bit)流，从网卡发送出去，再把比特转换成电子、光学或微波信号在网络中传输

上面的6个步骤可总结为：DNS解析URL地址、生成HTTP请求报文、构建TCP连接、使用IP协议选择传输路线、数据链路层保证数据的可靠传输、物理层将数据转换成电子、光学或微波信号进行传输



mysql索引相关





联查两表中某用户创建的订单前十的用户名



sync.waitgroup的锁



# 13、map多线程操作是否线程安全

map：不是线程安全的。在同一时间段内，让不同 goroutine 中的代码，对同一个字典进行读写操作是不安全
的。字典值本身可能会因这些操作而产生混乱，相关的程序也可能会因此发生不可预知的问题。

sync.Map:并发安全的字典类型sync.Map。这个字典类型提供了一些常用的键值存取操作方法，并保证了这些操作的并发安全。同时，它的存、取、删等操作都可以基本保证在常数时间内执行完毕。换句话说，它们的算法复杂度与map类型一样都是O(1)的。

```
 var ma sync.Map// 该类型是开箱即用，只需要声明既可
    ma.Store("key", "value") // 存储值
    ma.Delete("key") //删除值
    ma.LoadOrStore("key", "value")// 获取值，如果没有则存储
    fmt.Println(ma.Load("key"))//获取值
    
    //遍历
    ma.Range(func(key, value interface{}) bool {
        fmt.Printf("key:%s ,value:%s \n", key, value)
        //如果返回：false，则退出循环，
        return true
    })
```

map的key可以是什么?不能是什么

```
键类型的值之间必须可以施加操作符==和!=。换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。
```

判断那些类型作为字典的键比较合适

```
map查找的流程中,“把键值转换为哈希值”以及“把要查找的键值与哈希桶中的键值做对比”是两个重要且比较耗时的操作。可以说：求哈希和判等操作的速度越快，对应的类型就越适合作为键类型。
```



# 14、进程

进程，直观点说，保存在硬盘上的程序运行以后，会在内存空间里形成一个独立的内存体，这个内存体**有自己独立的地址空间，有自己的堆**，上级挂靠单位是操作系统。**操作系统会以进程为单位，分配系统资源（CPU时间片、内存等资源），进程是资源分配的最小单位**。

线程,有时被称为轻量级进程(Lightweight Process，LWP），是操作系统调度（CPU调度）执行的最小单位。

#### 进程和线程的区别与联系

【区别】：

- **调度**：**线程作为调度和分配的基本单位，进程作为拥有资源的基本单位**；
- **并发性**：**不仅进程之间可以并发执行，同一个进程的多个线程之间也可并发执行**；
- **拥有资源**：**进程是拥有资源的一个独立单位，线程不拥有系统资源**，但可以访问隶属于进程的资源。进程所维护的是程序所包含的资源（静态资源）， 如：**地址空间，打开的文件句柄集，文件系统状态，信号处理handler等**；线程所维护的运行相关的资源（动态资源），如：**运行栈，调度相关的控制信息，待处理的信号集等**；
- **系统开销**：在创建或撤消进程时，由于系统都要为之分配和回收资源，导致系统的开销明显大于创建或撤消线程时的开销。但是进程有独立的地址空间，一个进程崩溃后，在保护模式下不会对其它进程产生影响，而线程只是一个进程中的不同执行路径。线程有自己的堆栈和局部变量，但线程之间没有单独的地址空间，一个进程死掉就等于所有的线程死掉，所以**多进程的程序要比多线程的程序健壮，但在进程切换时，耗费资源较大，效率要差一些**。

【联系】：

- **一个线程只能属于一个进程，而一个进程可以有多个线程，但至少有一个线程**；
- 资源分配给进程，同一进程的所有线程共享该进程的所有资源；
- 处理机分给线程，即**真正在处理机上运行的是线程**；
- 线程在执行过程中，需要协作同步。不同进程的线程间要利用消息通信的办法实现同步。

#### **说说进程、线程、协程之间的区别？**

进程是资源的分配和调度的一个独立单元，而线程是CPU调度的基本单元；

同一个进程中可以包括多个线程；一个进程至少包括一个线程；

进程结束后它拥有的所有线程都将销毁，而线程的结束不会影响同个进程中的其他线程的结束；

线程共享整个进程的资源（寄存器、堆栈、上下文）线程中执行时一般都要进行同步和互斥，因为他们共享同一进程的所有资源；

 

进程是资源分配的单位 

线程是操作系统调度的单位 

进程切换需要的资源很最大，效率很低 
线程切换需要的资源一般，效率一般 
协程切换任务资源很小，效率高 

进程拥有自己独立的堆和栈，既不共享堆，亦不共享栈，进程由操作系统调度。（全局变量保存在堆中，局部变量及函数保存在栈中）

线程拥有自己独立的栈和共享的堆，共享堆，不共享栈，线程亦由操作系统调度(标准线程是这样的)。

协程和线程一样共享堆，不共享栈，协程由程序员在协程的代码里显示调度。

一个应用程序一般对应一个进程，一个进程一般有一个主线程，还有若干个辅助线程，线程之间是平行运行的，在线程里面可以开启协程，让程序在特定的时间内运行。

协程和线程的区别是：协程避免了无意义的调度，由此可以提高性能，但也因此，程序员必须自己承担调度的责任，同时，协程也失去了标准线程使用多CPU的能力。



#### 一个形象的例子解释进程和线程的区别

![在这里插入图片描述](http://blog.chinaunix.net/attachment/201310/23/29270628_1382541951nJe7.jpg)

这副图是一个双向多车道的道路图，假如我们**把整条道路看成是一个“进程”的话**，那么图中由白色虚线分隔开来的**各个车道就是进程中的各个“线程”了**。

- **这些线程(车道)共享了进程(道路)的公共资源(土地资源)**。
- 这些线程(车道)必须依赖于进程(道路)，也就是说，**线程不能脱离于进程而存在(就像离开了道路，车道也就没有意义了)**。
- **这些线程(车道)之间可以并发执行(各个车道你走你的，我走我的)，也可以互相同步(某些车道在交通灯亮时禁止继续前行或转弯，必须等待其它车道的车辆通行完毕)**。
- **这些线程(车道)之间依靠代码逻辑(交通灯)来控制运行，一旦代码逻辑控制有误(死锁，多个线程同时竞争唯一资源)，那么线程将陷入混乱，无序之中**。
- **这些线程(车道)之间谁先运行是未知的**，只有在线程刚好被分配到CPU时间片(交通灯变化)的那一刻才能知道。

![在这里插入图片描述](http://5b0988e595225.cdn.sohucs.com/images/20180622/6765e36cc4604fba897976638af03524.jpeg)

协程，是一种比线程更加轻量级的存在，协程不是被操作系统内核所管理，而完全是由程序所控制（也就是在用户态执行）。这样带来的好处就是性能得到了很大的提升，不会像线程切换那样消耗资源。

  子程序，或者称为函数，在所有语言中都是层级调用，比如A调用B，B在执行过程中又调用了C，C执行完毕返回，B执行完毕返回，最后是A执行完毕。所以子程序调用是通过栈实现的，**一个线程就是执行一个子程序**。子程序调用总是一个入口，一次返回，调用顺序是明确的。而协程的调用和子程序不同。

  **协程在子程序内部是可中断的，然后转而执行别的子程序，在适当的时候再返回来接着执行**。



# 协程同步的方式 waitgroup和context区别 如何处理异常 defer



# tcp三次握手四次挥手 可靠性如何保证

# 网络不可达如何排查，例如我当前打不开qq.com

# tcp、udp区别、进程线程区别

1. go map slice 实现（[源码](https://www.nowcoder.com/jump/super-jump/word?word=源码)分析以及slice内存泄漏分析）
2. go 内存逃逸分析（分析了栈帧，讲五种例子，描述堆栈优缺点，点头）

1. go 协程怎么切换的



go的调度
go struct能不能比较

可以,也不可以

比如struct例有一个map就不行了,因为map是无法比较的

go中slice、map、function都无法比较

结论:

相同结构,只要成员类型可比较,则能比较

不同结构,如果能相互转化也能比较.前提是成员都是可比较的.



go defer（for defer）

栈操作，后进先出。



# select可以用于什么

select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

1. goroutine超时设置，防止goroutine一直执行导致内存不释放等问题。
2. 判断channel是否已满或空。如实现一个池线程，当channel已被写满，暂无空闲worker在进行读取，进入default，返回一个暂无可分配资源错误。



context包的用途
client如何实现长连接
主协程如何等其余协程完再操作
slice，len，cap，共享，扩容
map如何顺序读取
实现set
实现消息队列（多生产者，多消费者）
大文件排序
基本排序，哪些是稳定的

快速排序、希尔排序、堆排序、直接选择排序不是稳定的排序算法。

基数排序、冒泡排序、直接插入排序、折半插入排序、归并排序是稳定的排序算法。



http get跟head
http 401,403
http keep-alive
http能不能一次连接多次请求，不等后端返回
tcp与udp区别，udp优点，适用场景
time-wait的作用
数据库如何建索引
孤儿进程，僵尸进程
死锁条件，如何避免
linux命令，查看端口占用，cpu负载，内存占用，如何发送信号给一个进程
git文件版本，使用顺序，merge跟rebase

>
> 爬虫如何做的鉴权吗
> 怎么实现的分布式爬虫
> 电商系统图片多会造成带宽过高，如何解决
> mysql底层有哪几种实现方式
> channel底层实现
> java nio和go 区别
> 读写锁底层是怎么实现的
> go-micro 微服务架构怎么实现水平部署的，代码怎么实现
> micro怎么用
> 怎么做服务发现的
> mysql索引为什么要用B+树？
> mysql语句性能评测？
> 服务发现有哪些机制当go服务部署到线上了，发现有内存泄露，该怎么处理

## goalng相关

Q：context作用，原理，超时控制

**A: golang context的理解，context主要用于父子任务之间的同步取消信号，本质上是一种协程调度的方式。另外在使用context时有两点值得注意：上游任务仅仅使用context通知下游任务不****再****需要，但不会直接干涉和中断下游任务的执****行****，由下游任务自行决定后续的处理操作，也就是说context的取消操作是无侵入的；context是线程安全的，因为context本身是不可变的（immutable），因此可以放心地在多个协程中传递使用。**

**内存方面问题**

go语言有什么优点和缺点

**A: 优势：容易学习，生产力，并发，动态语法。劣势：包管理，错误处理，缺乏框架。**

Go GC算法，三色标记法描述

go的gc是基于 `标记-清扫`算法

### 标记-清扫(Mark And Sweep)算法

此算法主要有两个主要的步骤：

- 标记(Mark phase)
- 清除(Sweep phase)

第一步，找出不可达的对象，然后做上标记。

第二步，回收标记好的对象。

操作非常简单，但是有一点需要额外注意：mark and sweep算法在执行的时候，需要程序暂停！即 `stop the world`。

标记-清扫(Mark And Sweep)算法这种算法虽然非常的简单，但是还存在一些问题：

- STW，stop the world；让程序暂停，程序出现卡顿。
- 标记需要扫描整个heap
- 清除数据会产生heap碎片

### 三色并发标记法

1.首先将程序创建的对象全部标记为白色

2.gc开始扫描，并将可达的对象标记为灰色

3.再从灰色对象中找到其引用的对象，将其标记为灰色，将自身标记成黑色

重复以上2、3步骤，直至没有灰色对象

4.对所有白色对象进行清除



作者：98k_sw
链接：https://www.jianshu.com/p/cfc669f83eaa
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

#### Golang的内存模型，为什么小对象多了会造成gc压力。

通常小对象过多会导致GC三色法消耗过多的GPU。优化思路是，减少对象分配.

gc的过程一共分为四个阶段：

1. 栈扫描（开始时STW）
2. 第一次标记（并发）
3. 第二次标记（STW）
4. 清除（并发）

# go中多态的实现

Golang中的多态可以通过接口来实现。

定义接口的所有方法的任何类型都表示隐式实现该接口。类型接口的变量可以保存实现该接口的任何值。接口的这个属性用于实现GO的多态性。



map底层实现机制(无序、底层机制)

线程、协程、go程区别

slice和链表的区别?面试官想要知道的是:使用场景、优缺点、不同之处、原理、底层结构,从时间复杂度空间复杂度去分析

redis

线程的多种方式?

mpg模型?goroutie的实现机制

二叉树、排序二叉树、应用场景?

数据库的四大特性?

### 



移动互联网

海量用户(高并发、分布式、微服务server mesh服务治理,服务异常、失败的处理(网格计算))

异构系统、服务治理、服务网络(网格能力、集群声明能力)

golang 

grpc

