---
title: spring boot上传文件并自动解压
date: "2020-05-15 08:35:44"
updated: "2020-05-20 09:13:16"
url: https://p00q.cn/?p=101
categories:
    - Java
tags:
    - spring boot
    - zip4j
summary: 这段代码展示了如何将上传的文件进行解压到指定目录的业务实现。首先，使用zip4j库进行文件解压，该库支持加密、解密压缩，以及文件的添加、删除等功能。然后，在控制器中判断上传文件的后缀是否是.zip，如果是，则调用服务进行文件解压。服务部分主要实现了将元数据保存到数据库，并保存文件后进行解压的操作。文件工具类中的方法用于保存文件、获取应用程序路径以及进行zip文件的解压操作。最后，调用解压方法将文件解压到指定目录，并返回解压成功的提示信息。
id: "101"
---

业务需要把上传文件解压到指定目录
# 依赖
zip4j功能比较强大，支持加密、解密压缩，支持文件的添加、删除等
```xml
        <dependency>
            <groupId>net.lingala.zip4j</groupId>
            <artifactId>zip4j</artifactId>
            <version>1.3.2</version>
        </dependency>
```
# 控制器
判断文件后缀是不是.zip 进行服务调用
```java
@PostMapping("/uploadWallpaper")
    @ResponseBody
    public Response uploadFile(@RequestParam("file") MultipartFile file,Wallpaper wallpaper){
        //后缀.zip
        if(FileUtils.getPrefix(file.getOriginalFilename()).toLowerCase().equals(Wallpaper.File_PREFIX)){
            return wallpaperService.saveWallpaper(file,wallpaper);
        }
        return Response.Err("不是.zip结尾的压缩包");
    }
```
# 服务
主要业务实现,把元数据存入数据库.保存文件(zip)后进行解压.
```java
public Response saveWallpaper(MultipartFile file, Wallpaper wallpaper){
        //设置id(行数+1)
        wallpaper.setId(WallpaperSize()+1);
        wallpaper.setUrl("127.0.0.1:8081/wallpaper/"+wallpaper.getId()+"/index.html");
        wallpaper.setAudit(false);
        //添加到数据库
        wallpaperMapper.insert(wallpaper);
        String path=FileUtils.getApplicationPath()+"wallpaper"+File.separator+wallpaper.getId()+"."+Wallpaper.File_PREFIX;
        //保存到服务器
        FileUtils.saveFile(file,path);
        String destPath=FileUtils.getApplicationPath()+"wallpaper"+File.separator+wallpaper.getId()+File.separator;
        //解压
            if(FileUtils.unPackZip(new File(path),"",destPath)){
            return Response.Ok("上传成功",wallpaper.getUrl());
        }
        return Response.Err("上传失败");
    }
```
# 文件工具类
文件保存,文件解压的工具类
```java
public class FileUtils {
    public static String getPrefix(String filename){
        return filename.substring(filename.lastIndexOf(".")+1);
    }
    public static boolean saveFile(MultipartFile file,String path){
        File desFile = new File(path);
        if(!desFile.getParentFile().exists()){
            desFile.mkdirs();
        }
        try {
            file.transferTo(desFile);
        } catch (IOException e) {
            e.printStackTrace();
            return false;
        }
        return true;
    }
    public static String getApplicationPath(){
        //获取classpath
        ApplicationHome h = new ApplicationHome(WallpaperApplication.class);
        File jarF = h.getSource();
        return jarF.getParentFile()+ File.separator;
    }
    /**
     * zip文件解压
     *
     * @param destPath 解压文件路径
     * @param zipFile  压缩文件
     * @param password 解压密码(如果有)
     */
    public static boolean unPackZip(File zipFile, String password, String destPath) {
        try {
            ZipFile zip = new ZipFile(zipFile);
            zip.extractAll(destPath);
            // 如果解压需要密码
            if (zip.isEncrypted()) {
                zip.setPassword(password);
            }
        } catch (Exception e) {
            System.out.println(e);
            return false;
        }
        return true;
    }
}
```
