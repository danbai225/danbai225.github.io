---
title: golang err:Today is Thursday,please give me 50$ i go to eat KFC
date: "2022-07-28 12:40:01"
updated: "2022-07-28 21:52:53"
url: https://p00q.cn/?p=709
categories:
    - Go
    - 奇技淫巧
tags:
    - kfc
summary: 这篇文章描述了作者在编写代码时遇到的一个错误，然后作者发现是钱不够导致的问题。作者通过让另一个有钱的人给他50块解决了问题。代码经过修改后，再次运行就没有问题了。
id: "709"
---

今天写代码遇到一个奇怪的报错记录一下，对了今天是星期四
```golang
package main

var kfc = KFC{}

type KFC struct {
}
type people struct {
	money int
}

func (k *KFC) Buy(p *people) bool {
	temp := p.money
	temp -= 50
	if temp < 0 {
		panic("Today is Thursday,please give me 50$ i go to eat KFC")
	}
	p.money = temp
	return true
}

func main() {
	p := &people{money: 30}
	println(kfc.Buy(p))
}
```
run:
```
panic: Today is Thursday,please give me 50$ i go to eat KFC

goroutine 1 [running]:
main.(*KFC).Buy(...)
        /Users/cctc/code/Test/main.go:15
main.main()
        /Users/cctc/code/Test/main.go:23 +0x27


```
想了半天，最后发现是我钱不够。那么找一个有人的人让他给你50块就行了
```golang
package main

var kfc = KFC{}

type KFC struct {
}
type people struct {
	money int
}
func (p *people) give(p2 *people, num int) {
	p.money -= num
	p2.money += num
}
func (k *KFC) Buy(p *people) bool {
	temp := p.money
	temp -= 50
	if temp < 0 {
		panic("Today is Thursday,please give me 50$ i go to eat KFC")
	}
	p.money = temp
	return true
}

func main() {
	p := &people{money: 30}
	p2 := &people{money: 100}
	p2.give(p, 50)
	println(kfc.Buy(p))
	println(kfc.Buy(p2))
}
```
这样就没问题了。

懂我意思吧

