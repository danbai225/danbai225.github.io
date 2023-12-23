---
title: spring boot映射所在目录下的静态资源
date: 2020-05-14 15:48:10.198
updated: 2020-06-08 17:58:02.642
url: https://p00q.cn/?p=100
categories: 
- Java
tags: 
- spring boot
---

# SpringMvcConfig
```java
@Configuration
public class MyWebMvcConfigurer implements WebMvcConfigurer {
    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        //获取classpath
        ApplicationHome h = new ApplicationHome(getClass());
        File jarF = h.getSource();
        //拼接path classpath的上级+自定义的images
        String gitPath=jarF.getParentFile()+ File.separator+"images"+File.separator;
        //添加映射  (idea打包会映射在target目录)
        registry.addResourceHandler("/images/**").addResourceLocations("file:"+gitPath);
    }

}

```
jar包放在D:\TEST,目录下运行
访问:
http://127.0.0.1:8081/images/1.png
即可访问到:
D:\TEST\images\1.png