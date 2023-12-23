---
title: 在golang中使用clash作为http客户端的代理
date: "2023-07-28 14:23:57"
updated: "2023-07-28 14:23:57"
url: https://p00q.cn/?p=906
categories:
    - Go
    - 奇技淫巧
    - 开发
    - 编码协议
    - 解决办法
tags:
    - Go
    - proxy
    - http
summary: |-
    这是一篇使用Clash作为HTTP客户端代理的Go语言教程。在爬取网站时，经常会遇到一些反爬虫的网站，这些网站会检测你的IP地址，如果被检测到是爬虫，就会拒绝访问。为了解决这个问题，可以使用代理来隐藏真实IP地址。代理的原理是将请求发送到代理服务器，代理服务器再发送给目标服务器并返回数据。这样目标服务器无法知道真实IP地址。本文介绍了使用Clash作为HTTP客户端代理的方法。

    首先，需要获取HTTP代理。可以通过免费获取或者付费购买的方式来获取。作者选择了使用机场提供的v2或SSR协议的代理。

    接下来，将Clash代码引入Go程序中，并设置代理的配置信息。代码中使用了Clash的Start()方法来启动代理，并通过调用Proxies()方法获取代理列表。

    然后，使用代理URL并创建http.Transport和http.Client对象。在创建http.Transport对象时，将代理URL传入Proxy字段。最后，通过调用http.Client对象的Get方法来发送请求。

    以上就是使用Clash作为HTTP客户端代理的具体实现方法。通过这种方式，可以在Go程序中轻松地使用Clash作为HTTP客户端代理来访问需要隐藏真实IP地址的网站。
id: "906"
---

# 在golang中使用clash作为http客户端的代理

在爬取网站时，经常会遇到一些反爬虫的网站，这些网站会检测你的ip地址，如果你的ip地址被检测到是爬虫的话，就会拒绝你的访问，这时候就需要使用代理来解决这个问题。代理的原理就是你的请求先发送给代理服务器，代理服务器再发送给目标服务器，目标服务器返回数据给代理服务器，代理服务器再返回给你。这样的话，目标服务器就无法知道你的真实ip地址，从而达到隐藏你的真实ip地址的目的。网上有一些免费的获取http代理的网站或者可以采用付费购买的方式获取http代理。

我这里我想采用机场来做我的http代理但是机场一般提供的是v2或者ssr协议的代理，恰好clash是用golang写的所以我能在我的golang程序中直接引入clash的代码开启一个本地端口进行代理。

下面是我的代码：

代理实现
```go
package proxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Dreamacro/clash/config"
	C "github.com/Dreamacro/clash/constant"
	"github.com/Dreamacro/clash/hub"
	logs "github.com/danbai225/go-logs"
	"go.uber.org/automaxprocs/maxprocs"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func init() {
	rand.NewSource(time.Now().UnixNano())
}

type Clash struct {
	configPath string
}

func New(configPath string) *Clash {
	return &Clash{configPath: configPath}
}
func (c *Clash) Start() {
	maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))
	C.SetHomeDir(".")
	if c.configPath != "" {
		if !filepath.IsAbs(c.configPath) {
			currentDir, _ := os.Getwd()
			c.configPath = filepath.Join(currentDir, c.configPath)
		}
		C.SetConfig(c.configPath)
	} else {
		configFile := filepath.Join(C.Path.HomeDir(), C.Path.Config())
		C.SetConfig(configFile)
	}
	if err := config.Init(C.Path.HomeDir()); err != nil {
		logs.Err("Initial configuration directory error: %s", err.Error())
		return
	}
	if err := hub.Parse(hub.WithExternalController(":9091")); err != nil {
		logs.Err("Parse config error: %s", err.Error())
		return
	}
	time.Sleep(time.Second * 1)
	c.Speed()
	c.RandomSelect()
}
func (c *Clash) Speed() {
	proxies := c.Proxies()
	group := sync.WaitGroup{}
	for _, s := range proxies {
		group.Add(1)
		go func() {
			http.Get(fmt.Sprintf(`http://127.0.0.1:9091/proxies/%s/delay?timeout=5000&url=%s`, url.PathEscape(s.Name), url.QueryEscape(`https://baidu.com`)))
			group.Done()
		}()
	}
	group.Wait()
}
func (c *Clash) Proxies() []P {
	resp, _ := http.Get(`http://127.0.0.1:9091/proxies`)
	all, _ := io.ReadAll(resp.Body)
	m := make(map[string]interface{})
	m2 := make(map[string]P)
	json.Unmarshal(all, &m)
	marshal, _ := json.Marshal(m["proxies"])
	json.Unmarshal(marshal, &m2)
	ps := make([]P, 0)
	for k := range m2 {
		ps = append(ps, m2[k])
	}
	return ps
}
func (c *Clash) Switchover(name string) {
	json := fmt.Sprintf(`{"name":"%s"}`, name)
	req, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:9091/proxies/GLOBAL", bytes.NewBufferString(json))
	if err != nil {
		logs.Err(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		logs.Err(err)
		return
	}
}
func (c *Clash) EffectiveProxy() []P {
	ps := make([]P, 0)
	for _, p := range c.Proxies() {
		if len(p.History) > 0 && p.History[len(p.History)-1].Delay > 0 {
			ps = append(ps, p)
		}
	}
	return ps
}
func (c *Clash) RandomSelect() string {
	ps := c.EffectiveProxy()
	name := ps[rand.Int63n(int64(len(ps)))].Name
	c.Switchover(name)
	logs.Info("切换为", name)
	return name
}

type P struct {
	Alive   bool `json:"alive"`
	History []struct {
		Time      time.Time `json:"time"`
		Delay     int       `json:"delay"`
		MeanDelay int       `json:"meanDelay"`
	} `json:"history"`
	Name string `json:"name"`
	Type string `json:"type"`
	Udp  bool   `json:"udp"`
}

```

使用demo
```golang
package main

import (
	logs "github.com/danbai225/go-logs"
	"io"
	"net/http"
	"net/url"
	"test/proxy"
)

func main() {
	clash := proxy.New("clash.yml")
	clash.Start()
	// 创建代理URL
	proxyURL, err := url.Parse("http://127.0.0.1:7799")
	if err != nil {
		return
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	c := &http.Client{
		Transport: transport,
	}
	resp, err := c.Get("https://ip.useragentinfo.com/json")
	all, err := io.ReadAll(resp.Body)
	logs.Info(string(all))
}

```
