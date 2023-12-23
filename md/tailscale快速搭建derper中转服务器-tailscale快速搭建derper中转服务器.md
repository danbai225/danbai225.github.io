---
title: tailscale快速搭建derper中转服务器
date: 2022-08-01 12:58:25.673
updated: 2022-08-01 13:42:18.741
url: https://p00q.cn/?p=738
categories: 
- 瞎折腾
tags: 
- derper
- tailscale
---

# 使用docke-compose快速部署
```
version: '3.3'
services:
  yangchuansheng:
    restart: always
    network_mode: host
    container_name: derper
    environment:
      - DERP_HTTP_PORT=2180
      - 'DERP_ADDR=:21443'
    image: ghcr.io/yangchuansheng/ip_derper
```
`docker-compose up -d`
开放端口3478 还有配置中的端口
端口根据自身情况进行修改
# 修改管理配置
![image-1659330026532](../res/img/738-1.jpeg)```

```
"derpMap": {
		"Regions": {"900": {
			"RegionID":   900,
			"RegionCode": "my",
			"Nodes": [{
				"Name":             "1",
				"RegionID":         900,
				"HostName":         "1.14.75.115",
				"IPv4":             "1.14.75.115",
				"DERPPort":         21443,
				"InsecureForTests": true,
			}],
		}},
	},
```
保存重启客户端
![image-1659330163364](../res/img/738-2.jpeg)

教程来源：https://icloudnative.io/posts/custom-derp-servers/