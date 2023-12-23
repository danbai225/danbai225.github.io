---
title: 分享一个开源库，使用golang创建一些弹窗接收用户输入
date: 2022-07-19 20:32:32.201
updated: 2022-07-19 20:35:49.768
url: https://p00q.cn/?p=706
categories: 
- Go
- 奇技淫巧
- 开发
- 推荐分享
tags: 
---

适用于 Golang、Windows 和 macOS 的 Zenity 对话框

[zenity](https://github.com/ncruces/zenity)

安装:
`go get github.com/ncruces/zenity@latest`
```go
package main

import (
	logs "github.com/danbai225/go-logs"
	"github.com/ncruces/zenity"
	"time"
)

func main() {
	zenity.Info("你好") //提示框

	items, err := zenity.ListItems("请选择", "321", "123") //选择框
	if err == nil {
		logs.Info(items)
	}

	file, err := zenity.SelectFile() //文件选择
	if err == nil {
		logs.Info(file)
	}

	progress, err := zenity.Progress() //进度条
	progress.Text("下载中")
	for i := 0; i <= 100; i++ {
		_ = progress.Value(i)
		time.Sleep(time.Millisecond * 100)
		if progress.MaxValue() == i {
			_ = progress.Close()
		}
	}

	entry, err := zenity.Entry("请输入用户名") //输入框
	if err == nil {
		logs.Info(entry)
	}
}

```