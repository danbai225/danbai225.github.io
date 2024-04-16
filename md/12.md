---
title: golang项目练习-临时云
date: "2019-12-14 16:09:54"
updated: "2022-03-29 21:35:10"
url: https://p00q.cn/?p=12
categories:
    - Go
tags:
    - Beego
    - Go
summary: 这篇文章介绍了一个临时网盘项目，用于测试和临时存放json文件。用户可以通过主域名和路径访问文件，也可以通过子域名的方式访问文件。项目还可以进一步完善功能，如权限管理、统计、协同等。文章还提到了注意事项，即在前端页面传入的路径参数和用户名中不得出现 ".."。
id: "12"
---

前几天看了个帖子有人在问有没有 临时网上放json的网盘并且能直连访问的 用于测试. 加之这两天学了golang于是就写了这么个小项目
# 项目  
![logo.png](https://img.hacpai.com/file/2019/12/logo-c4e1c8ed.png)

[项目地址](https://github.com/danbai225/tempyun)  
还有很多功能可以完善,感兴趣的朋友可以去扩展  
例如 权限、统计、协同、等等(学完go可以用来练练手项目很小很容易修改)  
  
[线上地址](http://tempyun.com)
# 用途

前端页面测试,json测试,临时网盘

# 子域名访问实现
域名:tempyun.com
文件:a.txt
通过主域名访问路径如下:
http://tempyun.com/service/files/{用户名}/a.txt

实现访问路径:
http://{用户名}.tempyun.com/a.txt

实现:
将域名泛解析到服务器ip *.tempyun.com   192.168.1.1

通过过滤器将子域名(用户名)提取

服务代理访问路径并返回
```
func Ymfilter() {
	var Filter = func(ctx *context.Context) {
		doamin := beego.AppConfig.String("doamin")
		if ctx.Input.Host() != doamin {
			req := httplib.Get(strings.Replace(ctx.Input.Site(), ctx.Input.Host(), doamin, -1) + ":" + strconv.Itoa(ctx.Input.Port()) + "/" + strings.Replace(ctx.Input.Host(), "."+doamin, "", -1) + ctx.Input.URL())
			bytes, _ := req.Bytes()
			ctx.ResponseWriter.Write(bytes)
		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Filter)
}
```
# 注意
记得检测前端页面传来的路径参数和用户名不得出现 ".."


