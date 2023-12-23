---
title: 让java switch支持表达式和多匹配
date: 2020-07-22 09:54:19.968
updated: 2020-07-22 09:54:19.968
url: https://p00q.cn/?p=131
categories: 
- Java
tags: 
- java
- Choice
---

# choice
先上地址:[Github](https://github.com/danbai225/choice)
maven:
```
<dependency>
    <groupId>cn.p00q.choice</groupId>
    <artifactId>choice</artifactId>
    <version>1.0</version>
</dependency>
```
让java支持表达式和多匹配的实现.
使用例子:
```
 int a = 1, b = 2, c = 3, d = 4;
        new Choice(true).add(1L, () -> {
            System.out.println("这里是1");
        }).add(a < b ? c : d > a ? b : c, () -> {
            System.out.println("这里是2");
        }).add(3.0f, () -> {
            System.out.println("这里是3");
        }).add(Color.BLUE, () -> {
            System.out.println("这里是4");
        }).add(() -> {
            System.out.println("多个匹配");
        }, 1, 2, 3, 4, 5, 6, 7, 8, 9).Default(() -> {
            System.out.println("这里是默认方法");
        }).execute(4);
```
使用了Lambda表达式的特性.不仅能这样用,还能可以这样用:
在A类创建一个公开的静态Choice
```
public static Choice socketChoice;
```
在不同的任务类中去实现`Function`接口的`run`方法或者直接添加Lambda表达式.即可实现在不同请求调用不同方法的功能.
```
socketChoice.execute(请求值)
```
# 实现
java并不能直接像js那样将函数直接视为变量传入.所以通常以接口(只能有一个方法)当作函数进行传入.

接口定义:
```
public interface Function {
    /**
     * 运行方法
     */
    void run();

}
```
完整实现:
```
public class Choice {
    /**
     * 方法map
     */
    private Map<Object, Function> map;
    /**
     * 默认方法
     */
    private Function Default;
    /**
     * 表达式支持
     */
    private boolean expression;

    public Choice() {
        map = new ConcurrentHashMap();
    }

    public Choice(boolean expression) {
        this();
        this.expression = expression;
    }

    public Choice add(Object v, Function function) {
        if (v != null && function != null) {
            map.put(v, function);
        }
        return this;
    }

    public Choice add(Function function, Object... v) {
        if (function != null) {
            for (Object iv : v) {
                if (iv != null) {
                    map.put(iv, function);
                }
            }
        }
        return this;
    }

    /**
     * 不传值的执行
     */
    public void execute() {
        AtomicBoolean flg = new AtomicBoolean(false);
        //遍历map
        for (Map.Entry<Object, Function> entry : map.entrySet()) {
            Object iv = entry.getKey();
            if (iv != null) {
                Function iFun = entry.getValue();
                //类型判断
                if (expression && iv.getClass().equals(Boolean.class)) {
                    if (iv.equals(true)) {
                        //类型一样eq为true 执行方法
                        iFun.run();
                        flg.set(true);
                        return;
                    }
                }
            }
        }
        if (!flg.get() && Default != null) {
            Default.run();
        }
    }

    /**
     * 传值的执行
     *
     * @param v 匹配值
     */
    public void execute(Object v) {
        if (v == null) {
            return;
        }
        AtomicBoolean flg = new AtomicBoolean(false);
        Function fastTrue = null;
        for (Map.Entry<Object, Function> entry : map.entrySet()) {
            Object iv = entry.getKey();
            Function iFun = entry.getValue();
            //类型判断
            if (v.getClass().equals(iv.getClass())) {
                if (v.equals(iv)) {
                    iFun.run();
                    flg.set(true);
                    return;
                }
            }
            //表达式
            if (expression && iv.getClass().equals(Boolean.class)) {
                if (v.equals(true)) {
                    fastTrue = iFun;
                    flg.set(true);
                }
            }
        }
        if (flg.get()) {
            if (expression && fastTrue != null) {
                fastTrue.run();
            }
        } else if (Default != null) {
            Default.run();
        }
    }

    public Choice Default(Function function) {
        Default = function;
        return this;
    }

}
```
