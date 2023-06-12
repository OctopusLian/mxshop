# 慕学生鲜  

## 技术栈  

- redis  
- mysql  
- elasticsearch  
- rocketmq  
- grpc  
- gin  
- viper  
- jaeger  
- nacos  
- sentinel  
- yapi  
- 分布式锁  
- 分布式事务  
- gorm  
- 负载均衡  
- 服务注册  
- 服务发现oss  
- 幂等机制  
- kong  
- jenkins  
- docker  
- kubernetes  

## 运行  

### 启动 Consul  

```
$ consul agent -dev -client=0.0.0.0
$ dig @物理IP -p 8600 consul.service.consul SRV
```

访问`http://localhost:8500/`  

- 8500是HTTP端口  
- 8600是DNS端口  

