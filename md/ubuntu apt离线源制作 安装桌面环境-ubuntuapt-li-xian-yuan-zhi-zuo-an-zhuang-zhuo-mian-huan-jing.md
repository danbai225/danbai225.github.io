---
title: ubuntu apt离线源制作 安装桌面环境
date: 2022-07-26 13:53:43.21
updated: 2022-07-27 12:19:09.092
url: https://p00q.cn/?p=708
categories: 
- 瞎折腾
tags: 
- apt
- ubuntu
---

# 制作
```
#创建文件夹
mkdir -p /opt/offline-packages/archives
cd /opt/offline-packages/archives
#下载依赖
sudo apt-get download $(apt-cache depends --recurse --no-recommends --no-suggests --no-conflicts --no-breaks --no-replaces --no-enhances gdm3 ubuntu-desktop gnome | grep "^\w" | sort -u)
#建立依赖
cd /opt/offline-packages
sudo dpkg-scanpackages -m . > ./archives/Packages

#打包
cd ../
tar -zcvf offline-packages.tar.gz offline-packages
```
# 使用
```
sudo tar -zxvf offline-packages.tar.gz -C /opt
# 写入本地源，如有需要，提取备份原有源
echo "deb file:///opt/offline-packages  archives/"| sudo tee /etc/apt/sources.list 
sudo apt-get update --allow-insecure-repositories
```
原文：[ubuntu apt-get离线源制作](https://blog.csdn.net/HYESC/article/details/123551830)