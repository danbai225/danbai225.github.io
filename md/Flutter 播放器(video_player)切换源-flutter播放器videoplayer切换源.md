---
title: Flutter 播放器(video_player)切换源
date: 2020-04-27 15:22:40.542
updated: 2021-06-10 15:18:34.851
url: https://p00q.cn/?p=68
categories: 
- Flutter
tags: 
- Flutter
---

# 依赖
  chewie: ^0.9.10
  video_player: ^0.10.8+1

# 代码

定义变量

```
 VideoPlayerController _videoPlayerController;
  ChewieController _chewieController;
```
主要代码
```
 @override
  void dispose() {
    // 退出处理
    super.dispose();
    if (_videoPlayerController != null) {
      _videoPlayerController.dispose();
      _chewieController.dispose();
      _videoPlayerController = null;
      _chewieController = null;
    }
  }

//初始化控制器
  void _initController(String link) {
    _videoPlayerController = VideoPlayerController.network(link)
      ..initialize().then((_) {
        _chewieController = ChewieController(
          customControls: CustomControls(),
          allowedScreenSleep: false,
          videoPlayerController: _videoPlayerController,
          aspectRatio: _videoPlayerController.value.aspectRatio,
          autoPlay: true,
          looping: true,
        );
        setState(() {});
      });
  }

  //加载视频
  _loadVideo(String url) async {
    if (_videoPlayerController == null) {
      //没有视频在播放
      _initController(url);
    } else {
      // 如果有控制器，我们需要先处理旧的
      final oldController = _videoPlayerController;
      // 为下一帧的结束注册回调
      // 处理一个旧控制器
      // (调用setState后不再使用)
      WidgetsBinding.instance.addPostFrameCallback((_) async {
        await oldController.dispose();
      });
      // 通过将其设置为null来确保没有使用该控制器
      setState(() {
        _videoPlayerController = null;
        _initController(url);
      });
    }
  }


在build中...
	Container(
              child: _videoPlayerController != null &&
                      _chewieController != null &&
                      _videoPlayerController.value.initialized
                  ? Chewie(
                      controller: _chewieController,
                    )
                  : null,
            )
```