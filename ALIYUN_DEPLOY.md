# 阿里云容器镜像服务部署指南

## 问题说明

阿里云容器镜像构建服务默认无法访问Docker Hub，需要使用国内镜像源。

## 解决方案

### 方案一：使用阿里云容器镜像服务的海外构建功能

1. 登录阿里云容器镜像服务控制台
2. 创建命名空间
3. 创建镜像仓库
4. 在构建配置中，选择"海外机器构建"选项
5. 这样可以直接从Docker Hub拉取镜像

### 方案二：本地构建后推送（推荐）

#### 步骤1：配置Docker镜像加速

在本地Docker Desktop中配置镜像加速器（参考之前的说明）

#### 步骤2：本地构建镜像

```bash
# 构建镜像
docker build -t hello-hertz:latest .

# 测试镜像
docker run -d -p 8888:8888 hello-hertz:latest
curl http://localhost:8888/uppercase?text=hello
```

#### 步骤3：推送到阿里云镜像仓库

```bash
# 登录阿里云镜像仓库
docker login --username=你的阿里云账号 registry.cn-hangzhou.aliyuncs.com

# 标记镜像
docker tag hello-hertz:latest registry.cn-hangzhou.aliyuncs.com/你的命名空间/hello-hertz:latest

# 推送镜像
docker push registry.cn-hangzhou.aliyuncs.com/你的命名空间/hello-hertz:latest
```

#### 步骤4：在服务器上拉取并运行

```bash
# 登录镜像仓库
docker login --username=你的阿里云账号 registry.cn-hangzhou.aliyuncs.com

# 拉取镜像
docker pull registry.cn-hangzhou.aliyuncs.com/你的命名空间/hello-hertz:latest

# 运行容器
docker run -d \
  --name hello-hertz-app \
  -p 8888:8888 \
  --restart=always \
  registry.cn-hangzhou.aliyuncs.com/你的命名空间/hello-hertz:latest
```

### 方案三：使用GitHub Actions自动构建

创建 `.github/workflows/docker-build.yml` 文件：

```yaml
name: Build and Push Docker Image

on:
  push:
    branches: [ main ]

jobs:
  build:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v3
    
    - name: Set up Docker Buildx
      uses: docker/setup-buildx-action@v2
    
    - name: Login to Aliyun Container Registry
      uses: docker/login-action@v2
      with:
        registry: registry.cn-hangzhou.aliyuncs.com
        username: ${{ secrets.ALIYUN_USERNAME }}
        password: ${{ secrets.ALIYUN_PASSWORD }}
    
    - name: Build and push
      uses: docker/build-push-action@v4
      with:
        context: .
        push: true
        tags: registry.cn-hangzhou.aliyuncs.com/你的命名空间/hello-hertz:latest
```

### 方案四：在服务器上直接构建

```bash
# 在阿里云服务器上
git clone https://github.com/yanshuzh/hello_ziming.git
cd hello_ziming

# 配置Docker镜像加速
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://docker.m.daocloud.io"
  ]
}
EOF

# 重启Docker
sudo systemctl daemon-reload
sudo systemctl restart docker

# 构建镜像
docker build -t hello-hertz:latest .

# 运行容器
docker run -d --name hello-hertz-app -p 8888:8888 --restart=always hello-hertz:latest
```

## 镜像仓库地址

根据你的地域选择合适的镜像仓库地址：

- 华东1（杭州）：`registry.cn-hangzhou.aliyuncs.com`
- 华东2（上海）：`registry.cn-shanghai.aliyuncs.com`
- 华北2（北京）：`registry.cn-beijing.aliyuncs.com`
- 华南1（深圳）：`registry.cn-shenzhen.aliyuncs.com`

## 注意事项

1. 确保阿里云账号已开通容器镜像服务
2. 需要创建命名空间和镜像仓库
3. 推送镜像前需要先登录
4. 建议使用访问凭证而不是账号密码
