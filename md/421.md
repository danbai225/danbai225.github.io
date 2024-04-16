---
title: 手动升级py后apt update无法执行，显示没有apt_pkg模块
date: "2021-09-16 11:09:38"
updated: "2021-09-16 11:12:38"
url: https://p00q.cn/?p=421
categories:
    - 报错处理
tags:
    - Python
    - apt
summary: |-
    这个错误是由于没有找到名为'apt_pkg'的模块导致的。解决方法是创建一个软链接。首先进入/usr/lib/python3/dist-packages目录，然后以管理员身份运行以下命令：
    ```
    sudo ln -s apt_pkg.cpython-36m-x86_64-linux-gnu.so apt_pkg.so
    ```
    其中的'36m'需要根据你安装的Python版本进行替换，例如'38m'或'39m'。
id: "421"
---

ModuleNotFoundError: No module named ‘apt_pkg’
# 解决方法
```shell
# 创建一个软链接
cd /usr/lib/python3/dist-packages
sudo ln -s apt_pkg.cpython-36m-x86_64-linux-gnu.so apt_pkg.so
```
其中 `36m`需要替换你升级得py版本 例如38m、39m。
