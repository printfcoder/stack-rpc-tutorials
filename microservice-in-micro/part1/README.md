# 第一章 用户服务

本章节我们实现用户服务，用户服务分为两层，web层（user-web）与服务层（user-service），前者提供http接口，后者向web提供RPC服务。

- user-web 以下简称web
- user-service 以下简称service

web服务主要向用户提供如下接口

- 登录与token颁发
- 鉴权

我们不提供注册接口，一来增加不必要的代码量，我们的核心还是介绍如何使用Micro组件。

## 开始编写

我们先从下往上编写，也就是从服务层**user-service**开始

### user-service

### user-web