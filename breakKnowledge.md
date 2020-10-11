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









1. docker k8s?

2. 算法:树、链表、冒泡排序

   

   

   

   ### 



1. 多态
2. 定时任务crontab linux 定时任务写法
3. go mod和go vendor