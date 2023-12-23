---
title: 通过字符串生成手机号
date: 2023-04-18 12:12:42.466
updated: 2023-04-29 09:40:19.621
url: https://p00q.cn/?p=841
categories: 
- Go
- 奇技淫巧
- 解决办法
tags: 
- golang
---

用户信息中没有手机号，通过用户名的哈希来生成一个手机号。
```
package main

import (
	"crypto/md5"
	"fmt"
	"math/big"
)

func main() {
	// 假设用户名为"john"
	username := "john"

	// 生成一个128位的MD5哈希值
	hash := md5.Sum([]byte(username))

	// 将哈希值转换为一个大整数
	num := new(big.Int)
	num.SetBytes(hash[:])

	// 取前9位数字作为手机号码
	mobile := num.Mod(num, big.NewInt(1e9)).String()

	// 在手机号码前面添加"13"-"19"中的一个数字
	mobile = "1" + string("3456789"[num.Mod(num, big.NewInt(7)).Uint64()]) + mobile

	fmt.Println("Username:", username)
	fmt.Println("Mobile:", mobile)
}

```