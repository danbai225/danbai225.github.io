---
title: socks5协议Go实现
date: 2021-08-07 14:34:18.718
updated: 2021-08-07 14:42:22.936
url: https://p00q.cn/?p=354
categories: 
- Go
- 开发
tags: 
- socks5
- 协议实现
---

# Socks5协议介绍
socks5 是 socks4 升级版的协议相比 4，本代协议版本多了 udp 代理。
socks5 代理比 http 协议更低层次并且支持 tcp 代理。更多内容详见[百科](https://zh.wikipedia.org/wiki/SOCKS)。

# Go的相关实现
下面是一些 github 上的协议实现

- [armon/go-socks5](https://github.com/armon/go-socks5)
- [txthinking/socks5](https://github.com/txthinking/socks5)

# 协议基本代理流程

客户端->建立连接->服务端
客户端<-验证成功<-服务端
客户端->我要访问百度->服务端->访问百度->百度
客户端<-给你百度数据<-服务端<-返回数据<-百度

# 自己动手开撸

为了更熟悉协议，自己用 go 写一下基础功能

## 客户端
```GO
package socks5

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	Host     string
	Port     string
	UserName string
	Password string
}

//连接
func (c *Client) conn() (net.Conn, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))
	if err != nil {
		return nil, err
	}
	err = c.auth(conn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

//认证
func (c *Client) auth(conn net.Conn) error {
	//组织发送支持的认证方法
	authPackage := AuthPackage{}
	if c.UserName != "" && c.Password != "" {
		authPackage.addMethod(AccountPasswordAuthentication)
	}
	authPackage.addMethod(NoAuthenticationRequired)
	_, err := conn.Write(authPackage.toData())
	if err != nil {
		return err
	}
	data := make([]byte, 2)
	l, err := conn.Read(data)
	if err != nil {
		return err
	}
	if l != 2 {
		return errors.New("返回数据有误非两个字节")
	}
	if data[0] != Version {
		return errors.New("当前协议Socks5与服务端协议不匹配")
	}
	buffer := bytes.Buffer{}
	switch data[1] {
	case NoAuthenticationRequired:
		return nil
	case AccountPasswordAuthentication:
		//认证协议 0x01
		buffer.WriteByte(0x01)
		//用户名长度
		buffer.WriteByte(byte(len(c.UserName)))
		//用户名
		buffer.WriteString(c.UserName)
		//密码长度
		buffer.WriteByte(byte(len(c.Password)))
		//密码
		buffer.WriteString(c.Password)
	}
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	l, err = conn.Read(data)
	if err != nil {
		return err
	}
	if l != 2 {
		return errors.New("返回数据有误非两个字节")
	}
	if data[0] != 0x01 {
		return errors.New("当前认证协议Socks5与服务端协议不匹配")
	}
	if data[1] > 0 {
		return errors.New("认证失败")
	}
	return nil
}

//向服务器发送连接请求
func (c *Client) requisition(conn net.Conn, host string, port uint16, cmd uint8) (net.Conn, error) {
	Type, addr, err := addressResolution(host)
	if err != nil {
		return nil, err
	}
	buffer := bytes.Buffer{}
	buffer.Write([]byte{Version, cmd, Zero, Type})
	buffer.Write(addr)
	//写入端口
	buffer.Write(portToBytes(port))
	_, err = conn.Write(buffer.Bytes())
	if err != nil {
		return nil, err
	}
	read, err := conn.Read(addr)
	if err != nil {
		return nil, err
	}
	rdata := addr[:read]
	if rdata[1] != Zero {
		return nil, errors.New(fmt.Sprintf("请求错误:%x", rdata[1]))
	}
	if cmd == UDP {
		ipdata := make([]byte, 1024)
		if len(rdata) == 4 {
			n, err := conn.Read(ipdata)
			if err != nil {
				return nil, err
			}
			ipdata = ipdata[:n]
		} else {
			ipdata = rdata[4:]
		}
		addr, err := addressResolutionFormByteArray(ipdata, rdata[3])
		if err != nil {
			if conn != nil {
				conn.Close()
			}
			return nil, err
		}
		dial, err := net.Dial("udp", addr)
		if err != nil {
			conn.Close()
			return nil, err
		}
		return dial, nil
	}
	return nil, nil
}
func (c *Client) udp(conn net.Conn, host string, port uint16) (net.Conn, error) {
	return c.requisition(conn, host, port, UDP)
}
func (c *Client) bind(conn net.Conn, host string, port uint16) error {
	_, err := c.requisition(conn, host, port, Bind)
	return err
}
func (c *Client) tcp(conn net.Conn, host string, port uint16) error {
	_, err := c.requisition(conn, host, port, Connect)
	return err
}
func (c *Client) TcpProxy(host string, port uint16) (net.Conn, error) {
	conn, err := c.conn()
	if err != nil {
		return nil, err
	}
	return conn, c.tcp(conn, host, port)
}
func (c *Client) GetHttpProxyClient() *http.Client {
	httpTransport := &http.Transport{}
	httpTransport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		split := strings.Split(addr, ":")
		if len(split) < 2 {
			return c.TcpProxy(split[0], 80)
		}
		port, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		return c.TcpProxy(split[0], uint16(port))
	}
	return &http.Client{Transport: httpTransport}
}
func (c *Client) GetHttpProxyClientSpecify(transport *http.Transport, jar http.CookieJar, CheckRedirect func(req *http.Request, via []*http.Request) error, Timeout time.Duration) *http.Client {
	transport.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		split := strings.Split(addr, ":")
		if len(split) < 2 {
			return c.TcpProxy(split[0], 80)
		}
		port, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		return c.TcpProxy(split[0], uint16(port))
	}
	return &http.Client{Transport: transport, Jar: jar, CheckRedirect: CheckRedirect, Timeout: Timeout}
}
func (c *Client) UdpProxy(host string, port uint16) (*UdpProxy, error) {
	conn, err := c.conn()
	if err != nil {
		return nil, err
	}
	udp, err := c.udp(conn, "0.0.0.0", 0)
	if err != nil {
		return nil, err
	}
	proxy := UdpProxy{udp, host, port}
	return &proxy, nil
}

```
## 服务端
```GO
package socks5

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	listen  net.Listener
	Config  Config
	UserMap map[string]string
}
type User struct {
	Name     []byte
	Password []byte
}
type Config struct {
	host      string
	port      uint16
	BlackList []string
	AuthList  []uint8
}

func (s *Server) Start() (err error) {
	s.listen, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.Config.host, s.Config.port))
	if err != nil {
		return err
	}
	for s.listen != nil {
		accept, err := s.listen.Accept()
		if err == nil {
			go s.newConn(accept)
		}
	}
	return nil
}
func (s *Server) newConn(conn net.Conn) {
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	err2 := s.auth(conn)
	if err2 != nil {
		log.Println(err2)
		return
	}
	bytes := make([]byte, 1024)
	read, err := conn.Read(bytes)
	if err != nil {
		return
	}
	bytes = bytes[:read]
	if len(bytes) < 5 {
		return
	}
	cmd := bytes[1]
	array, err2 := addressResolutionFormByteArray(bytes[4:], bytes[3])
	if err2 != nil {
		return
	}

	switch cmd {
	case Connect:
		dial, err2 := net.Dial("tcp", array)
		if err2 != nil {
			return
		}
		addr := dial.LocalAddr().String()
		split := strings.Split(addr, ":")
		port := split[len(split)-1]
		parseUint, _ := strconv.ParseUint(port, 10, 16)
		Type, host, _ := addressResolution(strings.Join(split[:len(split)-1], ":"))
		p := make([]byte, 2)
		binary.BigEndian.PutUint16(p, uint16(parseUint))
		data := []byte{Version, Zero, Zero, Type}
		data = append(data, host...)
		data = append(data, p...)
		conn.Write(data)
		ioCopy(conn, dial)
	case Bind:
	case UDP:
	}
}
func (s *Server) auth(conn net.Conn) error {
	bytes := make([]byte, 16)
	read, err := conn.Read(bytes)
	if err != nil {
		return err
	}
	bytes = bytes[:read]
	if len(bytes) < 3 {
		return errors.New("认证数据长度不符")
	}
	if bytes[0] != Version {
		return errors.New("协议不符合")
	}
	//支持的认证方法
	moth := uint8(Zero)
	for _, u := range s.Config.AuthList {
		for _, b := range bytes[2:] {
			if u == b {
				moth = u
				if moth == Zero {
					conn.Write([]byte{Version, Zero})
					return nil
				}
				break
			}
		}
	}
	if moth == Zero {
		return errors.New("没有支持的认证方法")
	}
	_, err = conn.Write([]byte{Version, moth})
	if err != nil {
		return err
	}
	switch moth {
	case NoAuthenticationRequired:
		return nil
	case AccountPasswordAuthentication:
		bytes = make([]byte, 1024)
		n, err := conn.Read(bytes)
		if err != nil {
			return err
		}
		bytes = bytes[:n]
		if len(bytes) < 3 {
			return errors.New("认证数据长度不符")
		}
		if bytes[0] != 0x01 {
			return errors.New("认证协议不符合")
		}
		username := bytes[2 : 2+bytes[1]]
		password := bytes[3+bytes[1]:]
		if pas, has := s.UserMap[string(username)]; has {
			if pas == string(password) {
				conn.Write([]byte{0x01, Zero})
				return nil
			}
		}
	}
	return nil
}

```

# Github
[socks5](https://github.com/danbai225/socks5)
# 参考资料

[[翻译] RFC 1928: SOCKS 协议第 5 版](https://luyuhuang.tech/2020/08/27/rfc1928.html#7-%E5%9F%BA%E4%BA%8E-udp-%E5%AE%A2%E6%88%B7%E7%AB%AF%E7%9A%84%E8%BF%87%E7%A8%8B)

[维基百科-SOCKS](https://zh.wikipedia.org/wiki/SOCKS)

[Socks5代理协议](https://juejin.cn/post/6844903923518537741)