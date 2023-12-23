---
title: golang beego框架orm操作mysql数据库
date: "2019-12-12 16:08:54"
updated: "2022-03-29 21:34:16"
url: https://p00q.cn/?p=11
categories:
    - Go
tags:
    - Beego
    - Go
summary: 本文介绍了在GoLang中使用beego框架进行ORM（对象关系映射）操作的基本知识。首先，文章介绍了如何连接数据库，然后说明了如何创建模型。在查询方面，文章展示了使用原始SQL查询和根据主键查询的例子。在插入和更新方面，文章展示了插入新数据和更新已有数据的例子。最后，文章总结了学习beego框架的一些感受，并提供了相关中文文档和参考链接。
id: "11"
---

上周学完了GoLang基础部分知识,这周学习下beego框架.
这篇文章记录了ORM部分知识

# 连接数据库

在包名dao下创建mysql.go

```
package dao

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "user:pass@tcp(185.207.153.189:3306)/tempyun")
	//参数3按照官网参数连接不上mysql
	// orm.RegisterDataBase("default", "mysql", "root:root@/orm_test?charset=utf8")
}

func Getcon() orm.Ormer {
	return orm.NewOrm()
}
```

# model

```
type User struct {
	Username string `orm:"pk"`
	Password string
	Email string
	Headurl string
}
```

如果表的主键不是id，那么需要加上pk注释，显式的说这个字段是主键
属性名首字母大写
第二个单词首字母不能大写,如要大写将对应 如HeadUrl  对应表字段 Head_url

[更多使用](https://beego.me/docs/mvc/model/models.md)

# 查询
```
func Ut()  {
	var o=dao.Getcon()
	var users []models.User
	//原始sql
	num, err := o.Raw("SELECT * FROM user").QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num)
		fmt.Println(users[0])
	}
	//根据主键
	user := models.User{Username: "danbai"}
	err = o.Read(&user)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Password)
	}
}
```
# 插入更新

```
func Ut1()  {
	var o=dao.Getcon()
	var user = models.User{"test","123","225@qq.com","https://www.123.com"}
	username, err := o.Insert(&user)
	if err == nil {
		fmt.Println(username)
	}
	user1 := models.User{Username: "test"}
	if o.Read(&user1) == nil {
		user1.Password = "pass"
		if num, err := o.Update(&user1); err == nil {
			fmt.Println(num)
		}
	}
}
```
# 总结
今天大概看了遍文档 和spring框架还是有点区别,但是内容都差不多路由、控制、服务、视图、过滤器等.
简单练习了下,发现启动速度也很快.占用内存比java少了不少.
接下来就开始边做边学了,加油!

[beego中文文档](https://beego.me/docs/intro/)

