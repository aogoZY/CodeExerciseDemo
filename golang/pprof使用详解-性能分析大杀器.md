# 				pprof--你想要的数据分析我都有

## PProf干啥的

Golang**自带**的一款开箱即用的**性能监控**和**分析**工具。

pprof开启后，每隔一段时间(10ms)就会收集当前的堆栈信息，获取各个函数占用的CPU以及内存资源，然后通过对这些采样数据进行分析，形成一个性能分析报告。



## PProf可以做什么

- CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置
- Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
- Block Profiling：阻塞分析，记录 goroutine 阻塞等待同步（包括定时器通道）的位置
- Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况



## PProf如何使用

### 一、安装

下载Gin框架专用pprof包并在项目中注册函数

1、go get github.com/gin-contrib/pprof

2、pprof.Register(engine)

![image-20210220141620230](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220141620230.png)

3、浏览器访问http://127.0.0.1:8001/debug/pprof/

![image-20210220141844569](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220141844569.png)

### 二、指标分析

| 类型         | 描述                                      |
| ------------ | ----------------------------------------- |
| allocs       | **内**存分配情况的采样信息                |
| blocks       | **阻塞**操作情况的采样信息                |
| cmdline      | 显示程序启动**命令参数**及其参数          |
| goroutine    | 显示当前所有**协程**的堆栈信息            |
| heap         | **堆**上的内存分配情况的采样信息          |
| mutex        | **锁**竞争情况的采样信息                  |
| profile      | **cpu**占用情况的采样信息，点击会下载文件 |
| threadcreate | 系统**线程**创建情况的采样信息            |
| trace        | 程序**运行跟踪**信息                      |

### 三、使用go tool pprof采集数据

使用pprof的三种方式

#### 1、通过 Web 界面

 `http://127.0.0.1:8001/debug/pprof/`

#### 2、**通过交互式终端使用**

获取当前协程的堆栈信息,采集协程数据并持续20S

```
go tool pprof --seconds 20 http://localhost:8001/debug/pprof/goroutine

或者采用如下写法

go tool pprof http://localhost:8001/debug/pprof/goroutine?second=20
```

![image-20210220143504002](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220143504002.png)

查看堆

```
go tool pprof http://localhost:8001/debug/pprof/heap
```

查看 30-second CPU信息:

```
go tool pprof http://localhost:8001/debug/pprof/profile
```

#### 3、PProf 可视化界面

法一：将上述文件以8080端口对外提供web访问页面

```go
go tool pprof -http=:8080 /Users/zhouyang/pprof/pprof.goroutine.003.pb.gz
```

法二：(其实是默认生成一个svg文件，使用web指令需自己设置为自动通过浏览器打开)

```
go tool pprof /Users/zhouyang/pprof/pprof.goroutine.003.pb.gz
$ (pprof) web
```

![image-20210220160009304](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220160009304.png)

访问localhost：8080端口，可以看到可视化的图形页面，包括火焰图、调用链关系图等。

![image-20210220164123832](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220164123832.png)

![image-20210220160102615](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220160102615.png)







若是出现了Could not execute dot; may need to install graphviz 需要下载安装graphviz。

```
brew install graphviz # for macos
apt-get install graphviz # for ubuntu
yum install graphviz # for centos
```

查看版本

```
dot -v
```

![image-20210220164224939](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220164224939.png)

###  四、pprof命令行常见指令

```
go tool pprof <binary> <source>
inary：代表二进制文件路径。
source：代表生成的分析数据来源，可以是本地文件（前文生成的cpu.prof），也可以是http地址
```

#### top

`top`默认查看程序中占用cpu前10位的函数。

`top 3` 可以查看程序中占用CPU前三位的函数。

最后一列为函数名称，其他各项内容意义如下：

- flat:当前函数占用CPU的耗时
- flat%:当前函数占用CPU的耗时百分比
- sum%:函数占用CPU的累积耗时百分比
- cum：当前函数+调用当前函数的占用CPU总耗时
- cum%: 当前函数+调用当前函数的占用CPU总耗时百分比

![](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220144318099.png)

#### list

我们还可以使用`list 函数名`命令查看具体的函数分析

![image-20210220145102444](/Users/zhouyang/Library/Application Support/typora-user-images/image-20210220145102444.png)

#### pdf

`pdf`命令可以生成可视化的pdf文件。

#### help

`help`命令可以提供所有pprof支持的命令说明。

#### web

以页面浏览器形式打开文件



## PProf实战分析

三步走：

1、查看你关注的指标项

```
go tool pprof http://localhost:6060/debug/pprof/{填上你想查看的内容}
```

```
top->list FuncName->web
```

总的思路就是通过`top` 和`web` 找出关键函数，再通过`list Func` 查看函数代码，找到关键代码行并确认优化方案



## 参考文章：

[golang pprof 实战](https://blog.wolfogre.com/posts/go-ppof-practice/#%E4%BD%BF%E7%94%A8-pprof)

[Golang 大杀器之性能剖析 PProf](https://www.jianshu.com/p/4e4ff6be6af9)

[pprof源码库](https://github.com/gin-contrib/pprof)

