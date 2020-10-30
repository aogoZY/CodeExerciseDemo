## Sync.waitgroup



## 前沿

为了支持并发——>引入goroutine——>导致goroutine没执行完main就退出——>使用time.sleep()等待——>时间无法确定——>引入channel作为线程间的通信——>channel占内存开销大——>引入sync.waitgroup



## 意义

能够一直等到所有的goroutine执行完成,并且阻塞主线程的执行,直到所有的goroutine执行完成



## 介绍

waitgroup

三个方法

> - Add()            Add(n)` 把计数器设置为`n
>
> - Done()         Done()` 每次把计数器`-1
>
> - Wait()           wait()` 会阻塞代码的运行，直到计数器地值减为`0

