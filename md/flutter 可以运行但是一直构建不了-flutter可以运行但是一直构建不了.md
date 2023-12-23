---
title: flutter 可以运行但是一直构建不了
date: 2020-07-10 12:42:34.221
updated: 2021-06-10 15:22:06.414
url: https://p00q.cn/?p=109
categories: 
- Flutter
tags: 
- Flutter
---

# 记录一个 Flutter build 问题

&emsp;&emsp;昨天flutter项目里面的一个 SDK 升级了我去更新了一下版本。
然后在虚拟机里启动了一下查看更新效果。

&emsp;&emsp;之后去 CMD ```futter build apk```执行。结果先开始 CPU 利用率达到了80 - 100之后就没怎么占用了我以为在下载依赖但发现网速并没有上去。

就像这样：![NXM6S4VE1C_5NBORSML8.png](../res/img/594354499462.jpeg)

等了十几分钟都没动。

# 解决之路

- 我尝试清除构建缓存```flutter clean``` ❌
- 直接删除项目下 ```build``` 文件夹 ❌
- 删除 ```C:\Users\DanBai\AppData\Local\Pub\Cache``` ❌
- 将代码回滚至上个版本 ❌
- 将项目文件夹整个删除重新在 Github 重新拉取下来 ✔

&emsp;&emsp;之后在项目中我又升级了依赖版本好在这次构建并没有问题一两分钟完事。在这期间都是能在模拟器运行的。虽然没能找到问题的根本，但问题是解决了。