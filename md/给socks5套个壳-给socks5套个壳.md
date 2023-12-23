---
title: 给socks5套个壳
date: 2021-09-06 20:05:18.677
updated: 2021-09-14 15:43:20.03
url: https://p00q.cn/?p=389
categories: 
- Go
tags: 
- Go
- socks5
---

# 限制我学习

由于公司限制了哔哩哔哩等学习网站我用的代理也是明文的socks5于是上级路由还是能够通过http给我拦截了。 注:我用socks5主要的原因是方便访问局域网

于是我便给我的socks5套了一层壳
通过修改[socks5proxy](https://github.com/shikanon/socks5proxy)开源库即可实现。

# 解除限制

前面的开源库本身就是一个socks5服务器和客户端，但为了连接我之前的socks（有代理）需要用到另一个socks5的开源实现[socks5](https://github.com/txthinking/socks5)
修改 `socks5proxy` 代码

`server.go`文件
```Go
//添加依赖
"github.com/txthinking/socks5"

//修改41行附近

// 连接真正的远程服务
	newClient, err := socks5.NewClient("socksServer", "", "", 0, 60)
	if err!=nil{
		return
	}
	dstServer, err := newClient.Dial("tcp", request.RAWADDR.String())
	//dstServer, err := net.DialTCP("tcp", nil, request.RAWADDR)
	if err != nil {
		log.Print(client.RemoteAddr(), err)
		return
	}
	defer dstServer.Close()
```
然后编译服务端和客户端代码，在socks5原来的服务器端上运行这个新的套壳服务端。
在本地电脑上运行客户端。注意启动参数，启动客户端后把网络代理的socks5替换成本地客户端的代理端口，就能访问b站了。

# 追更 9.14
前面的实现有些问题，后面我重写了。
[Github](https://github.com/danbai225/tcpproxy)
