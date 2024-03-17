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

## 组件安装  

### 安装Consul

```shell
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp  consul consul agent  -dev -client=0.0.0.0
```

访问`http://localhost:8500/`

- 8500是HTTP端口
- 8600是DNS端口

### 安装Nacos

```shell
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=512m -e JVM_XMX=512m -e JVM_XMN=256m -p 8848:8848 -d nacos/nacos-server:latest
```

### 安装ES

```shell
#新建es的config配置文件夹  注意：这里/data/elasticsearch/config位置在可随意
mkdir -p /data/elasticsearch/config
#新建es的data目录         注意：这个文件夹使用来存放es中我们要存储和使用的真实数据
mkdir -p /data/elasticsearch/data
#新建es的plugins目录
mkdir -p /data/elasticsearch/plugins
#给目录设置权限
chmod 777 -R /data/elasticsearch

#写入配置到elasticsearch.yml中， 下面的 > 表示覆盖的方式写入， >>表示追加的方式写入，但是要确保http.host: 0.0.0.0不能被写入多次
echo "http.host: 0.0.0.0" >> /data/elasticsearch/config/elasticsearch.yml

#安装es       注意： -e ES_JAVA_OPTS="-Xms128m -Xmx256m"可根据自己设备存储来配置大小
docker run --name elasticsearch -p 9200:9200 -p 9300:9300 \
    -e "discovery.type=single-node" \
  -e ES_JAVA_OPTS="-Xms128m -Xmx256m" \
  -v /data/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml \
  -v /data/elasticsearch/data:/usr/share/elasticsearch/data \
  -v /data/elasticsearch/plugins:/usr/share/elasticsearch/plugins \
  -d elasticsearch:7.10.1
```

在浏览器中访问：http://localhost:9200  

### 安装kibana  

注意：kibana 的版本需要和 es 的版本一致  

```shell
docker run -d --name kibana -e ELASTICSEARCH_HOSTS="http://192.168.0.104:9200" -p 5601:5601 kibana:7.10.1
```

在浏览器中访问：http://localhost:5601  


## 项目运行  

