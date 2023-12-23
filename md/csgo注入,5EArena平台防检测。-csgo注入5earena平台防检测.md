---
title: csgo注入,5EArena平台防检测。
date: 2022-10-05 21:01:38.098
updated: 2023-03-08 18:49:46.635
url: https://p00q.cn/?p=804
categories: 
- Go
- 瞎折腾
- 奇技淫巧
- 开发
- 解决办法
tags: 
- csgo
---

5EArena的防注入主要通过‘cheano.dll’来实现，只要在这之前注入`Osiris.dll`就可以转了。

另外在通过火绒自定义规则来阻止敏感位置的扫描。

实现代码:[下载](http://cloud.p00q.cn/s/bOSA)
