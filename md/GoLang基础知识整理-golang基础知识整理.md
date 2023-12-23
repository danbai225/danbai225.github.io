---
title: GoLang基础知识整理
date: 2019-12-15 16:10:32.321
updated: 2022-03-29 21:34:52.759
url: https://p00q.cn/?p=13
categories: 
- Go
tags: 
- Go
- 整理
---

# Go语言的特点:

简洁、快速、安全
并行、有趣、开源
内存管理、数组安全、编译迅速
跨平台

# 第一个程序

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

# Go内置关键字

- break 退出本层循环
- default 流程控制关键字 在没有匹配结果下默认执行
- func 方法/函数定义关键字
- interface 接口定义关键字
- select 用于选择不同类型的通讯
- case 流程控制-条件
- defer 用于资源的释放,会在函数返回之前进行调用。     
- go 用于开启一个协程(goroutine)         
- map map类型关键字         
- struct 结构体关键字
- chan 通道关键字       
- else 流程控制-否则               
- goto 跳转       
- package 定义包关键字        
- switch 流程控制-条件选择
- const 常量定义
- fallthrough 在switch中，使用fallthrough可以强制执行后面的case代码。   
- if 流程控制-条件判断       
- range 用于读取slice、map、channel数据       
- type 用于声明自定义类型
- continue 跳过本次循环   
- for 流程控制-循环                  
- import 引入包关键字
- return 退出函数/方法            
- var 变量定义关键字

# 命名规范
在main函数的包必须为main
## 变量

- 变量命名基本上遵循相应的英文表达或简写。
- 在相对简单的环境（对象数量少、针对性强）中，可以将一些名称由完整单词简写为单个字母
- 若变量类型为 bool 类型，则名称应以 `Has`, `Is`, `Can` 或 `Allow` 开头
- 如果变量为私有，且特有名词为首个单词，则使用小写，如 apiClient
- 其它情况都应当使用该名词原有的写法，如 APIClient、repoID、UserID。
## 文件命名规范
由于文件跟包无任何关系， 而又避免windows大小写的问题，所以推荐的明明规范如下： 文件名应一律使用小写， 不同单词之间用下划线分割, 命名应尽可能地见名知意

## 函数结构体

- 使用驼峰命名
- 如果包外不需要访问请用小写开头的函数
- 如果需要暴露出去给包外访问需要使用大写开头的函数名称

# 数组和切片
固定长度的列表叫数组
长度不固定的叫切片

## 数组使用
`var 数组变量名 [元素数量]Type`
- 数组变量名：数组声明及使用时的变量名。
- 元素数量：数组的元素数量，可以是一个表达式，但最终通过编译期计算的结果必须是整型数值，元素数量不能含有到运行时才能确认大小的数值。
- Type：可以是任意基本类型，包括数组本身，类型为数组本身时，可以实现多维数组。
```
var NumberList = [3]string
NumberList = [3]string{"Python", "Golang", "Java"}
```
- 获取数组长度: len
- 获取数组容量: cap
- 比较两个数组是否相等 在数组类型长度相同的情况下可以直接用==或!=进行判断
- 遍历通过for+range
```
var team [3]string
team[0] = "hammer"
team[1] = "soldier"
team[2] = "mum"
for k, v := range team {
    fmt.Println(k, v)
}
```

## 切片使用
Go中的slice依赖于数组，它的底层就是数组，所以数组具有的优点，slice都有。且slice支持可以通过append向slice中追加元素，长度不够时会动态扩展，通过再次slice切片，可以得到得到更小的slice结构，可以迭代、遍历等。

实际上slice是这样的结构：先创建一个有特定长度和数据类型的底层数组，然后从这个底层数组中选取一部分元素，返回这些元素组成的集合(或容器)，并将slice指向集合中的第一个元素。换句话说，slice自身维护了一个指针属性，指向它底层数组中的某些元素的集合。
## 创建、初始化、访问slice
一种是使用make()：
```
// 创建一个length和capacity都等于5的slice
slice := make([]int,5)

// length=3,capacity=5的slice
slice := make([]int,3,5)
```
直接赋值初始化的方式创建slice：
```
// 创建长度和容量都为4的slice，并初始化赋值
color_slice := []string{"red","blue","black","green"}

// 创建长度和容量为100的slice，并为第100个元素赋值为3
slice := []int{99:3}
```
注意区分array和slice：
```
// 创建长度为3的int数组
array := [3]int{10, 20, 30}

// 创建长度和容量都为3的slice
slice := []int{10, 20, 30}
```
对slice切片的语法为：
```
SLICE[A:B]
SLICE[A:B:C]
```
其中A表示从SLICE的第几个元素开始切，B控制切片的长度(B-A)，C控制切片的容量(C-A)，如果没有给定C，则表示切到底层数组的最尾部。
还有几种简化形式：
```
SLICE[A:]  // 从A切到最尾部
SLICE[:B]  // 从最开头切到B(不包含B)
SLICE[:]   // 从头切到尾，等价于复制整个SLICE
```
其他函数
 - copy()函数 
 - append()函数
 
扩容

当slice的length已经等于capacity的时候，再使用append()给slice追加元素，会自动扩展底层数组的长度。
# 结构体

Go语言中提供了对struct的支持,struct,中文翻译称为结构体，与数组一样，属于复合类型，并非引用类型。

## 定义
使用struct关键字可以定义一个结构体,结构体中的成员，称为结构体的字段或属性。
```
type Member struct {
    Id     int
    Name   string
    Email  string
    Gender int
    Age    int
}
``` 
## 声明
```
var m2 = Member{1,"小明","xiaoming@163.com",1,18} // 简短变量声明方式
var m3 = Member{id:2,"name":"小红"}// 简短变量声明方式
```
## 特性
- 值传递: 结构体与数组一样，是复合类型，无论是作为实参传递给函数时，还是赋值给其他变量，都是值传递，即复一个副本。
- 没有继承: Go语言是支持面向对象编程的，但却没有继承的概念，在结构体中，可以通过组合其他结构体来构建更复杂的结构体。
- 结构体不能包含自己: 一个结构体，并没有包含自身，比如Member中的字段不能是Member类型，但却可能是*Member。

# Map
字典: 键值型

## 定义
```
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type

/* 使用 make 函数 */
map_variable := make(map[key_data_type]value_data_type)
```
## 使用
```
package main

import "fmt"

func main() {
    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

    /* map插入key - value对,各个国家对应的首都 */
    countryCapitalMap [ "France" ] = "巴黎"
    countryCapitalMap [ "Italy" ] = "罗马"
    countryCapitalMap [ "Japan" ] = "东京"
    countryCapitalMap [ "India " ] = "新德里"

    /*使用键输出地图值 */
    for country := range countryCapitalMap {
        fmt.Println(country, "首都是", countryCapitalMap [country])
    }

    /*查看元素在集合中是否存在 */
    capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(capital) */
    /*fmt.Println(ok) */
    if (ok) {
        fmt.Println("American 的首都是", capital)
    } else {
        fmt.Println("American 的首都不存在")
    }
}
```
delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。

# 协程
　我们知道，协程（coroutine）是Go语言中的轻量级线程实现，由Go运行时（runtime）管理。

　　在一个函数调用前加上go关键字，这次调用就会在一个新的goroutine中并发执行。当被调用的函数返回时，这个goroutine也自动结束。需要注意的是，如果这个函数有返回值，那么这个返回值会被丢弃。
## Golang 协程的应用
通过go关键字开启协程
```
func Add(x, y int) {
    z := x + y
    fmt.Println(z)

}
func main() {
    for i:=0; i<10; i++ {
        go Add(i, i)
    }
}
```
执行上面的代码，会发现屏幕什么也没打印出来，程序就退出了。
　　对于上面的例子，main()函数启动了10个goroutine，然后返回，这时程序就退出了，而被启动的执行 Add() 的 goroutine 没来得及执行。我们想要让 main() 函数等待所有 goroutine 退出后再返回，但如何知道 goroutine 都退出了呢？这就引出了多个goroutine之间通信的问题。
　　在工程上，有两种最常见的并发通信模型：共享内存 和 消息。
 　　下面的例子，使用了锁变量（属于一种共享内存）来同步协程，事实上 Go 语言主要使用消息机制（channel）来作为通信模型
```
package main

import (
    "fmt"
    "sync"
    "runtime"
)

var counter int = 0
func Count(lock *sync.Mutex) {
    lock.Lock() // 上锁
    counter++
    fmt.Println("counter =", counter)
    lock.Unlock()   // 解锁
}

func main() {
    lock := &sync.Mutex{}
    for i:=0; i<10; i++ {
        go Count(lock)
    }
    for {
        lock.Lock() // 上锁
        c := counter
        lock.Unlock()   // 解锁
        runtime.Gosched() // 出让时间片
        if c >= 10 {
            break
        }
    }
}
```
## channel
消息机制认为每个并发单元是自包含的、独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享。每个并发单元的输入和输出只有一种，那就是消息。

channel 是 Go 语言在语言级别提供的 goroutine 间的通信方式，我们可以使用 channel 在多个 goroutine 之间传递消息。channel是进程内的通信方式，因此通过 channel 传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。channel 是类型相关的，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。

channel的声明形式为：
```
var chanName chan ElementType
```
举个例子，声明一个传递int类型的channel：
```
var ch chan int
```
 　　使用内置函数 `make()` 定义一个channel：
```
ch := make(chan int)
```
在channel的用法中，最常见的包括写入和读出：
```
// 将一个数据value写入至channel，这会导致阻塞，直到有其他goroutine从这个channel中读取数据
ch <- value

// 从channel中读取数据，如果channel之前没有写入数据，也会导致阻塞，直到channel中被写入数据为止
value := <-ch
//默认情况下，channel的接收和发送都是阻塞的，除非另一端已准备好。

//我们还可以创建一个带缓冲的channel：

c := make(chan int, 1024)

// 从带缓冲的channel中读数据
for i:=range c {
　　...
}
```
此时，创建一个大小为1024的int类型的channel，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被填完之前都不会阻塞。

可以关闭不再使用的channel：
`close(ch)`

应该在生产者的地方关闭channel，如果在消费者的地方关闭，容易引起panic；

现在利用channel来重写上面的例子：
```
func Count(ch chan int) {
    ch <- 1
    fmt.Println("Counting")
}

func main() {
    chs := make([] chan int, 10)
    for i:=0; i<10; i++ {
        chs[i] = make(chan int)
        go Count(chs[i])
    }
    for _, ch := range(chs) {
        <-ch
    }
}
```

## 通道
```

var ch1 chan int  　　　　// 普通channel

var ch2 chan <- int 　　 // 只用于写int数据

var ch3 <-chan int 　　 // 只用于读int数据

//可以通过类型转换，将一个channel转换为单向的：

ch4 := make(chan int)

ch5 := <-chan int(ch4)   // 单向读

ch6 := chan<- int(ch4)  //单向写
```

各知识点收集于网络
