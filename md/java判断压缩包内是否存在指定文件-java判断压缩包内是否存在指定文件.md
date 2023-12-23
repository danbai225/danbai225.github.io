---
title: java判断压缩包内是否存在指定文件
date: 2020-05-20 09:02:03.975
updated: 2023-03-08 18:55:31.697
url: https://p00q.cn/?p=102
categories: 
- Java
tags: 
- zip
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