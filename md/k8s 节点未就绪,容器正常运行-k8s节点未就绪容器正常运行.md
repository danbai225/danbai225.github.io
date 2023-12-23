---
title: k8s 节点未就绪,容器正常运行
date: 2020-09-18 18:54:33.43
updated: 2020-09-18 18:59:42.023
url: https://p00q.cn/?p=165
categories: 
- 瞎折腾
tags: 
- k8s
---

# 异常状态
今天打开k8s控制面板,看到主节点和k8s2两个节点都未就绪.
查看容器服务有些服务处于未运行状态.
运行在容器上的非k8s容器正常运行.
![image](../res/img/165-1.jpeg)

# 重启kubelet.service
我尝试百度后,重启k8s服务
` systemctl restart kubelet.service` 解决了

![image.png](../res/img/165-1.jpeg)