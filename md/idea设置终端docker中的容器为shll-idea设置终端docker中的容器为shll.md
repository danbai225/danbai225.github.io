---
title: idea设置终端docker中的容器为shll
date: 2022-03-16 17:25:13.341
updated: 2022-03-16 17:33:18.988
url: https://p00q.cn/?p=549
categories: 
- 奇技淫巧
tags: 
- docker
---

# shell
把下面代码保存为脚本，在idea终端设置中设置shell路径为脚本
![image.png](../res/img/549.jpeg)
# win
```
@echo off
docker exec -it ubuntu bash
pause
```
# uinx
```
#!/bin/bash
docker exec -it ubuntu bash
```
# 代码
- tools->deployment->brower remote host 代码同步
- JetBrains Gateway 新的远程开发功能在远程装了个idea
- docker映射开发工作目录