---
title: 分享一个开源库，使用golang创建一些弹窗接收用户输入
date: "2022-07-19 20:32:32"
updated: "2022-07-19 20:35:49"
url: https://p00q.cn/?p=706
categories:
    - Go
    - 奇技淫巧
    - 开发
    - 推荐分享
tags: []
summary: 本文介绍了一个适用于 Golang、Windows 和 macOS 的 Zenity 对话框库，可以用于创建提示框、选择框、文件选择、进度条和输入框等功能。通过安装 zenity 包，可以在 Golang 程序中使用这些对话框功能。代码示例展示了如何使用 zenity 创建不同类型的对话框，并获取用户的交互结果。适用于需要在 Golang 程序中使用对话框的开发者。
id: "706"
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
