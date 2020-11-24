### 开篇词-这样学redis,才能技高一筹

#### 前言

问题一:为了保证数据的可靠性,redis需要在磁盘上读写AOF和RDB,但是在高并发场景下会带来两个新的问题:1、写AOF和RDB会造成redis性能抖动,2、redis集群数据同步和实例恢复时,读RDB较慢,限制速度

解决方案:使用非易失内存NVM,可保证高速读写+快速持久话数据



#### redis使用场景

缓存、数据库、分布式锁



#### 常见的坑

- CPU-数据结构复杂度、跨CPU核的访问
- 内存-主从同步和AOF内存竞争
- 存储持久化-SSD做快照的性能抖动
- 网络通信-多实例的异常网络丢包



大部分技术人有一个误区:只关注零散的知识点,缺乏系统观.但是系统观才是至关重要的,拥有了系统观意味着你能有依据、有章法的定位、解决问题

#### Redis知识全景图

两大维度(系统维度+应用维度),三大主线(高性能、高可靠、高可扩展)

![B89BF969-5E23-4A0E-A983-50E113CFA29D](/var/folders/rn/j1054c3j5q55dcxvlcdwjs_00000gn/T/com.yinxiang.Mac/com.yinxiang.Mac/WebKitDnD.ZQVCbB/B89BF969-5E23-4A0E-A983-50E113CFA29D.png)

高性能主线:线程模型、数据结构、持久化、网络框架

高可靠主线主从复置、哨兵机制

高可扩展主线:数据分配、负载均衡



应用维度按照两种方式学习:“应用场景驱动”+“典型案例驱动” == 面的梳理+点的掌握



#### redis的问题画像图

![A6CC124F-133C-461D-8CAE-46265CDCF40D](/var/folders/rn/j1054c3j5q55dcxvlcdwjs_00000gn/T/com.yinxiang.Mac/com.yinxiang.Mac/WebKitDnD.Mz48bH/A6CC124F-133C-461D-8CAE-46265CDCF40D.png)

eg:如果你遇到redis响应变慢的问题,参照上图可看出是与性能主线相关,而性能主线又与数据结构、异步机制、RDB、AOF重写相关.



附上一张课程表的目录

![836603FC-4C5C-4B62-9050-D5B6086F03C7](/var/folders/rn/j1054c3j5q55dcxvlcdwjs_00000gn/T/com.yinxiang.Mac/com.yinxiang.Mac/WebKitDnD.0efkSU/836603FC-4C5C-4B62-9050-D5B6086F03C7.png)