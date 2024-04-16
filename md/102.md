---
title: java判断压缩包内是否存在指定文件
date: "2020-05-20 09:02:03"
updated: "2023-03-08 18:55:31"
url: https://p00q.cn/?p=102
categories:
    - Java
tags:
    - zip
summary: |-
    以上代码是用于判断压缩包内是否存在名为index.html的文件。

    其中，通过`ZipInputStream`从输入流中读取压缩包的文件流。通过`zin.getNextEntry()`获得压缩包中的文件，然后判断文件名是否为index.html。若存在index.html文件，则将变量`ifThere`置为true，跳出循环。

    最后，根据`ifThere`的值来判断是否存在index.html文件，并返回对应的响应。

    另外，如果有文件对象，可以直接使用`ZipFile`来读取压缩包内的文件。可以通过`zipfile.getEntry("index.html")`获取index.html文件的入口，并通过`.getName()`获取文件名。
id: "102"
---

业务需求,需要判断压缩包内是否存在index.html文件.不通过解压进行判断.
通过ZipInputStream获取文件流进行判断.

上代码:
```
//判断是否存在index.html
        boolean ifThere=false;
        try {
            InputStream in = new BufferedInputStream(file.getInputStream());
            ZipInputStream zin = new ZipInputStream(in,Charset.forName("gbk"));
            ZipEntry ze;
            while ((ze = zin.getNextEntry()) != null) {
                if (ze.isDirectory()) {
                } else {
                    String name = ze.getName();
                    if(name.toLowerCase().equals(WallpaperConstant.INDEX)){
                        ifThere=true;
                        break;
                    }
                }
            }
            zin.closeEntry();
        } catch (IOException e) {
            e.printStackTrace();
        }
        if(!ifThere){
            return Response.Err(WallpaperConstant.NOT_FIND_INDEX);
        }
```
上面的代码是读取http请求中的文件流
如果有文件对象能直接读取指定文件
```
ZipFile zipfile=new ZipFile(file.getInputStream(),Charset.forName("gbk"));
System.out.println(zipfile.getEntry("index.html").getName());
```
