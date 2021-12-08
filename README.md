# Sprots

## 开发
### Swagger
```bash
# 安装 swag 工具
## linux or macos
go get -u github.com/swaggo/swag/cmd/swag
mv $GOPATH/bin/swag /usr/local/bin

# 生成或者更新 swag 文档
go get -u github.com/swaggo/swag/cmd/swag
# 寻找 github.com/swaggo/swag@latest 目录，$SWAGPATH
mv $SWAGPATH/cmd/swag/swag.exe $GOPATH
```
参考：
[]

### 基础设施
#### 数据库
```bash
# 数据库
## 生成镜像
docker-compose -f docker-compose-db.yaml build mariadb --no-cache
## 运行
docker-compose -f docker-compose-db.yaml up -d mariadb

## 停止
docker-compose -f docker-compose-db.yaml down mariadb
```

#### emq
```bash
docker-compose -f docker-compose-emq.yaml -p my_emqx up -d
# 查看集群状态
docker exec -it my_emqx_emqx1_1 sh -c "emqx_ctl cluster status"

# 停止
docker-compose -f docker-compose-emq.yaml -p my_emqx down
```
#### redis
```bash
# 运行
docker-compose -f docker-compose-redis.yaml up -d

# 停止
docker-compose -f docker-compose-redis.yaml down

#redis的认证密码一般配置在配置文件的requirepass字段。如果不使用配置文件，可以使用 command: redis-server --requirepass yourpass 配置认证密码；
```

#### rabbitmq
```bash
## 运行
docker-compose -f docker-compose-rabbitmq.yaml up -d

## 停止
docker-compose -f docker-compose-rabbitmq.yaml down

## 启动监控插件
docker-compose -f docker-compose-rabbitmq.yaml ps
docker exec -it 容器Name /bin/bash
rabbitmq-plugins enable rabbitmq_management # 在容器中执行
```

### 运行
```bash
go run main.go

# or run make
make run

#build docker 
make build 

# docker-composer run
docker-compose -f docker-compose-api-server.yaml up -d
docker-compose -f docker-compose-api-server.yaml down
```

### 单元测试
```bash
make test
```

### other
```bash
make lint
make format
```


## 基本方案
IOT --> Mqtt --> Redis --> API Server  --> Fronted
          |
          |
          DB

