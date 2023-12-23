---
title: 通过拳头lucAPI来拉人进游戏
date: "2021-06-21 18:08:13"
updated: "2021-06-21 18:20:09"
url: https://p00q.cn/?p=327
categories:
    - Go
    - 奇技淫巧
    - 游戏娱乐
tags:
    - Go
    - luc
    - lol
summary: 该文档介绍了拳头官方提供给开发者的LOLApi（LUCApi）。该API可以用来开发各种有趣的工具，包括wegame的天赋配置。文章中提供了相关的参考资料和简单使用方法，以及使用Go代码实现一个LUCA连接器的示例。最后，文章展示了如何使用API调用邀请召唤师的功能。
id: "327"
---

# LUCApi
该Api是由拳头官方提供给开发者的api，所以没有封号风险。
相关文档：
[apiSwagger文档](https://www.mingweisamuel.com/lcu-schema/tool/)
[参考资料](https://github.com/XHXIAIEIN/LeagueCustomLobby#%E5%8F%82%E8%80%83%E8%B5%84%E6%96%99)
[参考资料2](https://riot-api-libraries.readthedocs.io/en/latest/lcu.html)
[参考资料3](https://github.com/molenzwiebel/Mimic)
# 简单使用
在运行lol客户端后`英雄联盟\LeagueClient`目录下会出现`lockfile`该文件记录着进程名、PID、端口号、Token、协议等信息（通过:分割）。
例如:
```
LeagueClient:13808:1667:p4Sw2_-_uUcKIeHhPmvoMQ:https
```
通过提供的api可以实现许多有意思的工具包括wegame的天赋配置也是调用api实现的。
通过go代码简单实现一个luc连接器
```Go
package riotgames

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"test/utils"
)

type LucConnector struct {
	Name string
	Token string
	ExePath string
	PID int
	Protocol string
	Port int
	Url string
}

func (luc *LucConnector)Init(){
	if luc.ExePath==""{
		//通过 github.com/shirou/gopsutil/process 开源库 获取进程
		PS := utils.GetProcessByName("LeagueClient.exe")
		for _, process := range PS {
			exe, err := process.Exe()
			if err==nil{
				//目录下是否存在lockfile
				if utils.Exists(strings.Replace(exe,"LeagueClient.exe","lockfile",1)){
					luc.ExePath=exe
					break
				}
			}
		}
	}
	//加载配置
	luc.loadLockFile()
}
func (luc *LucConnector)loadLockFile()  {
	if luc.ExePath!=""{
		data, err := ioutil.ReadFile(strings.Replace(luc.ExePath,"LeagueClient.exe","lockfile",1))
		if err==nil{
			//data数据如下
			//LeagueClient:13600:9907:bnTiqCnmxh0x7fEmqmc2DQ:https
			split := strings.Split(string(data), ":")
			if len(split)>=5{
				luc.PID=utils.String2Int(split[1])
				luc.Port=utils.String2Int(split[2])
				luc.Token=split[3]
				luc.Protocol=split[4]
				luc.Url=fmt.Sprintf("%s://riot:%s@%s%d", luc.Protocol, luc.Token,"127.0.0.1:",luc.Port)
			}
		}
	}
}
//请求构造方法
func (luc *LucConnector)Request(method string,data string,path string)[]byte  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	req, _ := http.NewRequest(method, luc.Url+path, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{Transport: tr}
	do, err := client.Do(req)
	if err==nil{
		all, _ := ioutil.ReadAll(do.Body)
		return all
	}
	return nil
}
```
# 测试api调用邀请召唤师
```
package riotgames

import (
	"fmt"
	"github.com/tidwall/gjson"
	"test/log"
	"testing"
)

func TestLoL(t *testing.T) {
	connector := LucConnector{}
	connector.Init()
	//通过召唤师名字查询 SummonerId
	request := connector.Request("GET", ``, "/lol-summoner/v1/summoners?name=fer热")
	//组装邀请召唤师数组
	data:=fmt.Sprintf(`[{
		"invitationId": "string",
		"state": "Requested",
		"timestamp": "string",
		"toSummonerId": %d,
		"toSummonerName": "string"
	}]`,gjson.GetBytes(request,"summonerId").Int())
	//发送邀请
	log.LogInfo(string(connector.Request("POST",data,`/lol-lobby/v2/lobby/invitations`)))
}
```

