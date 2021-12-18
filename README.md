# simpleRouter
自制路由器,目标是实现一个go版本的spring-cloud-gateway



## 技术栈说明

gin - 基本web框架

redis - 缓存

jwt - 解析token

mysql - 存储

ratelimit - 限流

nacos - 注册中心及配置中心



#### TODO

- [x] db路由配置
- [x] yaml路由配置
- [x] router匹配 - 根据uri
- [x] Filter初始化
- [x] 转发
- [x] 缓存
- [ ] jwt
- [ ] 鉴权
- [ ] 黑白名单统计
- [ ] trace
- [ ] 单机限流
- [ ] 分布式限流
- [ ] 注册中心接入
  - [x] nacos  -- 配置使用yaml配置
- [ ] 灰度
