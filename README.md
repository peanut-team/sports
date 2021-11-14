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
docker-compose build --no-cache
## 运行
docker-compose up -d

## or make
make mysql.up
make mysql.down
```

### 运行
```bash
go run main.go

# or run make
make run
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
