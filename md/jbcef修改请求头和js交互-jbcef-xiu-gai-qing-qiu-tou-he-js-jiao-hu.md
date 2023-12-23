---
title: jbcef修改请求头和js交互
date: 2021-11-29 19:59:09.249
updated: 2021-11-29 19:59:09.249
url: https://p00q.cn/?p=516
categories: 
- Java
- 开发
tags: 
- 插件开发
- jcef
---


[CefBrowser文档](http://www.xuanyimao.com/doc/jcef/202011/org/cef/browser/CefBrowser.html) [帮助文档](http://www.xuanyimao.com/jcef/index.html)

# 修改请求头
## 实现CefContextMenuHandler
```
package com.github.danbai225.pwlchat.handler

import com.github.danbai225.pwlchat.component.WebChat
import org.cef.browser.CefBrowser
import org.cef.browser.CefFrame
import org.cef.callback.CefContextMenuParams
import org.cef.callback.CefMenuModel
import org.cef.handler.CefContextMenuHandler


class PwlContextMenuHandler() : CefContextMenuHandler {
    private var wb :WebChat?=null

    constructor(w:WebChat) : this() {
       wb=w
    }
    override fun onBeforeContextMenu(p0: CefBrowser?, p1: CefFrame?, p2: CefContextMenuParams?, p3: CefMenuModel?) {
        p3?.clear();//清除默认的菜单项
        p3?.addItem(CefMenuModel.MenuId.MENU_ID_USER_LAST, "Open DevTools")
        p3?.addItem(RefreshEmptyID, "RefreshEmpty")
    }

    override fun onContextMenuCommand(
        p0: CefBrowser?,
        p1: CefFrame?,
        p2: CefContextMenuParams?,
        p3: Int,
        p4: Int
    ): Boolean {
        when(p3){
            CefMenuModel.MenuId.MENU_ID_USER_LAST->{
                wb?.openDevtools()
                return true
            }
            RefreshEmptyID->{
                wb?.loadChatPage()
                return true
            }
        }
      return false
    }

    override fun onContextMenuDismissed(p0: CefBrowser?, p1: CefFrame?) {
    }
    companion object {
        private const val RefreshEmptyID=29101
    }
}
```
我是继承了JBCefBrowser类来实现功能的。
在初始化时通过`addRequestHandler`添加刚才实现的处理方法
```
        jbCefClient.addRequestHandler(PwlCefResourceRequestHandler(), cefBrowser)
```

# 和js交互
在加载之后加载js代码
```
        val openRedPacket=JBCefJSQuery.create(this)
        openRedPacket.addHandler { args: String ->
            clientApi?.openPacket(args)
            JBCefJSQuery.Response("ok")
        }
        jbCefClient.addLoadHandler(object : CefLoadHandlerAdapter() {
            override fun onLoadEnd(browser: CefBrowser, frame: CefFrame, httpStatusCode: Int) {
                clientApi?.more(1)
                browser.executeJavaScript(
                    "window.openRedPacket = function(arg) {${openRedPacket.inject(
                        "arg",
                        "response => console.log('调用成功', response)",
                        "(error_code, error_message) => console.log('调用失败', error_code, error_message)"
                    )}};",
                    "https://exp.com/js/openRedPacket.js", 0
                )
            }
        }, cefBrowser)
```
把js方法注册到`window.openRedPacket`调用这个方法就会调用到上面添加的`Handler`在回掉中通过`JBCefJSQuery.Response("ok")`返回调用结果。
`executeJavaScript`中的第二个参数是虚拟一个加载url
[代码](https://github.com/danbai225/pwl-chat)