---
title: Flutter缓存M3U8
date: "2020-04-26 13:03:18"
updated: "2021-06-10 15:18:02"
url: https://p00q.cn/?p=67
categories:
    - Flutter
tags:
    - Flutter
summary: 本文介绍了如何使用Flutter实现缓存影视功能，在安卓端添加了M3U8Manger的依赖库，并通过建立Flutter与安卓的通信实现了下载功能。在Flutter代码中调用安卓的M3U8Download方法进行下载任务，并通过UDP发送回调信息给Flutter端进行展示。具体的代码实现可以参考文章内容，其中包括了Flutter代码和安卓代码的链接。
id: "67"
---

# flutter实现缓存影视功能

安卓添加依赖

```
compile 'com.jwkj:M3U8Manger:v2.3.0'
```
## 主要实现内容
1.建立flutter与安卓直接的通信
2.flutter添加下载任务 调用安卓 M3U8Download
3.安卓处理任务并发回任务回调信息(在启动应用时初始化创建一个udp服务,用于发送回调信息)
4.flutter这边进入下载页面就会创建一个监听udp的链接 根据信息数据进行下载信息的展示.

[Flutter代码](https://github.com/danbai225/dbys_flutterhttps://github.com/danbai225/dbys_flutter/tree/master/lib/Page/Download)
[安卓代码](https://github.com/danbai225/dbys_flutter/tree/master/android/app/src/main/kotlin/cn/p00q/dbys)
## 代码

Flutter:
```
import 'dart:convert';

import 'package:flutter/services.dart';
import 'package:path_provider/path_provider.dart';

class DownloadManagement {
  static Map<String, String> nameMap = {};
  static List urlList = [];
  static int xz = 0;

  //与原生交互的通道
  static const platform = const MethodChannel('cn.p00q.dbys/M3U8Download');

  static init() async {
    platform.invokeMethod('Path', (await getExternalStorageDirectory()).path);
  }

  static add(String url, String pm, String jiName) {
    platform.invokeMethod(
        'Add', jsonEncode({"url": url, "pm": pm, "jiName": jiName}));
  }

  static cancel() {
    platform.invokeMethod('Cancel');
  }

  //关掉udp
  static removeBind() {
    platform.invokeMethod('removeBind');
  }
}

```
安卓:
```
package cn.p00q.dbys;

import android.os.Environment;
import android.util.Log;

import com.hdl.m3u8.M3U8DownloadTask;
import com.hdl.m3u8.M3U8Manger;
import com.hdl.m3u8.bean.M3U8;
import com.hdl.m3u8.bean.M3U8Listener;
import com.hdl.m3u8.bean.OnDownloadListener;
import com.hdl.m3u8.utils.NetSpeedUtils;

import org.eclipse.jetty.util.ajax.JSON;
import org.json.JSONException;
import org.json.JSONObject;

import java.io.File;
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.InetAddress;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class M3U8Download {
    private static String SaveFilePath=Environment.getExternalStorageDirectory().getPath() + File.separator;
    public static List<Download> downloadList=new ArrayList<>();
    private static  MyM3U8DownloadTask downloadTask;
    //1.创建服务器端DatagramSocket，指定端口
    public static DatagramSocket  socket;
    public static void init(){
        try {
            if(socket!=null){
                socket.close();
            }
            socket = new DatagramSocket(2252);
        }catch (Exception e) {
                e.printStackTrace();
            }
    }
    public static void sendData(String msg){
        byte[] buf = msg.getBytes();
        try {
            //目标主机地址，这里发送到本机所以是127.0.0.1
            InetAddress address = InetAddress.getByName("127.0.0.1");
            int port = 2256;  //目标主机的端口号
            //创建发送方的数据报信息
            DatagramPacket dataGramPacket = new DatagramPacket(buf, buf.length, address, port);
            socket.send(dataGramPacket);  //通过套接字发送数据
    } catch (UnknownHostException e) {
            e.printStackTrace();
        } catch (SocketException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
    public static void setPath(String path){
        SaveFilePath=path;
    }
    public static void Add(Object data){
        JSONObject jObject= null;
        try {
            jObject = new JSONObject(data.toString());
            String id=jObject.getString("pm")+"-"+jObject.getString("jiName");
            Download d=new Download(id,jObject.getString("pm"),jObject.getString("url"),jObject.getString("jiName"));
            downloadList.add(d);
            if(downloadList.size()>0&&(downloadTask==null||!downloadTask.isRunning())){
                Download(downloadList.get(0));
            }
        } catch (JSONException e) {
            e.printStackTrace();
        }
    }
    public static void cancel(){
        if(downloadTask!=null&&downloadTask.isRunning()){
            downloadTask.stop();
            downloadList.remove(0);
            if(downloadList.size()>0){
                Download(downloadList.get(0));
            }
        }
    }
        public static void Download(Download d){
            downloadTask = new MyM3U8DownloadTask(d.getId());
            MyM3U8DownloadTask.curTs=0;
            downloadTask.setThreadCount(10);
            downloadTask.setSaveFilePath(SaveFilePath+"/下载/"+d.getPm()+"/"+d.getJiName()+".mp4");
            downloadTask.download(d.getUrl(), new OnDownloadListener() {
                private long lastLength = 0;
                @Override
                public void onDownloading(final long itemFileSize, final int totalTs, final int curTs) {
                    Log.d("TEST",totalTs+"--curTs:"+curTs);
                    Map rs=new HashMap();
                    rs.put("type","onDownloading");
                    rs.put("taskId",d.getId());
                    rs.put("schedule",((double)curTs/totalTs));
                    rs.put("pm",d.getPm());
                    rs.put("JiName",d.getJiName());
                    Runnable networkTask = () -> M3U8Download.sendData(JSON.toString(rs));
                    new Thread(networkTask).start();
                }
                /**
                 * 下载成功
                 */
                @Override
                public void onSuccess() {
                    downloadList.remove(0);
                    if(downloadList.size()>0&&!downloadTask.isRunning()){
                        Download(downloadList.get(0));
                    }
                    Log.d("M3U8Download","下载完成了:"+d.getId());
                    Map rs=new HashMap();
                    rs.put("type","onSuccess");
                    Runnable networkTask = () -> M3U8Download.sendData(JSON.toString(rs));
                    new Thread(networkTask).start();
                }
                /**
                 * 当前的进度回调
                 *
                 * @param curLenght
                 */
                @Override
                public void onProgress(final long curLenght) {
                    if (curLenght - lastLength > 0) {
                        final String speed = NetSpeedUtils.getInstance().displayFileSize(curLenght - lastLength) + "/s";
                        lastLength = curLenght;
                        Map rs=new HashMap();
                        rs.put("type","onProgress");
                        rs.put("taskId",d.getId());
                        rs.put("speed",speed);
                        Runnable networkTask = () -> M3U8Download.sendData(JSON.toString(rs));
                        new Thread(networkTask).start();
                    }
                }
                @Override
                public void onStart() {
                    Log.d("M3U8Download","开始下载了");
                }

                @Override
                public void onError(Throwable errorMsg) {
                    Log.e("M3U8Download",errorMsg.getMessage());
                }
            });
    }
}

```

