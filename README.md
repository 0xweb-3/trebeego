# Beego上手

## 安装 bee 工具

```bash
go install github.com/beego/bee/v2@latest
```

⚠️ 注意：

- `bee` 会被安装到 `$GOPATH/bin` 目录下（默认是 `~/go/bin`）。
- 确保把 `~/go/bin` 加入到环境变量 PATH：

```bash
echo 'export PATH=$PATH:~/go/bin' >> ~/.zshrc
source ~/.zshrc
```

## 创建项目

```bash
bee new trybeego

cd trybeego
go run main.go
```

默认会监听在 `http://localhost:8080`。

## 目录结构说明

- `main.go`：入口文件
- `controllers/`：放业务逻辑控制器
- `models/`：放数据模型
- `routers/`：路由定义
- `conf/app.conf`：配置文件（端口、数据库等）

## 数据库验证
```shell
docker exec -it beego-mysql mysql -u beego -pbeego123
```