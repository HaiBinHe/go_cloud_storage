# 网盘系统

一款私有云存储的网盘系统。不包括前端内容，目前还在完善....

## 目录结构

```
├─api
├─cmd
├─conf
├─docs
├─internal
│  ├─cache
│  ├─middleware
│  ├─model
│  │  └─test
│  ├─mq
│  │  └─transfer
│  └─service
│      ├─file
│      ├─share
│      ├─upload
│      └─userService
├─pkg
│  ├─app
│  ├─auth
│  ├─error
│  ├─logger
│  │  └─storage
│  │      └─logs
│  └─response
├─routers
│  └─api
├─storage
│  └─logs
├─tools
└─web
    └─view

```

## 特性

☑ 支持分块上传和断点续传

☑ 使用Redis中的Stream做消息队列，来实现将文件异步存储至七牛云。

## 后续更新

