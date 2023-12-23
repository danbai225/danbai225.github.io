---
title: IDEA插件开发-发送事件日志(eventLog)
date: 2021-11-28 18:32:05.401
updated: 2021-11-29 19:41:59.978
url: https://p00q.cn/?p=515
categories: 
- 开发
tags: 
- IDEA
- 插件开发
---


想把文字展示到这个位置，但是没有找到直接的方法。
![image.png](../res/img/515.jpeg)
后发现用过事件日志可以实现，于是把消息写入事件日志。
代码如下:
```
fun sendNotify(project: Project, title: String?, content: String?, type: NotificationType?) {
    NotificationGroupManager.getInstance().getNotificationGroup("com.github.danbai225.pwlchat.notify")
        .createNotification(content!!, type!!)
        .setTitle(title)
        .notify(project)
}
```
需要注册一个组
在`plugin.xml`下添加
```
   <notificationGroup id="com.github.danbai225.pwlchat.notify" displayType="NONE"/>
```
`displayType`类型可以决定消息出现的方式

[代码](https://github.com/danbai225/pwl-chat/blob/main/src/main/kotlin/com/github/danbai225/pwlchat/notify/Notification.kt)