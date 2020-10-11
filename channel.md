[TOC]



# channel

## channel状态

channel作为go的一种基本数据类型，它有3种基本状态：nil、open、closed：

```Go
/* nil channel */
var ch = chan string // A channel is in a nil state when it is declared to its zero value
ch = nil // A channel can be placed in a nil state
 
/* open channel */
ch := make(chan string) // A channel is in a open state when it’s made using the built-in function make.
 
/* closed channel */
close(ch) // A channel is in a closed state when it’s closed using the built-in function close.
```

当channel处于这3种不同的状态时，对于channel上的操作也会有不同的行为，理解这些行为对于正确的使用channel非常重要。 

![img](https://img-blog.csdnimg.cn/2020031220024044.jpg?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3UwMTIyOTk1OTQ=,size_16,color_FFFFFF,t_70)

## 创建channel

channel是**引用**类型,需要实用make来创建channel

```
ch:=make(chan int,3)  带缓存
ch:=make(chan string) 不带缓存
```

## channel操作

### 放元素

- 我们可以使用`<-`符号指向`channel`来将元素放入channel中
- *注意向通道中传值必须要求该通道还有容量(缓冲),而且通道不能关闭*
- *对于无缓冲的或者缓冲已经满了的channel不可以轻易的传入值,必须要有goroutine同时在取元素才可以放入*

向一个有缓冲,非满的channel传值

```go
c := make(chan int, 1) // 定义一个带有一个缓冲的通道
c <- 1				   // 向通道中传入一个1,正常
```

**向一个有缓冲,满的channel传值**

```go
c := make(chan int, 1) // 定义一个带有一个缓冲的通道
c <- 1				   // 向通道中传入一个值,这个值传入后填满了该通道
c <- 2                 // 再向通道中传入一个值,报错!!!!!!!!!!!!!!!!!!!!!!!!!!!
```

**向一个无缓冲的channel传值**

```go
c := make(chan int) // 定义一个无缓冲通道
c <- 1              // 向无缓冲通道传值,报错!!!!!!!!!!!!!!!!!!!!!!!!!!!!
```

其实上述两种错误均是由于通道满了而引起的(无缓冲的通道可以看成是缓冲为0的通道)

### 取元素

- 我们可以使用`<-`符号指向`变量`来将channel中的元素放`变量`中

  此时可以接收两个值一个数值一个状态

  ```go
  v, ok := <-c   // c是通道,v是取到的值,ok是状态,正常时是true,从关闭的空通道取值是false
  ```

- 可以通过`range`取值

- *注意从通道中取值必须要求该通道还有值*

- *对于无缓冲的或者缓冲已经空了的channel不可以轻易的取出值,必须要同时在放元素才可以取出*

- 可以向已经关闭的通道取值

**ok为true的例子**

```go
c := make(chan int, 1)
c <- 1
a, ok := <-c
fmt.Println(a, ok)  // 输出 1 true
```

**ok为false的例子**

```go
c := make(chan int, 1)
close(c)
a, ok := <-c
fmt.Println(a, ok)  // 输出 0 false
```

**ok的应用--循环取值**

```go
for {
    v, ok := <- c
    if !ok {
        break
    }
    fmt.Println(v)
}
```

**range取值**

```go
func rangeFunc() {
	c := make(chan int, 10)
	for i:=0; i< 10; i++ {
		c <- i
	}
	close(c)

	for v := range c{
		fmt.Println(v)
	}
}
```

**向一个有缓冲,非空的channel取值**

```go
c := make(chan int, 1) // 定义一个带有一个缓冲的通道
c <- 1				   // 向通道中传入一个1,使通道非空
i := <-c			   // 从通道中取出一个值赋给变量i
// 如果只是想取出值而不想对该值做任何其他操作,可以这么写    <-c   左边省略接收者
```

**向一个有缓冲,空的channel取值**

```go
c := make(chan int, 1) // 定义一个带有一个缓冲的通道
<-c                    // 向空通道中取出一个值,报错!!!!!!!!!!!!!!!!!!!!!!!!!!!
```

**向一个无缓冲的channel取值**

```go
c := make(chan int) // 定义一个无缓冲通道
<-c                 // 向无缓冲通道取值,报错!!!!!!!!!!!!!!!!!!!!!!!!!!!!
```

其实上述两种错误均是由于通道空了而引起的(无缓冲的通道可以看成是缓冲为0的通道)

### 关闭通道

对于一个通道我们可以使用`close`内置函数来进行关闭,关闭后的通道具有以下特点

- 向一个已经关闭的通道**发送值是不允许的**,会报错
- 从一个已经关闭但是里面还有值的通道取值是允许的,可以正常获取到值
- 从一个已经关闭但是为空的通道取值是允许的,会获取通道类型元素的零值
- 不可以再次关闭一个已经关闭的通道,会报错
- 已经关闭的通道**无法再次打开**

**例子1: 向一个已经关闭的通道发送值**

```go
c := make(chan int, 1)
close(c)
c <- 1		// 报错!!!!!!!!!!!
```

**例子2: 从一个已经关闭但是里面还有值的通道取值**

```go
c := make(chan int, 1)
c <- 1
close(c)
a := <-c
fmt.Println(a)	// 输出 1
```

**例子3: 从一个已经关闭但是为空的通道取值**

```go
c := make(chan int, 1)
close(c)
a := <-c
fmt.Println(a)  // 输出 0 
```

**例子4: 关闭一个已经关闭的通道**

```go
c := make(chan int, 1)
close(c)
close(c) // 报错: panic: close of closed channel
```

## 单向通道

- 在函数中使用通道时我们可以限制其为只读通道或者只写通道

**定义只读通道的例子**

```go
func doWork(i <-chan int) {
	<-i		 // 只能取值
	//i <- 1 // 存值操作将不被允许
}
```

**定义只写通道的例子**

```go
func doWork(i chan <- int) {
	i <- 1      // 只能存值
	//<-i		// 取值操作将不被允许
}
```