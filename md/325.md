---
title: 用go实现一个通过模拟输入剪切板内容的程序
date: "2021-06-08 20:32:38"
updated: "2021-06-10 15:07:06"
url: https://p00q.cn/?p=325
categories:
    - Go
    - 奇技淫巧
tags:
    - 工具
    - Go
summary: 该程序的目的是为了解决在禁用CV大法的程序中输入复杂或长字符串时的不便。它通过监听Ctrl + I的快捷键，将剪切板中的内容作为字符串读取，然后用按键库模拟输入字符串的每个字符。该程序使用了robotgo、clipboard和systray等库，并且集成了一个带有图标的托盘应用程序。在运行该程序时，会在托盘中显示一个图标，右键点击可以选择退出程序。然后通过按下Ctrl + I快捷键，即可模拟输入剪切板中的字符串。但需要注意的是，并不是所有程序都支持输入，需要进一步使用虚拟键盘才能实现对这些程序的支持。
id: "325"
---

# 简介
在有些程序中为了安全起见，禁用了CV大法。在输入很复杂很长窜的字符串时感觉很难受。于是便有了下面这个程序，该代码的核心部分是通过robotgo这个库封装的C代码实现的，所以没有什么硬代码。
# 代码
```Go

package main

import (
	_ "embed" //Go1.16支持将资源文件一同打包
	"github.com/atotto/clipboard" //读取剪切板
	"github.com/getlantern/systray" //托盘图标
	"github.com/go-vgo/robotgo" //按键库
	"os"
	"strings"
	"test/log"
	"time"
)
//下面这个注释就是添加本地静态资源的
//go:embed assets/jqb.ico
var iconData []byte
func main() {
	 systray.Run(onReady, onExit)
	for {
		if robotgo.AddEvents("i", "ctrl"){
			// 读取剪切板中的内容到字符串
			content, err := clipboard.ReadAll()
			if err != nil {
				log.LogErr(err)
			}
			split := strings.Split(content, "")
			for _, s := range split {
				time.Sleep(time.Millisecond*66)
				robotgo.KeyTap(s)
			}
		}
	}
}
func onReady() {
	systray.SetIcon(iconData)
	systray.SetTitle("帮你输入剪切板")
	systray.SetTooltip("帮你输入剪切板")
	mQuit := systray.AddMenuItem("退出", "点击退出应用程序")
	go func() {
		select {
		case <-mQuit.ClickedCh:
			os.Exit(0)
		}
	}()
}

func onExit() {

}
```
# 编译运行
通过命令`go build -ldflags "-s -w -H=windowsgui"`编译出来的程序没有命令行窗口。运行后在托盘会展示图标，右键可以选择退出。
运行效果![image.png](1)
快捷键`Ctrl`+`I`就能模拟输入剪切板内的字符串了。
注：并不是所有程序都支持输入，WeGame测试不行。
部分程序不支持，进一步支持需要用到虚拟键盘
https://github.com/ddxoft/master

