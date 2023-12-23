---
title: golang解析sql从中获取表名
date: 2022-04-24 20:34:12.673
updated: 2022-04-24 20:37:30.787
url: https://p00q.cn/?p=610
categories: 
- 学习
- Go
- 开发
tags: 
- Go
- MySQL
- sql
- antlr
---

# antlr
ANTLR是Another Tool for Language Recognition的简写，是一个用Java语言编写的识别器工具。它能够自动生成解析器，并将用户编写的ANTLR语法规则直接生成目标语言的解析器，它能够生成Java、Go、C等语言的解析器客户端。
作者还维护了很多语法规则其中就有[mysql规则](https://github.com/antlr/grammars-v4/tree/master/sql/mysql)

# 使用

下载最新的jar包[https://www.antlr.org/download/antlr-4.10.1-complete.jar](https://www.antlr.org/download/antlr-4.10.1-complete.jar)
从github中下载`MySqlLexer.g4`和`MySqlParser.g4`

执行命令生成go文件 (需要java环境)
`java -jar antlr-4.10.1-complete.jar -visitor -Dlanguage=Go -o parser MySqlLexer.g4`

`java -jar antlr-4.10.1-complete.jar -visitor -Dlanguage=Go -o parser MySqlParser.g4`

目录结构如下：
```
├── MySqlLexer.g4
├── MySqlParser.g4
├── antlr-4.10.1-complete.jar
├── antlr.go
└── parser
    ├── MySqlLexer.interp
    ├── MySqlLexer.tokens
    ├── MySqlParser.interp
    ├── MySqlParser.tokens
    ├── mysql_lexer.go
    ├── mysql_parser.go
    ├── mysqlparser_base_listener.go
    └── mysqlparser_listener.go

```
`antlr.go`:
```
package antlr

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	logs "github.com/danbai225/go-logs"
	"strings"
	"test/antlr/parser"
)

type Ml struct {
	*parser.BaseMySqlParserListener
	tableNames map[string]struct{}
}

func (m *Ml) EnterTableName(ctx *parser.TableNameContext) {
	if m.tableNames == nil {
		m.tableNames = make(map[string]struct{})
	}
	m.tableNames[ctx.GetText()] = struct{}{}
}
func (m *Ml) GetTableNames() []string {
	arr := make([]string, 0)
	if m.tableNames != nil {
		for k := range m.tableNames {
			arr = append(arr, k)
		}
	}
	return arr
}
func GetTableNames(sql string, sqlType string) []string {
	tokenStream := antlr.NewCommonTokenStream(parser.NewMySqlLexer(antlr.NewInputStream(strings.ToUpper(sql))), antlr.TokenDefaultChannel)
	sqlParser := parser.NewMySqlParser(tokenStream)
	ml := Ml{}
	switch sqlType {
	case "dml":
		antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DmlStatement())
	case "ddl":
		antlr.ParseTreeWalkerDefault.Walk(&ml, sqlParser.DdlStatement())
	}
	return ml.GetTableNames()
}

func Antlr() {
	sql := "SELECT without FROM fails"
	logs.Info(GetTableNames(sql, "dml"))
}
```
通过覆盖实现`BaseMySqlParserListener`的`EnterTableName`接口
再解析的过程中每个表名节点都将调用这个方法收集表名。