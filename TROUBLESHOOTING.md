# 阿里云部署排查指南

## 问题：部署后没有Go服务启动

### 排查步骤

#### 1. 检查容器是否运行

SSH登录到阿里云服务器，执行：

```bash
# 查看所有容器（包括停止的）
docker ps -a

# 查看容器日志
docker logs <容器ID或名称>

# 查看容器详细信息
docker inspect <容器ID或名称>
```

#### 2. 检查容器状态

```bash
# 如果容器已停止，尝试启动
docker start <容器ID或名称>

# 如果容器不断重启，查看日志
docker logs -f <容器ID或名称>
```

#### 3. 检查端口占用

```bash
# 查看端口是否被占用
netstat -tunlp | grep 8888

# 或者使用 ss 命令
ss -tunlp | grep 8888
```

#### 4. 手动运行容器测试

```bash
# 拉取镜像
docker pull crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com/core_ziming/ziming:latest

# 手动运行容器
docker run -d \
  --name hello-hertz-test \
  -p 8888:8888 \
  crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com/core_ziming/ziming:latest

# 查看容器日志
docker logs -f hello-hertz-test

# 测试服务
curl http://localhost:8888/uppercase?text=hello
```

#### 5. 检查阿里云Flow部署配置

在阿里云Flow控制台检查：

1. **部署命令是否正确**
   - 确保使用了正确的端口映射 `-p 8888:8888`
   - 确保容器名称正确
   - 确保使用了 `--restart=always` 参数

2. **部署脚本示例**
   ```bash
   # 停止并删除旧容器
   docker stop hello-hertz-app 2>/dev/null || true
   docker rm hello-hertz-app 2>/dev/null || true
   
   # 拉取最新镜像
   docker pull crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com/core_ziming/ziming:latest
   
   # 运行新容器
   docker run -d \
     --name hello-hertz-app \
     -p 8888:8888 \
     --restart=always \
     crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com/core_ziming/ziming:latest
   ```

#### 6. 检查安全组配置

在阿里云控制台检查：

1. **安全组规则**
   - 确保开放了 8888 端口
   - 入方向规则：允许 TCP 8888 端口

2. **防火墙设置**
   ```bash
   # 检查防火墙状态
   sudo firewall-cmd --state
   
   # 如果防火墙开启，开放端口
   sudo firewall-cmd --zone=public --add-port=8888/tcp --permanent
   sudo firewall-cmd --reload
   ```

#### 7. 检查容器内部

```bash
# 进入容器内部检查
docker exec -it <容器ID或名称> sh

# 检查进程
ps aux | grep main

# 检查端口
netstat -tunlp | grep 8888

# 手动启动服务测试
./main
```

### 常见问题及解决方案

#### 问题1：容器启动后立即退出

**原因**：容器内没有前台进程运行

**解决方案**：确保 Dockerfile 中的 CMD 正确
```dockerfile
CMD ["./main"]
```

#### 问题2：端口无法访问

**原因**：端口映射错误或安全组未开放

**解决方案**：
1. 检查端口映射：`docker port <容器ID>`
2. 开放安全组端口 8888

#### 问题3：镜像拉取失败

**原因**：未登录镜像仓库

**解决方案**：
```bash
# 登录阿里云镜像仓库
docker login --username=你的账号 crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com
```

### 完整部署脚本

创建一个完整的部署脚本 `deploy-to-server.sh`：

```bash
#!/bin/bash

# 配置
IMAGE="crpi-d7yjhyilsfuhe4x4.cn-hangzhou.personal.cr.aliyuncs.com/core_ziming/ziming:latest"
CONTAINER_NAME="hello-hertz-app"
PORT=8888

echo "开始部署..."

# 停止并删除旧容器
echo "停止旧容器..."
docker stop ${CONTAINER_NAME} 2>/dev/null || true
docker rm ${CONTAINER_NAME} 2>/dev/null || true

# 拉取最新镜像
echo "拉取最新镜像..."
docker pull ${IMAGE}

# 运行新容器
echo "启动新容器..."
docker run -d \
  --name ${CONTAINER_NAME} \
  -p ${PORT}:${PORT} \
  --restart=always \
  ${IMAGE}

# 等待容器启动
sleep 3

# 检查容器状态
echo "检查容器状态..."
docker ps | grep ${CONTAINER_NAME}

# 查看日志
echo "容器日志："
docker logs ${CONTAINER_NAME}

# 测试服务
echo "测试服务..."
curl -s "http://localhost:${PORT}/uppercase?text=hello" || echo "服务测试失败"

echo "部署完成！"
```

### 联系支持

如果以上步骤都无法解决问题，请提供以下信息：

1. `docker ps -a` 的输出
2. `docker logs <容器ID>` 的输出
3. 阿里云Flow的部署日志
4. 安全组配置截图
