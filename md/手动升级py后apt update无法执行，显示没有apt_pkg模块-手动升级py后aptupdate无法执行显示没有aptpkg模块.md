---
title: 手动升级py后apt update无法执行，显示没有apt_pkg模块
date: 2021-09-16 11:09:38.457
updated: 2021-09-16 11:12:38.585
url: https://p00q.cn/?p=421
categories: 
- 报错处理
tags: 
- Python 
- apt
---

ModuleNotFoundError: No module named ‘apt_pkg’
# 解决方法
```shell
# 创建一个软链接
cd /usr/lib/python3/dist-packages
sudo ln -s apt_pkg.cpython-36m-x86_64-linux-gnu.so apt_pkg.so
```
其中 `36m`需要替换你升级得py版本 例如38m、39m。