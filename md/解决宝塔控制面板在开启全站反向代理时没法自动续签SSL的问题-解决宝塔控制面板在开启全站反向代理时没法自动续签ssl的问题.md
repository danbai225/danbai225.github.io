---
title: 解决宝塔控制面板在开启全站反向代理时没法自动续签SSL的问题
date: 2022-03-12 12:01:14.998
updated: 2022-03-12 12:10:00.34
url: https://p00q.cn/?p=546
categories: 
- 解决办法
tags: 
- nginx
- 宝塔面板
- 自动续签
- ssl证书
---

# 免费ssl续签

宝塔提供的一个自动向`Let's Encrypt`申请的免费ssl证书，有效期3个月。
所以在到期前1个月会自动重新申请，但如果你的网站开启了反向代理那么在验证时没法访问到放在你网站根目录下的验证文件。

# 解决办法

很简单，在你的位置设置里找到Nginx配置文件选项

![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1647057652981.png?x-oss-process=style/blog)

在引用反向代理配置前一行添加
![image.png](https://danbai.oss-cn-chengdu.aliyuncs.com/bk/image_1647057738457.png?x-oss-process=style/blog)
其中请将`/www/wwwlogs/p00q.cn`替换成你的网站根目录
```
    location /.well-known{
        root /www/wwwlogs/p00q.cn;
    }
```
# 验证

访问一下你网站`https://p00q.cn/.well-known/`路径出现nginx403就好了
再尝试续签能正常完成即可。