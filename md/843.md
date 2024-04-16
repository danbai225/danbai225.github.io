---
title: http代理服务器的代理
date: "2023-04-29 15:07:29"
updated: "2023-04-29 15:20:24"
url: https://p00q.cn/?p=843
categories:
    - Go
    - 奇技淫巧
    - 开发
tags:
    - Go
summary: |-
    该代码是一个基于HTTP的代理服务器，主要实现了自动密码输入的功能。通过传入本地监听地址、远程地址、用户名和密码，可以实现在不支持密码输入的HTTP客户端上使用密码进行访问。

    具体实现步骤如下：
    1. 构造待编码的用户名和密码字符串，使用base64进行编码。
    2. 监听本地地址，等待连接。
    3. 当有连接请求时，创建到远程地址的连接。
    4. 在两个连接之间进行数据传输。
    5. 在传输过程中，检查是否存在需要密码认证的请求头或者响应头，如果存在，则将用户名和密码信息进行替换或拼接。
    6. 对传输的数据进行读取和写入操作，实现数据的流转。

    通过这种方式，可以在不修改HTTP客户端的情况下，实现密码认证的功能。
id: "843"
---

代理http代理服务器，实现自动密码输入，让不支持密码输入的http客户端使用
```
func main() {
	logs.Info("http代理服务器的代理")
	proxyHttpProxy(":8080", "1.14.75.115:7777", "user", "pass")
}
func proxyHttpProxy(lAddr, addr, user, pass string) {
	// 构造待编码的字符串
	auth := user + ":" + pass
	// 对字符串进行base64编码
	encoded := base64.StdEncoding.EncodeToString([]byte(auth))
	authStr := fmt.Sprintf("Proxy-Authorization: Basic %s", encoded)
	listen, err := net.Listen("tcp", lAddr)
	if err != nil {
		logs.Err(err)
		return
	}
	for {
		accept, acceptErr := listen.Accept()
		if err != nil {
			logs.Err(acceptErr)
			continue
		}
		dial, err2 := net.Dial("tcp", addr)
		if err2 != nil {
			logs.Err(err2)
			continue
		}
		go hand(accept, dial, authStr)
	}
}

var regex = regexp.MustCompile(`Proxy-Authorization: Basic\s+(.*)`)
var regex2 = regexp.MustCompile(`HTTP\/1\.[0-2]`)

func hand(conn, dial net.Conn, authStr string) {
	defer conn.Close()
	go func() {
		defer dial.Close()
		re := false
		bs := make([]byte, 1024)
		for {
			var n int
			n, err2 := conn.Read(bs)
			bs = bs[:n]
			if !re {
				if regex.MatchString(string(bs)) {
					bs = regex.ReplaceAll(bs, []byte(authStr))
				} else {
					if bytes.Contains(bs, []byte("Proxy-Connection: Keep-Alive")) {
						bs = bytes.ReplaceAll(bs, []byte("Proxy-Connection: Keep-Alive"), []byte("Proxy-Connection: Keep-Alive\n"+authStr))
					} else if bytes.Contains(bs, []byte("Proxy-Connection: keep-alive")) {
						bs = bytes.ReplaceAll(bs, []byte("Proxy-Connection: keep-alive"), []byte("Proxy-Connection: Keep-Alive\n"+authStr))
					} else if regex2.MatchString(string(bs)) {
						findString := regex2.FindString(string(bs))
						bs = bytes.ReplaceAll(bs, []byte(findString), []byte(findString+"\n"+authStr))
					}
				}
				re = true
			}
			if err2 != nil {
				break
			}
			_, err2 = dial.Write(bs)
			if err2 != nil {
				break
			}
		}
	}()
	io.Copy(conn, dial)
}

```
