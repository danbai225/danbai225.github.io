---
title: 使用vue+element开发移动端页面
date: "2019-10-23 15:19:04"
updated: "2022-03-29 21:33:45"
url: https://p00q.cn/?p=5
categories:
    - 学习
tags:
    - Vue
summary: 该文介绍了作者用vue开发一套手机前端页面的经历。作者在介绍了项目的创建和vue项目开发的基本结构后，分享了自己在开发过程中遇到的困难和解决方法。其中困难包括对如何引入和使用组件的不熟悉、使用axios获取数据时的跨域和cookie问题、以及修改element-ui默认样式的困难。作者最终通过修改样式元素的class来覆盖原有样式。
id: "5"
---

![](https://img.hacpai.com/bing/20180824.jpg?imageView2/1/w/960/h/540/interlace/1/q/100)

# 介绍
为了更好的使[淡白影视](https://github.com/danbai225/dbys)在手机web页面有更好的表现。所以决定用vue开发一套手机前端页面.~~编译的源码可以通过[cordova](https://cordova.apache.org/)编译成跨平台的应用.~~
ps:经过开发一段时间后,发现在各大浏览器都有自己内置的播放器所以导致很多弹幕显示问题.并且使用cordova打包的app应用也存在播放问题.
所以app端采用原生开发,移动web端对弹幕不支持.

项目地址:[淡白影视app](https://github.com/danbai225/dbysapp)

# 创建项目
项目基于nodejs开发所以需要安装nodejs
[下载地址](https://nodejs.org/zh-cn/download/)
然后安装vue-cli、vue和cordova加element-ui等所需依赖
推荐使用vue-cli的vue ui创建vue项目
在cmd窗口输入vue ui即可,稍后会在浏览器打开管理页面
![TIM截图20191023200955.png](https://img.hacpai.com/file/2019/10/TIM截图20191023200955-a2caf185.png)
进入到工作目录创建项目按照需要配置创建项目即可

# vue项目开发
## 项目结构
![TIM截图20191023201836.png](https://img.hacpai.com/file/2019/10/TIM截图20191023201836-ba40a69f.png)
这是已经开发了两天的项目结构了
已经实现了一些基本功能
下面讲讲这两天遇到的困难
1. 由于vue是我最近才学的还不熟悉,所以在开发组件和使用组件的时候对如何引入组件和使用组件是按照vue-cli生成的初始项目修改添加才搞懂的.
2.使用axios获取数据也遇到了一定困难,跨域问题还有cookie问题
3.使用element-ui修改默认样式也让我搞了半天,由于我主要注重移动端页面所以element-ui一些组件偏大,需要修改css,根据官网提供的工具编译不通过.所以我用的方法是找到需要修改样式元素的class通过`!important`覆盖它原有的样式


