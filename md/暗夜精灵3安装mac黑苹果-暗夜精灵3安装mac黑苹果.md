---
title: 暗夜精灵3安装mac黑苹果
date: 2021-06-06 20:32:25.03
updated: 2021-06-06 22:03:03.046
url: https://p00q.cn/?p=322
categories: 
- 瞎折腾
tags: 
- mac
- 黑苹果
---

# 前言
买了台式后笔记本没有怎么使用比较闲置，正好有时间准备给装个黑苹果玩玩。也顺便学习学习看看在mac上开发是怎样一种体验。
# 准备
开始前先准备以下东西：
- 一个16gU盘（需要10g左右空间）
- MacOS镜像
- TransMac 或 etcher(看到过其他教程用过，我这里没用过刷树莓派到是用过)
- 分区工具DiskGenius
- EFI引导

镜像：这里提供一下我下载镜像的地址：[点击进入下载](https://mirrors.dtops.cc/ISO/MacOS/)
版本呢我选择的是10.15.7(19H2)不支持macOS Big Sur 11.x
TransMac ：[点击下载](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/TransMac_1622982963187.exe)
DiskGenius：[进入官网下载](https://www.diskgenius.cn/download.php)
EFI引导：[下载](https://danbai-cloud.oss-cn-chengdu.aliyuncs.com/uploads%2F2021%2F06%2F06%2FgxVuUgR4_EFI.zip?Expires=1622983634)
# 开始
准备工作完成后开始进入正题。

## 制作Mac启动盘
1.初始化u盘 记得备份数据
 插入U盘打开DiskGenius
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622983877347.png?x-oss-process=style/blog)
选择u盘设备右键，删除所有分区保存更改。然后新建一个分区，完成后关闭软件。![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622983995259.png?x-oss-process=style/blog)

2.写入mac镜像
 打开TransMac.exe软件
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622984122096.png?x-oss-process=style/blog)
同样会有对应的U盘设备鼠标移动到设备上右键选择![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622984208074.png?x-oss-process=style/blog)
提示框选择Yes、OK无需其余修改。
操作完成后再鼠标右键选择![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622984387791.png?x-oss-process=style/blog)
弹出提示框选择文件，选我们下号的mac镜像文件
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622984421758.png?x-oss-process=style/blog)
点击yes、OK后继续等待进度条完成大概十几分钟（取决u盘写入速度）关闭软件。![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622985579821.png?x-oss-process=style/blog)完成
3. 替换EFI
再打开DiskGenius
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622985648295.png?x-oss-process=style/blog)
找到这个位置选择全部文件右键强力删除
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622985690859.png?x-oss-process=style/blog)
然后再吧我们下载的EFI文件给放进去![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622985730598.png?x-oss-process=style/blog)
之后可以关闭软件拔出U盘开始安装系统了。

## 安装Mac系统
将启动盘插入笔记本,通电开机笔记本然后点按ESC进入启动菜单
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622987060153.png?x-oss-process=style/blog)
F9进入引导菜单
选择第一个你u盘名字的启动项回车。
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622987087382.png?x-oss-process=style/blog)
进入这个界面选择Insstall mac Catalina
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622987212726.png?x-oss-process=style/blog)
等等加载完成，进入mac安装界面。
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622987357254.png?x-oss-process=style/blog)
选择磁盘工具
然后对你原来的windows盘或者一个你需要安装mac的盘进行抹除格式改为Mac OS扩展（日志式）
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622987496861.png?x-oss-process=style/blog)
分区后再关闭磁盘工具窗口回到首界面选择安装macOS，期间大约30分钟左右。
安装完成后通过U盘启动选择界面选择MAC完成安装。
引导修复，安装完成mac后不需要再用u盘启动，需要最后用u盘进入winpe系统也就是在前面引导菜单那里出现的两个同名u盘名的启动项的第二个。
进入PE使用BOOTICE软件。找到UEFI-修改启动序列-添加 ，再挂载的EFI分区里面，找到OC文件夹下面的OpenCore.efi文件，选择并打开:
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1622988032087.png?x-oss-process=style/blog)
添加完后关机拔掉u盘能顺利启动mac就成功了。

参考文章：https://www.sqlsec.com/2018/08/clover.html

