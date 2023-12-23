---
title: 修复一个open too many files
date: 2021-09-09 18:00:03.538
updated: 2021-09-10 10:05:26.246
url: https://p00q.cn/?p=390
categories: 
- Go
tags: 
- Go
- socket
---

在使用[socks5代理客户端](github.com/shikanon/socks5proxy)发现，很多连接没有被正常关闭。
在文件`client.go`中当代理类型不为`http`时会执行以下代码


```
	// 和远程端建立安全信道
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// 本地的内容copy到远程端
	go func() {
		defer wg.Done()
		SecureCopy(localClient, dstServer, auth.Encrypt)
	}()

	// 远程得到的内容copy到源地址
	go func() {
		defer wg.Done()
		SecureCopy(dstServer, localClient, auth.Decrypt)
	}()
	wg.Wait()
```
在`cryptogram.go`中的copy方法如下
```
// 加密io复制，可接收加密函数作为参数
func SecureCopy(src io.ReadWriteCloser, dst io.ReadWriteCloser, secure func(b []byte) error) (written int64, err error) {
	size := 1024
	buf := make([]byte, size)
	for {
		nr, er := src.Read(buf)
		secure(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}
	return written, err

```
当有一方连接断开时，另一方链接可能不会被正常关闭导致锁一直没有释放。本地连接也一直没有关闭。打开了太多连接也就出现了报错。

修改后的代码：
`cryptogram.go`
```
func SecureCopy(src io.ReadWriteCloser, dst io.ReadWriteCloser, secure func(b []byte) error) (written int64, err error) {
	defer func() {
		if src != nil {
			src.Close()
		}
		if dst != nil {
			dst.Close()
		}
	}()
	buf := make([]byte, 1024)
	for {
		nr, er := src.Read(buf)
		if er != nil {
			break
		}
		if nr > 0 {
			err = secure(buf)
			if err != nil {
				return written, err
			}
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}
	}
	return written, err
}
```

`client.go` 把连接的关闭放到copy内部 当一方连接被关闭则关闭两个连接，另一个拷贝就能感知从而读取错误正常关闭并退出。

```
	go SecureCopy(dstServer, localClient, auth.Decrypt)
	SecureCopy(localClient, dstServer, auth.Encrypt)
```