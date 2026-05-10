# Hello Hertz Service

一个使用 Hertz 框架的 Go Web 服务，提供小写字母转大写的功能。

## 功能特性

- ✅ GET 接口：通过 URL 参数转换文本
- ✅ POST 接口：通过 JSON body 转换文本
- ✅ Docker 支持：提供完整的 Docker 部署方案
- ✅ 轻量级：使用 Alpine 基础镜像，镜像大小约 10MB

## 快速开始

### 本地运行

```bash
# 安装依赖
go mod tidy

# 运行服务
go run main.go
```

服务将在 `http://localhost:8888` 启动

### Docker 运行

```bash
# 构建镜像
docker build -t hello-hertz:latest .

# 运行容器
docker run -d --name hello-hertz-app -p 8888:8888 hello-hertz:latest
```

或使用部署脚本：

```bash
./deploy.sh
```

## API 接口

### GET /uppercase

通过 URL 参数转换文本为小写

**请求示例：**
```bash
curl "http://localhost:8888/uppercase?text=hello%20world"
```

**响应示例：**
```json
{
  "original": "hello world",
  "result": "HELLO WORLD"
}
```

### POST /uppercase

通过 JSON body 转换文本为大写

**请求示例：**
```bash
curl -X POST http://localhost:8888/uppercase \
  -H "Content-Type: application/json" \
  -d '{"text":"hello hertz"}'
```

**响应示例：**
```json
{
  "original": "hello hertz",
  "result": "HELLO HERTZ"
}
```

## 项目结构

```
hello/
├── main.go          # 主程序文件
├── go.mod           # Go 模块文件
├── go.sum           # 依赖校验文件
├── Dockerfile       # Docker 构建文件
├── .dockerignore    # Docker 忽略文件
├── .gitignore       # Git 忽略文件
├── deploy.sh        # 部署脚本
└── README.md        # 项目说明文档
```

## 技术栈

- **框架**: [Hertz](https://github.com/cloudwego/hertz) - CloudWeGo 出品的 HTTP 框架
- **语言**: Go 1.21+
- **容器化**: Docker

## 部署到阿里云

详细的阿里云部署说明请参考项目文档。

## License

MIT
