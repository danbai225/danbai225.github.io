---
title: IDEA常用插件以及快捷键
date: "2019-12-05 16:06:57"
updated: "2022-03-29 21:34:42"
url: https://p00q.cn/?p=9
categories:
    - 学习
tags:
    - 工具
    - IDEA
summary: idea是一个强大的开发工具，配合好用的插件和快捷键使开发更快捷便利。其中的插件包括各语言的支持插件、文件管理插件、代码生成插件、代码格式化插件等等。而快捷键则可以帮助开发者快速完成一些常见操作，如删除行、剪切行、查找类、替换、查找当前代码等等。这些功能和快捷键的使用可以大大提高开发效率。
id: "9"
---

idea是一个强大的开发工具配合好用的插件和快捷键让你开发能快捷遍历.

# 插件

* 各语言支持插件如go,python等(不用换开发工具也能开发其他语言项目了)
* .ignore 管理git提交忽略文件
* JRebel java热部署插件
* Lombok 生成java get,set方法插件
* Alibaba Java Coding Guidelines 阿里巴巴代码规范插件
* HighlightBracketPair 高亮显示当前光标括号位置
* Translation 翻译插件
* CodeGlance 右侧代码缩略图快速定位代码

# 快捷键

快捷键收集于网络

## Ctrl

Ctrl+Y    删除行

ctrl+X    剪切行

Ctrl+N   查找类

Ctrl+R   替换

Ctrl+F  当前代码中查找

Ctrl+J  自动代码提示（提示的是自己定义的代码格式）

Ctrl+D 复制行或是块(默认是这个意思)(但是我一般习惯改成专门复制行,不包括块,搜索Duplicate Lines)

Ctrl+F9     编译

Ctrl+P  方法参数提示显示        
     
Ctrl+1，2，3，4….  快速定位到书签代码处(必须先ctrl+shift+1,2,3,4…添加书签)

Ctrl+delete  删除光标后面的单词

Ctrl+BackSpace  删除光标前面的单词

Ctrl+W  选中光标所在的单词 ，连续按会有其他效果  (相反的是Ctrl+Shift+W)

Ctrl+方向左/右    光标跳到上/下个单词

Ctrl+方向上/下    相当于你用鼠标滑滚轮(为了方便鼠标党)

Ctrl+Tab  编辑窗口切换 (如果在切换的过程又加按上delete,则是关闭对应选中的窗口)  

Ctrl-F12  当前编辑的文件中快速导航(可以直接键入字母,来快速定位)

Ctrl+F3 调转到所选中的词的下一个同名位置

Ctrl+/ 或 Ctrl+Shift+/  注释（// 或者/*…*/ ）

Ctrl+Q   提示错误问题

Ctrl+B 快速打开光标处的类或方法（对于前端支持很好,比如可以直接通过class的name定位到css的文件位置）(等同于ctrl+光标指向)

Ctrl+E  最近打开的文件

Ctrl+U   前往父类的方法/父类 

Ctrl+K    VCS提交项目

Ctrl+T    VCS更新项目

Ctrl+[ / ]   移动光标到块的初/末括号地方

Ctrl+G  跳到指定行

Ctrl+小键盘+/-  折叠/展开代码

Ctrl+F1 显示错误


## Ctrl+Alt

Ctrl+Alt+方向左/右  退回/前进到上一个操作的地方

Ctrl+Alt+方向上/下  在Find模式下，挑到上/下个查找的文件

Ctrl+Alt+W  关闭所有编辑的快捷键（自己添加，在close all）

Ctrl+Alt+空格   类名或接口名提示(最常用的,一般敲入字母都会提示,但是如果你不小心esc了,可以再按这个出来)  
            
Ctrl+Alt+O 优化导入的类和包

Ctrl+Alt+L  格式化代码

Ctrl+Alt+I  选中部分自动缩进行（有点类似格式化，但是只是整理行格式而已）

Ctrl+Alt+T  选中的地方代码环绕提示

Ctrl+Alt+Enter  光标所在行上空出一行，光标跳上

Ctrl+Alt+Space 类名或接口名提示

Ctrl+Alt+O  智能导包

Ctrl+Alt+B  跳到具体的实现方法，查找抽象方法的具体实现很好用

Ctrl + Alt + V 快速引进一个变量

## Ctrl+Shift
Ctrl+Shift+F 查找文件（通过某个词，指定要搜索的文件类型，目录（跟myeclipse中的ctrl+H功能一样）） 

Ctrl+Shift+U // 大/小写都是这个快捷键

Ctrl+Shift+空格    代码补全功能(最常用的,一般敲入字母都会提示,但是如果你不小心esc了,可以再按这个出来) 

Ctrl+Shift+Up/Down  移动光标所在行/区域移动到上面/下面

Ctrl+Shift+方向左/右  选中临边左/右的单词或是符号

Ctrl+Shirt+F12  编辑器全屏

Ctrl+Shift+V 粘贴最近复制过的一些信息

Ctrl+Shift+F7 ，高亮显示所有该选中文本，按Esc高亮消失。

Ctrl+Shift+Del ，删除环绕的标签

Ctrl+Shift+Z ，取消撤销（恢复上一次操作）

Ctrl+Shift+1，2，3…   快速添加书签

Ctrl+Shift + C 复制当前文件磁盘路径到剪贴板

Ctrl+Shift + J 自动将下一行合并到当前行末尾

Ctrl+Shift + E 最近更改的文件

Ctrl+Shift + R 搜索指定范围文件，替换文字

Ctrl+Shift+[/]  选中从光标所在位置到它的父级区域

Ctrl+Shift+Space 自动补全代码（智能提示）

Ctrl+Shift+Alt+N 查找类中的方法或变量

Ctrl+Shift+小键盘+/-  折叠/展开所有代码

Ctrl+Shift+Enter  自动给末尾加;完成代码

## Alt

Alt+/  智能补全代码

Alt+F1 弹出文件选择目标，这个很好用的

Alt+F7 查看该方法/变量被调用的地方

Alt+1     打开/关闭project选项卡

Alt+A  在SVN中把新创建的文件加入进来（自己添加，在Subversion类别）

Alt+回车   快速修复(可以用来导入单个包)

Alt+Insert 生成代码(如get,set方法,构造函数等)

Alt+ left/right 切换代码视图

Alt+ Up/Down 在方法间快速移动定位

Alt+Shirt+Up/Down 移动光标所在行到上/下

Alt+Shift+N      添加任务

Alt+Shift+F   添加到收藏夹

Alt+F3  选中文本，逐个往下查找相同文本，并高亮显示。

Alt+鼠标选取，可以直接方块区域选择（很有用）

Alt + Home  跳到文件导航bar

## Shift

Shift+F6 重构：重新命名

Shift＋Click  可以关闭文件

Shift+F11     查看书签

Shift+end   选中从光标到end处

Shift+home   选中从光标到home处

Shift+Enter  光标所在行下空出一行，光标跳下

## 其他

F11      添加书签

F2 或Shift+F2 高亮错误或警告快速定位

代码标签输入完成后，按Tab，生成代码。

在Ctrl+F查找模式下，按F3下一个点

在debug模式下，F8下一步，F9下一个断点

更改下移的快捷  搜索down

