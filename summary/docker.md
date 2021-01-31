### docker

# [docker面试题和解答(一)](https://www.cnblogs.com/xiaoyangjia/p/11388806.html)



1. 什么Docker
   Docker是一个容器化平台，它以容器的形式将您的应用程序及其所有依赖项打包在一起，以确保您的应用程序在任何环境中无缝运行。

2. Docker与虚拟机有何不同
   Docker不是虚拟化方法。它依赖于实际实现基于容器的虚拟化或操作系统级虚拟化的其他工具。

   Docker主要专注于在应用程序容器内自动部署应用程序。应用程序容器旨在打包和运行单个服务，而系统容器则设计为运行多个进程，如虚拟机。因此，Docker被视为容器化系统上的容器管理或应用程序部署工具。

   -  容器不需要引导操作系统内核，因此可以在不到一秒的时间内创建容器。此功能使基于容器的虚拟化比其他虚拟化方法更加独特和可取。
   -  由于基于容器的虚拟化为主机增加了很少或没有开销，因此基于容器的虚拟化具有接近本机的性能。
   -  对于基于容器的虚拟化，与其他虚拟化不同，不需要其他软件。
   -  主机上的所有容器共享主机的调度程序，从而节省了额外资源的需求。
   -  与虚拟机映像相比，容器状态（Docker或LXC映像）的大小很小，因此容器映像很容易分发。
   -  容器中的资源管理是通过cgroup实现的。Cgroups不允许容器消耗比分配给它们更多的资源。虽然主机的所有资源都在虚拟机中可见，但无法使用。这可以通过在容器和主机上同时运行top或htop来实现。所有环境的输出看起来都很相似。

3. 什么是Docker镜像
   Docker镜像是Docker容器的**源代码**，Docker镜像用于创建容器。使用build命令创建镜像。

4. 什么是Docker容器
   Docker容器包括**应用程序及其所有依赖项**，作为操作系统的**独立进程**运行。

5. Docker容器有几种状态
   四种状态：运行、已暂停、重新启动、已退出。

6. Dockerfile中最常见的指令是什么
   FROM：指定基础镜像
   LABEL：功能是为镜像指定标签
   RUN：运行指定的命令
   CMD：容器启动时要运行的命令

7. docker常用命令
   docker pull 拉取或者更新指定镜像
   docker push 将镜像推送至远程仓库
   docker rm 删除容器
   docker rmi 删除镜像
   docker images 列出所有镜像
   docker ps 列出所有容器

8. 容器与主机之间的数据拷贝命令
   docker cp 命令用于容器与主机之间的数据拷贝
   主机到容器：
   docker cp /www 96f7f14e99ab:/www/
   容器到主机：
   docker cp 96f7f14e99ab:/www /tmp/

9. 启动nginx容器（随机端口映射），并挂载本地文件目录到容器html的命令
   docker run -d -P --name nginx2 -v /home/nginx:/usr/share/nginx/html nginx

10. 解释一下dockerfile的ONBUILD指令
    当镜像用作另一个镜像构建的基础时，ONBUILD指令向镜像添加将在稍后执行的触发指令。如果要构建将用作构建其他镜像的基础的镜像（例如，可以使用特定于用户的配置自定义的应用程序构建环境或守护程序），这将非常有用。

11. 什么是Docker Swarm
    Docker Swarm是Docker的本机群集。它将Docker主机池转变为单个虚拟Docker主机。Docker Swarm提供标准的Docker API，任何已经与Docker守护进程通信的工具都可以使用Swarm透明地扩展到多个主机。

12. 如何在生产中监控Docker
    Docker提供docker stats和docker事件等工具来监控生产中的Docker。我们可以使用这些命令获取重要统计数据的报告。
    Docker统计数据：当我们使用容器ID调用docker stats时，我们获得容器的CPU，内存使用情况等。它类似于Linux中的top命令。
    Docker事件：Docker事件是一个命令，用于查看Docker守护程序中正在进行的活动流。
    一些常见的Docker事件是：attach，commit，die，detach，rename，destroy等。我们还可以使用各种选项来限制或过滤我们感兴趣的事件