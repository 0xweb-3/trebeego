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

## 路由
* 固定路由
* 动态路由（带参数）
* 正则路由
* 自动路由
* 注解路由
* RESTful 风格支持


## 版本说明
* v1.0.0 连接mysql数据库并，实现一个简单api接
* v2.0.0 路由管理
  * v2.0.1 自动路由实现
  * v2.0.2 注解路由
  * v2.0.3 RESTful风格路由
  * v2.0.4 实现跨域中间件的使用
* v3.0.0 控制器 (Controllers) 参考笔记
* v4.0.0 参数校验
  * v4.0.1 使用 Beego 内置 Validator
  * 使用第三方库 `go-playground/validator`
* 