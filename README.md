# fox
fox 小程序服务端

### 背景
在生活中使用小程序频率越来越来多，有些好奇心，随想尝试自己做一个...

### 项目技术
* golang gin框架
* jwt 做restful接口token验证
* 使用mysql数据库，gorm可以快速开发
* redis 使用go-redis库

### 项目功能
* 1.获取wxUser accessToken.
* 2.定时刷新accessToken.
* 3.accessToken 存储数据库和Redis.
* 4.统一接口获取有效accessToken.
* 5.wxUser登录成功返回有效Session_key
