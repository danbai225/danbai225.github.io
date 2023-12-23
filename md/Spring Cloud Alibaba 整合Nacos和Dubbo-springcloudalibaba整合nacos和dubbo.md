---
title: Spring Cloud Alibaba 整合Nacos和Dubbo
date: 2020-06-04 16:18:25.78
updated: 2021-06-10 15:19:58.982
url: https://p00q.cn/?p=103
categories: 
- Java
tags: 
- Spring Cloud Alibaba
- nacos
- dubbo
---

# Spring Cloud Alibaba

Spring Cloud Alibaba 致力于提供微服务开发的一站式解决方案。此项目包含开发分布式应用微服务的必需组件，方便开发者通过 Spring Cloud 编程模型轻松使用这些组件来开发分布式应用服务。

依托 Spring Cloud Alibaba，您只需要添加一些注解和少量配置，就可以将 Spring Cloud 应用接入阿里微服务解决方案，通过阿里中间件来迅速搭建分布式应用系统。

[Spring Cloud Alibaba Github](https://github.com/alibaba/spring-cloud-alibaba)
# Nacos
什么是 Nacos？
服务（Service）是 Nacos 世界的一等公民。Nacos 支持几乎所有主流类型的“服务”的发现、配置和管理。
[Nacos官网](https://nacos.io/)
# Apache Dubbo
Apache Dubbo™ 是一款高性能Java RPC框架。
[Apache Dubbo](https://dubbo.apache.org)
# 搭建项目
项目结构
```
├─ demo
│   ├─ demo-api #api
│   ├─ demo-consumer #消费者
│   ├─ demo-provider #提供者
│   └─ pom.xml
├─ dubbo 
│   └─ pom.xml #dubbo依赖
├─ nacos
│   └─ pom.xml #nacos依赖
├─ spring-cloud-alibaba-starter 起始依赖
│   └─ pom.xml
└─ pom.xml
```
## 根POM.xml
基于spring-boot 2.3.0RELEASE﹑spring-cloud-alibaba2.2.1.RELEASE
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>
    <parent>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-parent</artifactId>
        <version>2.3.0.RELEASE</version>
        <relativePath/> <!-- lookup parent from repository -->
    </parent>
    <packaging>pom</packaging>
    <groupId>cn.p00q</groupId>
    <artifactId>spring-cloud-alibaba-parent</artifactId>
    <version>1.0.0-SNAPSHOT</version>
    <name>spring-cloud-alibaba-parent</name>
    <description>微服务结构项目</description>

    <properties>
        <java.version>1.8</java.version>
        <spring-cloud-alibaba.version>2.2.1.RELEASE</spring-cloud-alibaba.version>
    </properties>

    <dependencyManagement>
        <dependencies>
            <dependency>
                <groupId>com.alibaba.cloud</groupId>
                <artifactId>spring-cloud-alibaba-dependencies</artifactId>
                <version>${spring-cloud-alibaba.version}</version>
                <type>pom</type>
                <scope>import</scope>
            </dependency>
        </dependencies>
    </dependencyManagement>
</project>
```
## Nacos依赖
以根项目创建Nacos的POM项目
```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>spring-cloud-alibaba-parent</artifactId>
        <groupId>cn.p00q</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>
    <artifactId>nacos</artifactId>
    <dependencies>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-discovery</artifactId>
        </dependency>
    </dependencies>
</project>
```
## Dubbo依赖
```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>spring-cloud-alibaba-parent</artifactId>
        <groupId>cn.p00q</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>dubbo</artifactId>
    <dependencies>
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-dubbo</artifactId>
        </dependency>
    </dependencies>
</project>
```
## 起始依赖（spring-cloud-alibaba-starter）
```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>spring-cloud-alibaba-parent</artifactId>
        <groupId>cn.p00q</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>

    <artifactId>spring-cloud-alibaba-starter</artifactId>
    <packaging>pom</packaging>

    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-web</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-test</artifactId>
            <scope>test</scope>
            <exclusions>
                <exclusion>
                    <groupId>org.junit.vintage</groupId>
                    <artifactId>junit-vintage-engine</artifactId>
                </exclusion>
            </exclusions>
        </dependency>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>dubbo</artifactId>
            <version>${parent.version}</version>
        </dependency>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>nacos</artifactId>
            <version>${parent.version}</version>
        </dependency>
    </dependencies>
</project>
```
## 正式Demo项目开始

### Api
```
package cn.p00q.cloud.demo.api;

/**
 * @program: AlibabaCloud
 * @description: 例子Api
 * @author: DanBai
 * @create: 2020-06-02 20:09
 **/
public interface DemoService {

    /**
     * 返回信息
     * @param name
     * @return
     */
    String echo(String name);
}

```
### 服务提供者
### pom.xml
```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>demo</artifactId>
        <groupId>cn.p00q</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>
    <artifactId>demo-provider</artifactId>
    <dependencies>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>spring-cloud-alibaba-starter</artifactId>
            <version>${parent.version}</version>
        </dependency>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>demo-api</artifactId>
            <version>${parent.version}</version>
        </dependency>
    </dependencies>
</project>
```
#### 实现接口
Service接口是org.apache.dubbo.config.annotation.Service包下的
```
package cn.p00q.cloud.demo.service.impl;

import cn.p00q.cloud.demo.api.DemoService;
import org.apache.dubbo.config.annotation.Service;
import org.apache.dubbo.rpc.RpcContext;

/**
 * @program: AlibabaCloud
 * @description: DemoService实现
 * @author: DanBai
 * @create: 2020-06-02 20:39
 **/
@Service
public class DemoServiceImpl implements DemoService {
    @Override
    public String echo(String name) {
        RpcContext rpcContext = RpcContext.getContext();
        return String.format("Service [name :%s , port : %d] %s(\"%s\") : Hello,%s",
                "demo",
                rpcContext.getLocalPort(),
                rpcContext.getMethodName(),
                name,
                name);
    }
}

```
#### 启动类
```
package cn.p00q.cloud.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

/**
 * @program: AlibabaCloud
 * @description: demo提供应用
 * @author: DanBai
 * @create: 2020-06-02 20:35
 **/
@EnableDiscoveryClient
@SpringBootApplication
public class DemoProviderApplication {
    public static void main(String[] args) {
        SpringApplication.run(DemoProviderApplication.class, args);
    }
}
```
#### 配置
```yaml
dubbo:
  protocol:
    name: dubbo #通信协议 
    port: -1 #端口自增
  registry:
    address: nacos://127.0.0.1:8848 #nacos服务器地址
  scan:
    base-packages: cn.p00q.cloud.demo.service.impl #扫描服务包路径
server:
  port: 8001
spring:
  application:
    name: demo-provider
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848 #nacos服务器地址
```

### 消费者
#### pom.xml

```
<?xml version="1.0" encoding="UTF-8"?>
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
    <parent>
        <artifactId>demo</artifactId>
        <groupId>cn.p00q</groupId>
        <version>1.0.0-SNAPSHOT</version>
    </parent>
    <modelVersion>4.0.0</modelVersion>
    <artifactId>demo-consumer</artifactId>
    <dependencies>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>spring-cloud-alibaba-starter</artifactId>
            <version>${parent.version}</version>
        </dependency>
        <dependency>
            <groupId>cn.p00q</groupId>
            <artifactId>demo-api</artifactId>
            <version>${parent.version}</version>
        </dependency>
    </dependencies>
</project>
```
#### 调用服务
```
package cn.p00q.cloud.demo.controller;

import cn.p00q.cloud.demo.api.DemoService;
import org.apache.dubbo.config.annotation.Reference;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @program: AlibabaCloud
 * @description: Demo
 * @author: DanBai
 * @create: 2020-06-03 10:18
 **/
@RestController
public class DemoController {
    @Reference
    DemoService demoService;

    @GetMapping("/demo")
    public String getDemo(){
        return demoService.echo("danbai");
    }
}

```
#### 启动类
```
package cn.p00q.cloud.demo;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;

/**
 * @program: AlibabaCloud
 * @description: demo消费者启动类
 * @author: DanBai
 * @create: 2020-06-03 10:13
 **/
@EnableDiscoveryClient
@SpringBootApplication
public class DemoConsumerApplication {
    public static void main(String[] args) {
        SpringApplication.run(DemoConsumerApplication.class, args);
    }
}

```
#### 配置
```
dubbo:
  cloud:
    subscribed-services: demo-provider
  protocol:
    name: dubbo
    port: -1
  registry:
    address: nacos://127.0.0.1:8848
server:
  port: 8002
spring:
  application:
    name: demo-consumer
  cloud:
    nacos:
      discovery:
        server-addr: 127.0.0.1:8848

```
### 启动两个提供者 一个消费者
访问http://localhost:8002/demo
```
Service [name :demo , port : 20880] echo("danbai") : Hello,danbai
```
刷新
```
Service [name :demo , port : 20881] echo("danbai") : Hello,danbai
```
dubbo均衡负载

### Nacos配置中心

#### 在Nacos项目下的pom.xml增加依赖
```
        <dependency>
            <groupId>com.alibaba.cloud</groupId>
            <artifactId>spring-cloud-starter-alibaba-nacos-config</artifactId>
        </dependency>
```
#### 上传配置

登录nacos后台，进入配置列表在public命名空间下添加配置。
Data ID分别为demo-provider-dev.yaml和demo-consumer-dev.yaml。
组默认DEFAULT_GROUP
把各自的application.yaml添加进去

#### 配置修改

在消费者和提供者resources目录下添加bootstrap.yaml文件
```
spring:
  application:
    name: demo-provider #消费者为demo-consumer
  cloud:
    nacos:
      config:
        server-addr: 127.0.0.1:8848
        file-extension: yaml
  profiles:
    active: dev
```
完成配置

[本项目地址](https://github.com/danbai225/AlibabaCloud)

# 工具
[在线yaml转properties-在线properties转yaml-ToYaml.com](https://toyaml.com/index.html)