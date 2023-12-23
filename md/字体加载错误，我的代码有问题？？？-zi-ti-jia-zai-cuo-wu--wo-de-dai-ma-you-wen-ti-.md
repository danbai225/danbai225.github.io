---
title: 字体加载错误，我的代码有问题？？？
date: 2021-09-03 12:16:33.279
updated: 2021-09-03 12:20:35.179
url: https://p00q.cn/?p=388
categories: 
- 开发
tags: 
- css
- html
---

apache在使用一个字体文件时，浏览器加载出错。起初在我本机浏览器没问题，就没在意。但在其他主机上出现了字体现在不正常的问题，查明后发现原来是我本机安装了这个字体。:joy:

```
Failed to decode downloaded font: file:///Users/danbai/Downloads/website/zpixnew.ttf
index.html:1 OTS parsing error: OS/2: Failed to parse table
```
后尝试把字体转成其他格式。[`在线字体文件转换器`](https://convertio.co/zh/font-converter/)

```apache
index.html:1 Failed to decode downloaded font: file:///Users/danbai/Downloads/website/zpixb.woff2
index.html:1 OTS parsing error: OS/2: Failed to parse table
```
然而报错还是这样，最后下载了[仓库](https://github.com/SolidZORO/zpix-pixel-font)的使用 demo 发现用里面的字体文件是显示正常的后我下载了以前tag里的字体文件也没有问题。搞了半天原来是版本造成的，搞了半天。:tada:

> 另外顺便了解了一下字体格式兼容性的问题 https://segmentfault.com/a/1190000018215174
>
> @Font-face目前浏览器的兼容性：
>
> •Webkit/Safari(3.2+)：TrueType/OpenType TT (.ttf) 、OpenType PS (.otf)；
> •Opera (10+)： TrueType/OpenType TT (.ttf) 、 OpenType PS (.otf) 、 SVG (.svg)；
> •Internet Explorer： 自ie4开始，支持EOT格式的字体文件；ie9支持WOFF；
> •Firefox(3.5+)： TrueType/OpenType TT (.ttf)、 OpenType PS (.otf)、 WOFF (since Firefox 3.6)
> •Google Chrome：TrueType/OpenType TT (.ttf)、OpenType PS (.otf)、WOFF since version 6
>
> **由上面可以得出：.eot + .ttf /.otf + svg + woff = 所有浏览器的完美支持。**