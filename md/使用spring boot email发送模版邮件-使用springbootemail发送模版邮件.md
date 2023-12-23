---
title: 使用spring boot email发送模版邮件
date: "2020-08-03 11:38:25"
updated: "2021-06-10 15:25:41"
url: https://p00q.cn/?p=135
categories:
    - Java
tags:
    - email
summary: 本文介绍了在项目中发送邮件时的两种方式：使用原始的javax.mail发送简单邮件和使用html模板发送邮件。原始方式发送的邮件容易被识别成垃圾邮件，而使用html模板发送能够更好地呈现邮件内容。使用模板发送邮件的步骤包括添加相关依赖、设置模板内容、创建MimeMessage对象并设置相关信息，最后发送邮件。给出了具体的代码示例，方便读者参考。
id: "135"
---

# 缘由

在项目中最早我用原始的`javax.mail`
```
public void sendEmail(final String title, final String content, final String toMail) {
        try {
            SimpleMailMessage mailMessage = new SimpleMailMessage();
            String nick=javax.mail.internet.MimeUtility.encodeText("淡白影视");
            mailMessage.setFrom(String.valueOf(new InternetAddress(nick+" <"+FROM+">")));
            mailMessage.setSubject(title);
            mailMessage.setText(content);
            String[] toAddress = toMail.split(",");
            mailMessage.setTo(toAddress);
            //发送邮件
            javaMailSender.send(mailMessage);
        }catch(Exception e){
            System.out.println(e);
```
发送简单几句话 很容易被识别成垃圾邮件.

# 用html模版发送
添加依赖
```
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-mail</artifactId>
        </dependency>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-thymeleaf</artifactId>
        </dependency>
```
发送模版:
```
public void sendTemplateMail(String receiver, String subject, String emailTemplate, Map<String, Object> dataMap) throws Exception {
        Context context = new Context();
        for (Map.Entry<String, Object> entry : dataMap.entrySet()) {
            context.setVariable(entry.getKey(), entry.getValue());
        }
        String templateContent = templateEngine.process(emailTemplate, context);
        MimeMessage message = javaMailSender.createMimeMessage();
        MimeMessageHelper helper = new MimeMessageHelper(message, true);
        helper.setFrom(sender);
        helper.setTo(receiver);
        helper.setSubject(subject);
        helper.setText(templateContent, true);
        javaMailSender.send(message);
    }
```
使用:
```
 //发送激活链接邮件
        SimpleDateFormat sdf = new SimpleDateFormat("yyyy年MM月dd日 HH时mm分ss秒");
        String subject = "淡白影视验证码";
        String emailTemplate = "ValidationEmailTemplate";
        Map<String, Object> dataMap = new HashMap<>();
        dataMap.put("email", email);
        dataMap.put("code", yzm);
        dataMap.put("createTime", sdf.format(new Date()));
        try {
            emailUtil.sendTemplateMail(email, subject, emailTemplate, dataMap);
        } catch (Exception e) {
            return;
        }
```
[详细代码](https://github.com/Programming-With-Love/dbys/blob/master/src/main/java/com/danbai/ys/utils/EmailUtil.java)
