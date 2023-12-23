---
title: 在mysql中使用json数据类型 grom
date: 2022-03-16 15:01:22.599
updated: 2022-03-16 15:18:29.041
url: https://p00q.cn/?p=548
categories: 
- Go
- 数据库
tags: 
- Go
- MySQL
- gorm
---

# mysql json类型

MySQL 5.7起支持JSON数据类型的字段。JSON作为现在最为流行的数据交互形式，MySQL也不断跟进，在5.7版本开始新增JSON数据类型。本文基于`MySQL 8.0.28`

# gorm配置
目前gorm还不支持json类型需要自己通过[自定义类型](https://gorm.io/docs/data_types.html#Implements-Customized-Data-Type)来实现
```go

type Test struct {
	ID   int64 
	Name string
	Age  int64
	Data Data `gorm:"type:json"`
}
type Data struct {
	UserInfo UserInfo `json:"user_info"`
	DataSize int64    `json:"data_size"`
	Comment  string   `json:"comment"`
}

func (c Data) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Data) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type UserInfo struct {
	Site        string `json:"site"`
	PhoneNumber string `json:"phone_number"`
	HeadAddress string `json:"head_address"`
	Distance    int64  `json:"distance"`
}

```
# 查询

通过`字段名->'$.json属性名'`来访问json属性
```
var tests []modle.Test
	begin.Model(&modle.Test{}).Where("data->'$.data_size' < ?", 100).Find(&tests)
	log.Println(len(tests))
```
查询json中data_size小于100的数据

# 修改
由于没有支持json类型只能靠自己去实现Build接口去实现一些mysql内置函数的拼接[接口实现例子](https://github.com/go-gorm/datatypes/blob/master/json.go#L133)
```
	var test modle.Test
	begin.Model(&modle.Test{}).Where("data->'$.user_info.site' = ?", "khcUTxnAUBNZkAOBXmJnYwHPYtdmKfLRdHwnxfozlYhANGTEXKIbphMyzFoUDJFf").Find(&test)
	log.Println(test.ID)

	test.Data.DataSize++
	begin.Updates(&test)
```
sql
```
UPDATE `test` SET `name`='BqLpGYaP',`age`=23,`data`='{"user_info":{"site":"khcUTxnAUBNZkAOBXmJnYwHPYtdmKfLRdHwnxfozlYhANGTEXKIbphMyzFoUDJFf","phone_number":"padWxlWJrOQ","head_address":"YAZxSORyWOdoXdjvnuixBJU
PwAvzVXZsuZJPCOhuHzBDMbBGXUvuJPhaLEIYTWrmhRnhzYABtzNDaOSOXwRziKoEUBdZOIOhzZNoncfjFodkYglnqudLUevHrOlAoPki","distance":835},"data_size":849,"comment":"ufDErrgfktIVDCbnFinqfUQbWEAMiEEawhlPjWAIOOXYyjqzObjcvFrTBzEzChZlaKpBEpTfzDaJjgicZzdiNWMADdwIFvHzcbylnTKVjwWNUBVpfjdkUlysPtjdZxAk"}' WHERE `id` = 311091
```
通过函数更新
```sql

update test set `data` = JSON_SET(`data`, '$.user_info.site', "a") where id =311091;
```

# 内置函数
![924df2314068a62aad5d439c96e77585.jpg](../res/img/548-1.jpg)

# 更多
[聊聊MySQL 8.0中的Json增强](https://www.51cto.com/article/665187.html)
[MySQL · 最佳实践 · 如何索引JSON字段](http://mysql.taobao.org/monthly/2017/12/09/)