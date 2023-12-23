---
title: golang cacheMap
date: 2020-12-12 20:30:04.321
updated: 2020-12-12 20:30:04.321
url: https://p00q.cn/?p=291
categories: 
- Go
tags: 
- Go
---

# 缓存
为了缓解一些接口的压力(数据还重复计算和数据量大),又没有redis就用Map写了个简单的做缓存.

```
package utils

import (
	"sync"
	"time"
)
var Cache cacheMap
func init()  {
	go Cache.iniCache()
}
type cacheMap struct {
	cacheMap  map[string]cacheVal
	cacheLock sync.Mutex
}
type cacheVal struct {
	val     interface{}
	outTime time.Time
}
func (this *cacheMap) iniCache() {
	this.cacheMap = make(map[string]cacheVal)
	for {
		time.Sleep(time.Second)
		this.cacheLock.Lock()
		for k, val := range this.cacheMap {
			if val.outTime.Sub(time.Now()).Microseconds() < 0 {
				delete(this.cacheMap, k)
			}
		}
		this.cacheLock.Unlock()
	}
}
func (this *cacheMap) GetCache(key string) interface{} {
	this.cacheLock.Lock()
	defer this.cacheLock.Unlock()
	return this.cacheMap[key].val
}
func (this *cacheMap) Remove(key string){
	this.cacheLock.Lock()
	defer this.cacheLock.Unlock()
	delete(this.cacheMap,key)
}

func (this *cacheMap) SetCache(key string, val interface{}, Time time.Duration) {
	this.cacheLock.Lock()
	defer this.cacheLock.Unlock()
	cacheVal := cacheVal{
		val:     val,
		outTime: time.Now().Add(Time),
	}
	this.cacheMap[key] = cacheVal
}
```
# 使用

```
func main() {
	utils.Cache.SetCache("test","cache",time.Second*10)
	println(utils.Cache.GetCache("test").(string))
}
```