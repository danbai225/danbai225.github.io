---
title: Flutter Video_Flutter+Chewie切换播放源泄漏问题解决
date: "2020-09-24 16:09:34"
updated: "2020-10-17 00:14:35"
url: https://p00q.cn/?p=194
categories:
    - Flutter
tags:
    - Flutter
    - video_player
summary: 本文介绍了在切换播放源时可能出现的后台播放问题，并提供了一种解决方案。在切换播放源之前，先将新的`VideoPlayerController`保存在一个变量中，待其初始化完成后再销毁上一个控制器。具体操作通过判断当前控制器是否为空，如果为空则直接初始化新的控制器；如果不为空，则先保存当前控制器到一个变量，在下一次刷新时销毁之前的控制器，并初始化新的控制器。控制器初始化的过程是先初始化`VideoPlayerController`，然后使用其初始化完成的回调来初始化`ChewieController`。在页面销毁时，需要先销毁播放器以避免泄漏。
id: "194"
---

# 切换播放
如果在上`VideoPlayerController`初始化完成之前初始化一个新的`VideoPlayerController`而上一个没有完成初始化无法销毁,则可能造成后台播放的问题.

# 解决方案
在切换一个新的`VideoPlayerController`之前先将其保存放到一个变量上,待初始化完成后销毁.
切换播放源:
```
if (_videoPlayerController == null) {
     //初始化一个控制器
      _initController(url);
    } else {
      //如果当前控制器不为空则把当前控制器保留到old变量在下一次刷新时销毁
      final oldController = _videoPlayerController;
      WidgetsBinding.instance.addPostFrameCallback((_) async {
        await oldController.dispose();
      });
      setState(() {
        _videoPlayerController = null;
        _initController(url);
      });
    }
```
控制器初始化:
```
_initController(String link) {
    _videoPlayerController = VideoPlayerController.network(link)
      ..initialize().then((_) {
        _chewieController = ChewieController(
            customControls: CustomControls(),
            allowedScreenSleep: false,
            videoPlayerController: _videoPlayerController,
            aspectRatio: _videoPlayerController.value.aspectRatio,
            autoPlay: true,
            looping: true,
            startAt: startTime);
        setState(() {});
      });
  }
```

记得在页面销毁时先销毁播放否则也会造成泄漏
```
@override
  dispose() {
    super.dispose();
    if(_videoPlayerController!=null){
      _videoPlayerController.pause();
      _videoPlayerController.dispose();
      _videoPlayerController = null;
      _chewieController = null;
    }
  }
```

完整解决:https://www.thetopsites.net/article/58955831.shtml
