---
title: Flutter Video_Flutter+Chewie切换播放源泄漏问题解决
date: 2020-09-24 16:09:34.147
updated: 2020-10-17 00:14:35.947
url: https://p00q.cn/?p=194
categories: 
- Flutter
tags: 
- Flutter
- video_player
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