# 容器化 (in progress)

容器化是直把服务部署在容器之中，它与传统在服务器上直接部署应用的方式不同。

容器可以帮助我们可以快速发布业务软件系统。它将业务软件系统封装在一个最小化的独立操作系统中（minimum system)，就好像只有这个软件在一台独立的主机上运行一样。

## 什么是容器

容器是一个业务软件系统打包在一起的单元结构。容器可以在任何操作系统中发布，可以快速、可靠从旧环境迁移到另一个操作系统环境。比如Docker容器（引擎），在其之上运行镜像，便可以在镜像之中运行所有可执行程序。

容器单元彼此之间互相隔离，互不干扰。它不受运行和开发环境的影响。

### 为什么要什么容器？

接下来可能有朋友会问，为什么不用VMware、VirtualBox或者其它已经非常成熟虚拟化技术。我们引用[Docker](https://www.docker.com/resources/what-container)上的结构图说明：

容器             |  虚拟机
:-------------------------:|:-------------------------:
![](../docs/part8_docker-containerized-appliction-blue-border_1.png)  |  ![](../docs/part8_container-vm-whatcontainer_2.png)

虚拟机采用的是硬件级别的**系统级**虚拟化方案，需要CPU等硬件支持，而容器则是在虚拟机之上的更高级别的**进程级**虚拟化技术。

可见，容器依赖虚拟机但是比虚拟机虚拟出的系统更轻量，因为它并非为了派生操作系统，而是为了精简系统、复用虚拟机并在其之上运行软件。

## 相关资料

[Docker内部原理][Docker内部原理]
[为什么使用Docker][为什么使用Docker]

[Docker内部原理]: https://medium.com/@nagarwal/understanding-the-docker-internals-7ccb052ce9fe
[为什么使用Docker]: https://runnable.com/docker/why-use-docker