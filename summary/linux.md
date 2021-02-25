linux

# 进程管理

### 查看进程

#### 1. ps

查看某个时间点的进程信息。

示例：查看自己的进程

```
## ps -l
```

示例：查看系统所有进程

```
## ps aux
```

示例：查看特定的进程

```
## ps aux | grep threadx
```

#### 2. pstree

查看进程树。

示例：查看所有进程树

```
## pstree -A
```

#### 3. top

实时显示进程信息。

示例：两秒钟刷新一次

```
## top -d 2
```

#### 4. netstat

查看占用端口的进程

示例：查看特定端口的进程

```
## netstat -anp | grep port
```

### 进程状态

| 状态 | 说明                                                         |
| ---- | ------------------------------------------------------------ |
| R    | running or runnable (on run queue) 正在执行或者可执行，此时进程位于执行队列中。 |
| D    | uninterruptible sleep (usually I/O) 不可中断阻塞，通常为 IO 阻塞。 |
| S    | interruptible sleep (waiting for an event to complete) 可中断阻塞，此时进程正在等待某个事件完成。 |
| Z    | zombie (terminated but not reaped by its parent) 僵死，进程已经终止但是尚未被其父进程获取信息。 |
| T    | stopped (either by a job control signal or because it is being traced) 结束，进程既可以被作业控制信号结束，也可能是正在被追踪。 |



[![img](https://camo.githubusercontent.com/ae6622fbe9bed2d92a67e2cac0bbcb86fcf5e015a789aab89faf6b5e0fdb41e5/68747470733a2f2f63732d6e6f7465732d313235363130393739362e636f732e61702d6775616e677a686f752e6d7971636c6f75642e636f6d2f32626162343132372d336537642d343863632d393134652d3433366265383539666230352e706e67)](https://camo.githubusercontent.com/ae6622fbe9bed2d92a67e2cac0bbcb86fcf5e015a789aab89faf6b5e0fdb41e5/68747470733a2f2f63732d6e6f7465732d313235363130393739362e636f732e61702d6775616e677a686f752e6d7971636c6f75642e636f6d2f32626162343132372d336537642d343863632d393134652d3433366265383539666230352e706e67)



### SIGCHLD

当一个子进程改变了它的状态时（停止运行，继续运行或者退出），有两件事会发生在父进程中：

- 得到 SIGCHLD 信号；
- waitpid() 或者 wait() 调用会返回。

其中子进程发送的 SIGCHLD 信号包含了子进程的信息，比如进程 ID、进程状态、进程使用 CPU 的时间等。

在子进程退出时，它的进程描述符不会立即释放，这是为了让父进程得到子进程信息，父进程通过 wait() 和 waitpid() 来获得一个已经退出的子进程的信息。



### wait()

```
pid_t wait(int *status)
```

父进程调用 wait() 会一直阻塞，直到收到一个子进程退出的 SIGCHLD 信号，之后 wait() 函数会销毁子进程并返回。

如果成功，返回被收集的子进程的进程 ID；如果调用进程没有子进程，调用就会失败，此时返回 -1，同时 errno 被置为 ECHILD。

参数 status 用来保存被收集的子进程退出时的一些状态，如果对这个子进程是如何死掉的毫不在意，只想把这个子进程消灭掉，可以设置这个参数为 NULL。

### waitpid()

```
pid_t waitpid(pid_t pid, int *status, int options)
```

作用和 wait() 完全相同，但是多了两个可由用户控制的参数 pid 和 options。

pid 参数指示一个子进程的 ID，表示只关心这个子进程退出的 SIGCHLD 信号。如果 pid=-1 时，那么和 wait() 作用相同，都是关心所有子进程退出的 SIGCHLD 信号。

options 参数主要有 WNOHANG 和 WUNTRACED 两个选项，WNOHANG 可以使 waitpid() 调用变成非阻塞的，也就是说它会立即返回，父进程可以继续执行其它任务。

### 孤儿进程

一个父进程退出，而它的一个或多个子进程还在运行，那么这些子进程将成为孤儿进程。

孤儿进程将被 init 进程（进程号为 1）所收养，并由 init 进程对它们完成状态收集工作。

由于孤儿进程会被 init 进程收养，所以孤儿进程不会对系统造成危害。

### 僵尸进程

一个子进程的进程描述符在子进程退出时不会释放，只有当父进程通过 wait() 或 waitpid() 获取了子进程信息后才会释放。如果子进程退出，而父进程并没有调用 wait() 或 waitpid()，那么子进程的**进程描述符**仍然保存在系统中，这种进程称之为僵尸进程。

僵尸进程通过 ps 命令显示出来的状态为 Z（zombie）。

系统所能使用的进程号是有限的，如果产生大量僵尸进程，将因为没有可用的进程号而导致系统不能产生新的进程。

要消灭系统中大量的僵尸进程，只需要将其父进程杀死，此时僵尸进程就会变成孤儿进程，从而被 init 进程所收养，这样 init 进程就会释放所有的僵尸进程所占有的资源，从而结束僵尸进程。

### **守护进程**

Linux Daemon（守护进程）是运行在后台的一种特殊进程。它独立于控制终端并且周期性地执行某种任务或等待处理某些发生的事件。它不需要用户输入就能运行而且提供某种服务。Linux系统的大多数服务器就是通过守护进程实现的。常见的守护进程包括系统日志进程syslogd、 web服务器httpd、邮件服务器sendmail和数据库服务器mysqld等。

守护进程一般在系统启动时开始运行，除非强行终止，否则直到系统关机都保持运行。守护进程经常以超级用户（root）权限运行，因为它们要使用特殊的端口（1-1024）或访问某些特殊的资源。

一个守护进程的父进程是init进程，因为它真正的父进程在fork出子进程后就先于子进程exit退出了，所以它是一个由init继承的孤儿进程。守护进程是非交互式程序，没有控制终端，所以任何输出，无论是向标准输出设备stdout还是标准出错设备stderr的输出都需要特殊处理。

守护进程的名称通常以d结尾，比如sshd、xinetd、crond等。



# 进程 vs 线程

进程：**进程是资源分配的最小单位**。保存在硬盘上的程序运行以后，会在内存空间里形成一个独立的内存体，这个内存体**有自己独立的地址空间，有自己的堆**，上级挂靠单位是操作系统。操作系统会以进程为单位，分配系统资源（CPU时间片、内存等资源）。

线程：轻量级进程(Lightweight Process，LWP），是**操作系统调度（CPU调度）执行的最小单位**。

#### 进程和线程的区别与联系

【区别】：

- **调度**：**线程作为调度和分配的基本单位，进程作为拥有资源的基本单位**；
- **并发性**：**不仅进程之间可以并发执行，同一个进程的多个线程之间也可并发执行**；
- **拥有资源**：**进程是拥有资源的一个独立单位，线程不拥有系统资源**，但可以访问隶属于进程的资源。进程所维护的是程序所包含的资源（静态资源）， 如：**地址空间，打开的文件句柄集，文件系统状态，信号处理handler等**；线程所维护的运行相关的资源（动态资源），如：**运行栈，调度相关的控制信息，待处理的信号集等**；
- **系统开销**：在创建或撤消进程时，由于系统都要为之分配和回收资源，导致系统的开销明显大于创建或撤消线程时的开销。但是进程有独立的地址空间，一个进程崩溃后，在保护模式下不会对其它进程产生影响，而线程只是一个进程中的不同执行路径。线程有自己的堆栈和局部变量，但线程之间没有单独的地址空间，一个进程死掉就等于所有的线程死掉，所以**多进程的程序要比多线程的程序健壮，但在进程切换时，耗费资源较大，效率要差一些**。

【联系】：

- 线程在执行过程中，需要协作同步。不同进程的线程间要利用消息通信的办法实现同步。
- 线程共享整个进程的资源（寄存器、堆栈、上下文）线程中执行时一般都要进行同步和互斥，因为他们共享同一进程的所有资源；

 

|              | 进程                                                       | 线程                                       |
| ------------ | ---------------------------------------------------------- | ------------------------------------------ |
| 概念         | 资源分配的独立单元                                         | CPU调度的基本单元                          |
| 切换所需资源 | 大                                                         | 中                                         |
| 效率         | 低                                                         | 一般                                       |
| 拥有资源     | 独立的堆栈（全局变量保存在堆中，局部变量及函数保存在栈中） | 独立的栈，共享堆                           |
| 关系         | 一个进程包含多个线程，至少包含一个线程                     |                                            |
| 结束的影响   | 进程结束后所有线程将被销毁                                 | 线程结束不影响其他线程                     |
| 系统资源     | 拥有系统资源                                               | 不拥有系统资源，可以访问隶属于该进程的资源 |
|              |                                                            |                                            |



#### 一个形象的例子解释进程和线程的区别

![在这里插入图片描述](http://blog.chinaunix.net/attachment/201310/23/29270628_1382541951nJe7.jpg)

这副图是一个双向多车道的道路图，假如我们**把整条道路看成是一个“进程”的话**，那么图中由白色虚线分隔开来的**各个车道就是进程中的各个“线程”了**。

- **这些线程(车道)共享了进程(道路)的公共资源(土地资源)**。
- 这些线程(车道)必须依赖于进程(道路)，也就是说，**线程不能脱离于进程而存在(就像离开了道路，车道也就没有意义了)**。
- **这些线程(车道)之间可以并发执行(各个车道你走你的，我走我的)，也可以互相同步(某些车道在交通灯亮时禁止继续前行或转弯，必须等待其它车道的车辆通行完毕)**。
- **这些线程(车道)之间依靠代码逻辑(交通灯)来控制运行，一旦代码逻辑控制有误(死锁，多个线程同时竞争唯一资源)，那么线程将陷入混乱，无序之中**。
- **这些线程(车道)之间谁先运行是未知的**，只有在线程刚好被分配到CPU时间片(交通灯变化)的那一刻才能知道。

![在这里插入图片描述](http://5b0988e595225.cdn.sohucs.com/images/20180622/6765e36cc4604fba897976638af03524.jpeg)

#### 协程--一种比线程更加轻量级的存在

​		协程不是被操作系统内核所管理，而完全是由程序所控制（也就是在用户态执行）。这样带来的好处就是性能得到了很大的提升，不会像线程切换那样消耗资源。

  子程序，或者称为函数，在所有语言中都是层级调用，比如A调用B，B在执行过程中又调用了C，C执行完毕返回，B执行完毕返回，最后是A执行完毕。所以子程序调用是通过**栈实现的**，**一个线程就是执行一个子程序**。子程序调用总是一个入口，一次返回，调用顺序是明确的。而协程的调用和子程序不同。

  **协程在子程序内部是可中断的，然后转而执行别的子程序，在适当的时候再返回来接着执行**。



# 查看磁盘、cpu、内存、负载

#### *1. df -hl ：查看磁盘使用情况*

#### *2.top ：查看cpu 内存等使用情况*

#### *3. 内存*

- #### *1. free -m ：查询内存详情*

- #### *2. cat /proc/meminfo ：查看内存详细信息*
