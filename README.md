# 网盘系统

一款私有云存储的网盘系统。不包括前端内容，目前还在完善....

## 目录结构

```
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

- [x] 支持分块上传、断点续传、文件分享、文件秒传

- [x] 使用Redis中的Stream做消息队列，来实现将文件异步存储至七牛云。

- [x] 自定义日志框架，定义了两个层面的抽象，一个Logger统一了日志接入的方式，一个Helper统一了日志调用的方式。

## 后续更新

- [ ] 实现多用户
- [ ] 离线下载
- [ ] 客户端直传



