---
title: golang数组切片拷贝问题
date: "2020-11-08 19:31:37"
updated: "2020-11-08 19:42:31"
url: https://p00q.cn/?p=258
categories:
    - Go
tags:
    - Go
summary: 本文介绍了在Go语言中对数组切片进行传递时的一些注意事项。在Go语言中，数组的传递是采用引用传递的方式，即传递的是数组的内存地址。因此，即使重新赋值，也只是复制了一份内存地址。如果要解决这个问题，可以使用copy()函数来进行复制。此外，在遇到多维数组时，使用copy()函数无法解决子元素数组的复制问题，需要进行二次复制来解决。此外，本文还介绍了一个通用的深拷贝函数Clone()，可以用来对任意结构进行深拷贝。
id: "258"
---

# 数组切片传递
数组在传递中采用引用传递,即使重新赋值也是只是复制他的一份内存地址
例如:
```
func main() {
	array1:=[]string{"00","01"}
	array2:=array1
	array1[0]="asd"
	fmt.Println(array1)
	fmt.Println(array2)
}
```
输出:
```
[asd 01]
[asd 01]
```
## 使用COPY来解决这个问题
注意初始化时的长度要和拷贝数组长度>=
```
func main() {
	array1:=[]string{"00","01"}
	array2:=make([]string, len(array1))
	copy(array2,array1)
	array1[0]="asd"
	fmt.Println(array1)
	fmt.Println(array2)
}
```
输出:
```
[asd 01]
[00 01]
```
# 当遇到多维数组的时候
在多维数组的情况下COPY貌似也不行了,在拷贝数组的时候并没有拷贝子元素数组
```
func main() {
	array1:=[][]string{
		{"00","01"},
		{"10","11"},
	}
	array2:=make([][]string, len(array1), len(array1[0]))
	copy(array2,array1)
	array1[0][0]="asd"
	fmt.Println(array1)
	fmt.Println(array2)
}
```
输出:
```
[[asd 01] [10 11]]
[[asd 01] [10 11]]

```
## 在改变子元素数组的时候再copy一次
```
func main() {
	array1:=[][]string{
		{"00","01"},
		{"10","11"},
	}
	array2:=make([][]string, len(array1), len(array1[0]))
	copy(array2,array1)
	as := make([]string, 2)
	copy(as,array2[0])
	as[0]="asd"
	array1[0]=as
	fmt.Println(array1)
	fmt.Println(array2)

}
```
输出:
```
[[asd 01] [10 11]]
[[00 01] [10 11]]
```

# golang深拷贝任意结构代码
[来源](https://www.cnblogs.com/ayanmw/p/8666378.html)
```
func Clone(a, b interface{}) error {
	buff := new(bytes.Buffer)
	enc := gob.NewEncoder(buff)
	dec := gob.NewDecoder(buff)
	if err := enc.Encode(a); err != nil {
		return err
	}
	if err := dec.Decode(b); err != nil {
		return err
	}
	return nil
}
```
再试试
```
func main() {
	array1:=[][]string{
		{"00","01"},
		{"10","11"},
	}
	array2:=make([][]string, len(array1), len(array1[0]))
	Clone(&array1,&array2)
	array1[0][0]="asd"
	fmt.Println(array1)
	fmt.Println(array2)
}
```
输出:
```
[[asd 01] [10 11]]
[[00 01] [10 11]]
```

最近刚遇到了多维数组切片拷贝的问题所有记录一下....
