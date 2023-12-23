---
title: golang 申请一块内存
date: 2022-07-30 14:43:34.401
updated: 2022-07-30 14:45:11.316
url: https://p00q.cn/?p=710
categories: 
- Go
- 奇技淫巧
- 开发
tags: 
- Go
---

申请一块占用实际内存的内存空间。
```golang
func RequestMemory(size int64) []byte {
	bytes := make([]byte, size)
	for i := range bytes {
		if i == 0 {
			bytes[i] = 0
		} else {
			bytes[i] = bytes[i-1] + 1
		}
	}
	return bytes
}

```