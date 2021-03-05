

# golang知识点碎片

[TOC]

# 1、数组 vs 切片？

- 相同点：
  - 都是一系列用来存放对应数据的集合
- 不同点：
  - 初始化方式：
    - 数组长度固定(需指定大小)，定义后**只能修改，无法增删**
    - 切片可以进行后续操作改变(切片是动态的数组)
  - 语法定义方式：
    - 数组的语法为： var arr  [10]int
    - 切片的语法为： var arr []int
  - 传递方式：
    - 数组：**值类型**，进行函数传递值时，通常是值传递，拷贝一份后进行操作
    - 切片：**引用类型**，函数操作时，针对传递指针进行操作
  - 空间大小：
    - 数组：数组大小为初始值时，默认的长度以及类型进行开辟空间
    - 切片：切片大小默认为24。这是因为切片的结构体只存放三个3个变量
      - 指针，长度，容量
      - 切片可以进行增删值，当超出现有容量后，会在1024容量内进行翻倍，超出后则每次增加1/4

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

Channel是Go中的一个核心类型，可以把它看成一个管道，通过它并发核心单元就可以发送或者接收数据进行通讯，Channel也可以理解是一个先进先出的队列，通过管道进行通信。

不同于传统的多线程并发模型使用共享内存来实现线程间通信的方式，golang 的哲学是通过 channel 进行协程(goroutine)之间的通信来实现数据共享：前者就是传统的加锁，后者就是Channel。也就是说，设计Channel的主要目的就是在多任务间传递数据的，这当然是安全的。

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

### go path

GOPATH 解决了第三方源码依赖的问题，用户比较容易的能够引用一个 Github 上的第三方代码库，进行打包构建。但这引入了一个问题，所有的团队依赖的只互联网上的一份源代码，如果这个代码库被人篡改，删除，你们的构建也会受到影响。

![image-20201031010812759](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010812759.png)

### go vendor

为了解决这个问题，社区里提出了 Go Vendoring 的方法。

![image-20201031010849483](/Users/zhouyang/Library/Application Support/typora-user-images/image-20201031010849483.png)

但这又引入了新的问题，随着项目的依赖增多，你的代码库可能会越来越大，例如上图，你依赖了 mypkg@v1.2.3， 这个库里有几百兆的图片，当你依赖了这个库，在 pull 代码的时候就大大增加的你的代码库大小。



### go mod

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



### 使用方式

#### go vendor

1. 初始化vendor目录      govendor  init(生成一个vendor目录)
2.  govendor add +external  将GOPATH中本工程使用到的依赖包自动移动到vendor目录中,若本地GOPATH没有依赖包，先go get相应的依赖包

#### go mod

(版本go1.13)

1. 设置环境

   ```
   go env -w go111module=on
   go env -w goproxy=yun.paic.com.cn
   go env sumdb=off 不校验包的hash值
   ```

2. go mod init

   生成go.sum & go.mod

**go.sum** 详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值以备 Go 在今后的操作中保证项目所依赖的那些模块版本不会被篡改。

**go.mod** 是启用了 Go moduels 的项目所必须的最重要的文件，它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:(包引用路径+版本号)

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

  

# 5、make vs new

new：为所有的类型分配内存，并初始化为零值，返回指针。

make：只能为 slice，map，chan 分配内存，并初始化，返回的是类型。

**new和make的区别？**

- 相同点：

  - new和make都是用来开辟空间的

- 不同点：

  - new是初始化一个类型的指针,返回的是**指针**，值为默认初始值，只对值类型有效

    ```
    func new(Type) *Type
    ```

  - make**仅**用于**分配**和**初始化** slice、map 以及 channel 类型的对象，三种类型都是结构。返回值为类型，而不是指针。

    ```
    func make(t Type, size ...IntegerType) Type
    ```

    - slice 源码结构：

    ```Go
    type slice struct {
        array unsafe.Pointer  //指针
        len   int             //长度
        cap   int             //容量
    }
    ```

    - map 源码结构：

    ```Go
    // A header for a Go map.
    type hmap struct {
    	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
    	// Make sure this stays in sync with the compiler's definition.
    	count     int // # live cells == size of map.  Must be first (used by len() builtin)
    	flags     uint8
    	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
    	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
    	hash0     uint32 // hash seed
    
    	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
    	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
    	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)
    
    	extra *mapextra // optional fields
    }
    ```

    - channel 源码结构：

    ```Go
    type hchan struct {
    	qcount   uint           // total data in the queue
    	dataqsiz uint           // size of the circular queue
    	buf      unsafe.Pointer // points to an array of dataqsiz elements
    	elemsize uint16
    	closed   uint32
    	elemtype *_type // element type
    	sendx    uint   // send index
    	recvx    uint   // receive index
    	recvq    waitq  // list of recv waiters
    	sendq    waitq  // list of send waiters
    
    	// lock protects all fields in hchan, as well as several
    	// fields in sudogs blocked on this channel.
    	//
    	// Do not change another G's status while holding this lock
    	// (in particular, do not ready a G), as this can deadlock
    	// with stack shrinking.
    	lock mutex
    }
    ```

  

  







# 6、map的key可以有哪些类型?

结论:

相同结构,只要成员类型可比较,则能比较

不同结构,如果能相互转化也能比较.前提是成员都是可比较的.



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



# 7、go的并发流程?

**Go语言通过系统的线程来多路派遣这些函数的执行，使得每个用go关键字执行的函数可以运行成为一个单位协程。**

**当一个协程阻塞的时候，调度器就会自动把其他协程安排到另外的线程中去执行，从而实现了程序无等待并行化运行。**

**而且调度的开销非常小，一颗CPU调度的规模不下于每秒百万次，**这使得我们能够创建大量的goroutine，从而可以很轻松地编写高并发程序，达到我们想要的目的。



# 8、map多线程操作是否线程安全

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



# 9、map的key可以是什么?不能是什么

```
键类型的值之间必须可以施加操作符==和!=。换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。
```

判断那些类型作为字典的键比较合适

```
map查找的流程中,“把键值转换为哈希值”以及“把要查找的键值与哈希桶中的键值做对比”是两个重要且比较耗时的操作。可以说：求哈希和判等操作的速度越快，对应的类型就越适合作为键类型。
```



# 10、select可以用于什么

select是Go中的一个控制结构，类似于switch语句，用于**处理异步IO操作**。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

1. goroutine超时设置，防止goroutine一直执行导致内存不释放等问题。
2. 判断channel是否已满或空。如实现一个池线程，当channel已被写满，暂无空闲worker在进行读取，进入default，返回一个暂无可分配资源错误。







time-wait的作用
死锁条件，如何避免
mysql底层有哪几种实现方式
channel底层实现
mysql索引为什么要用B+树？
mysql语句性能评测？
服务发现有哪些机制当go服务部署到线上了，发现有内存泄露，该怎么处理

# 11、context作用，原理，超时控制

**A: golang context的理解，context主要用于父子任务之间的同步取消信号，本质上是一种协程调度的方式。另外在使用context时有两点值得注意：上游任务仅仅使用context通知下游任务不再需要，但不会直接干涉和中断下游任务的执行**，由下游任务自行决定后续的处理操作，也就是说context的取消操作是无侵入的

context是线程安全的，因为context本身是不可变的（immutable），因此可以放心地在多个协程中传递使用。



# 12、标记-清扫(Mark And Sweep)算法

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

#### Golang的内存模型，为什么小对象多了会造成gc压力。

通常小对象过多会导致GC三色法消耗过多的GPU。优化思路是，减少对象分配.

gc的过程一共分为四个阶段：

1. 栈扫描（开始时STW）
2. 第一次标记（并发）
3. 第二次标记（STW）
4. 清除（并发）

# 13、go中多态的实现

Golang中的多态可以通过接口来实现。

定义接口的所有方法的任何类型都表示隐式实现该接口。类型接口的变量可以保存实现该接口的任何值。接口的这个属性用于实现GO的多态性。



slice和链表的区别?面试官想要知道的是:使用场景、优缺点、不同之处、原理、底层结构,从时间复杂度空间复杂度去分析

二叉树、排序二叉树、应用场景?







