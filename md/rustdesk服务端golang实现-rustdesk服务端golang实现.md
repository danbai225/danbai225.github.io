---
title: rustdesk服务端golang实现
date: "2022-05-28 21:34:31"
updated: "2022-06-28 17:21:52"
url: https://p00q.cn/?p=644
categories:
    - Go
    - 开发
tags:
    - rust
    - golang
    - rustdesk
    - 远程桌面
    - 开源
summary: 作者在这篇文章中分享了自己寻找远程桌面软件的经历，并介绍了在Mac平台下使用开源软件rustdesk建立自己的远程桌面服务的过程。作者通过实现基础功能、局域网直连、中继器连接和加密连接等功能，最终完成了一个可以使用的远程桌面服务端。文章最后，作者表示从rustdesk项目中学到了很多，并提到了rust的语法难度。
id: "644"
---

# 自建远程桌面服务

前段时间我在寻找一款远程桌面软件，原来我一直使用的是 windows 自带的 rdp 在有公网的情况下速度很快。但是后面我换到 mac 平台下无法继续使用 windows 的 rdp ，我先后尝试了 teamviewer 、todesk 。

后面在github找到一款开源的远程软件 [rustdesk](https://github.com/rustdesk/rustdesk) 。它是跨平台的，且支持服务端自建（一段时间里是不支持的）。于是我根据客户端开源的proto协议尝试用golang实现一个服务端[go-rustdesk-server](https://github.com/danbai225/go-rustdesk-server)。
在仓库建立一段时间后rustdesk也正式开源了他们的服务端 [rustdesk-server](https://github.com/rustdesk/rustdesk-server) 我去对比了我的实现，并根据官方实现去优化和修改我前面的一些错误理解。

# 已实现功能

- 基础功能的实现
- 局域网（有公网的）直连
- 通过中继器连接
- 加密的连接

仓库地址：https://github.com/danbai225/go-rustdesk-server

# 后记

从 rustdesk 项目中学到了很多，rust的语法也是真的难。
