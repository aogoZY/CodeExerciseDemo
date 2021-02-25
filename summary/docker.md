### docker

## 什么是Docker

Docker是一个容器化平台，它以容器的形式将您的应用程序及其所有依赖项打包在一起，以确保您的应用程序在任何环境中无缝运行。

Docker专注于在应用程序容器内自动部署应用程序。应用程序容器旨在打包和运行单个服务，而系统容器则设计为运行多个进程，如虚拟机。因此，Docker被视为容器化系统上的容器管理或应用程序部署工具。

- 容器不需要引导操作系统内核，因此可以在不到一秒的时间内创建容器。此功能使基于容器的虚拟化比其他虚拟化方法更加独特和可取。

- 基于容器的虚拟化为主机减少了开销，因此基于容器的虚拟化具有接近本机的性能。

- 对于基于容器的虚拟化，与其他虚拟化不同，不需要其他软件。

- 主机上的所有容器共享主机的调度程序，从而节省了额外资源的需求。

- 与虚拟机映像相比，容器状态（Docker或LXC映像）的大小很小，因此容器映像很容易分发。

  实现什么是Docker镜像
  Docker镜像是Docker容器的**源代码**，Docker镜像用于创建容器。使用build命令创建镜像。



### Docker的应用场景

- Web 应用的自动化打包和发布。
- 自动化测试和持续集成、发布。
- 在服务型环境中部署和调整数据库或其他的后台应用。
- 从头编译或者扩展现有的OpenShift或Cloud Foundry平台来搭建自己的PaaS环境。

### Docker 的优点

- **1、简化程序：**
  Docker 让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 机器上，便可以实现虚拟化。方便快捷已经是 Docker的最大优势。
- **2、避免选择恐惧症：**
  如果你有选择恐惧症，还是资深患者。Docker打包你的纠结！比如 Docker 镜像；Docker 镜像中包含了运行环境和配置，所以 Docker 可以简化部署多种应用实例工作。比如 Web 应用、后台应用、数据库应用、大数据应用比如 Hadoop 集群、消息队列等等都可以打包成一个镜像部署。
- **3、节省开支：**
  一方面，云计算时代到来，使开发者不必为了追求效果而配置高额的硬件，Docker 改变了高性能必然高价格的思维定势。Docker 与云的结合，让云空间得到更充分的利用。不仅解决了硬件管理的问题，也改变了虚拟化的方式。

## Docker容器有几种状态

四种状态：运行、已暂停、重新启动、已退出。

## docker常用命令

1. 容器与主机之间的数据拷贝命令
   docker cp 命令用于容器与主机之间的数据拷贝
   主机到容器：
   docker cp /www 96f7f14e99ab:/www/
   容器到主机：
   docker cp 96f7f14e99ab:/www /tmp/

2. 启动nginx容器（随机端口映射），并挂载本地文件目录到容器html的命令
   docker run -d -P --name nginx2 -v /home/nginx:/usr/share/nginx/html nginx

   

   

   ## docker相关命令

   ### 容器生命周期管理

   **docker run ：**创建一个新的容器并运行一个命令

   **docker start** :启动一个或多个已经被停止的容器

   **docker stop** :停止一个运行中的容器

   **docker restart** :重启容器

   **docker kill** :杀掉一个运行中的容器

   **docker rm ：**删除一个或多少容器

   **docker pause** :暂停容器中所有的进程

   **docker unpause** :恢复容器中所有的进程

   **docker create ：**创建一个新的容器但不启动它

   **docker exec ：**在运行的容器中执行命令

   ### 容器操作

   **docker ps :** 列出容器

   **docker inspect :** 获取容器/镜像的元数据

   **docker top :**查看容器中运行的进程信息，支持 ps 命令参数

   **docker attach :**连接到正在运行中的容器

   **docker events :** 从服务器获取实时事件

   **docker logs :** 获取容器的日志

   **docker wait :** 阻塞运行直到容器停止，然后打印出它的退出代码

   **docker export :**将文件系统作为一个tar归档文件导出到STDOUT

   **docker port :**列出指定的容器的端口映射，或者查找将PRIVATE_PORT NAT到面向公众的端口

   ### 容器rootfs命令

   **docker commit :**从容器创建一个新的镜像（提交镜像）

   **docker cp :**用于容器与主机之间的数据拷贝

   **docker diff :** 检查容器里文件结构的更改

   ### 镜像仓库

   **docker login :** 登陆到一个Docker镜像仓库，如果未指定镜像仓库地址，默认为官方仓库 Docker Hub

   **docker logout :** 登出一个Docker镜像仓库，如果未指定镜像仓库地址，默认为官方仓库 Docker Hub

   **docker pull :** 从镜像仓库中拉取或者更新指定镜像(下载镜像)

   **docker push :** 将本地的镜像上传到镜像仓库,要先登陆到镜像仓库

   **docker search :** 从Docker Hub查找镜像

   ### 本地镜像管理

   **docker images :** 列出本地镜像

   **docker rmi :** 删除本地一个或多少镜像

   **docker tag :** 标记本地镜像，将其归入某一仓库

   **docker build** 命令用于使用 Dockerfile 创建镜像

   **docker history :** 查看指定镜像的创建历史

   **docker save :** 将指定镜像保存成 tar 归档文件

   **docker import :** 从归档文件中创建镜像

   ### info|version

   docker info : 显示 Docker 系统信息，包括镜像和容器数。。

   docker version :显示 Docker 版本信息

   