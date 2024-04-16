---
title: spring boot映射所在目录下的静态资源
date: "2020-05-14 15:48:10"
updated: "2020-06-08 17:58:02"
url: https://p00q.cn/?p=100
categories:
    - Java
tags:
    - spring boot
summary: |-
    这段代码是一个Spring MVC的配置类。它通过实现WebMvcConfigurer接口来自定义资源处理器，实现了addResourceHandlers方法。在这个方法中，它首先获取了应用的classpath路径，并拼接上自定义的路径/images。然后通过addResourceHandler方法将访问路径/images/**映射到实际文件的路径file:gitPath上。

    在运行这个jar包的时候，可以通过访问http://127.0.0.1:8081/images/1.png来访问到D:\TEST\images\1.png这个文件。也就是说，这个配置类为指定路径下的文件提供了访问的接口。
id: "100"
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
