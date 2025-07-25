# 2. 快速开始

本章节将指导您如何快速部署和使用A.I.G。

## 一键安装

### 一键Docker部署


安装git / docker 等基础组件

### 从Docker Hub安装

```bash
xxx
```

### 从Github下载安装最新版

```bash
git clone xxx
cd xxx
docker-compose -f docker-compose.images.yml up -d
```

安装完成后，您可以通过浏览器访问 `http://localhost:8088` 来使用A.I.G的Web界面。

### 从源码编译Docker

如果您希望从最新的源码编译docker镜像。

```bash
docker-compose up -d
```

## 一键配置

### 一键Docker部署

通过Web界面的“模型管理”功能可配置A.I.G使用的API KEY

通过配置文件xxx配置。



### 服务架构

```
┌─────────────────┐    ┌─────────────────┐
│   Webserver     │    │     Agent       │
│   (Go + Web)    │◄──►│ (Go + Python)   │
│   Port: 8088    │    │                 │
└─────────────────┘    └─────────────────┘
         │                       │
         └───────┬───────────────┘
                 │
    ┌─────────────────────────┐
    │   Shared Data Volume    │
    │      (data/ dir)        │
    └─────────────────────────┘
```



### 常见问题

1. **端口冲突**
   ```bash
   # 修改webserver端口映射
   ports:
     - "8080:8088"  # 使用8080端口
   ```

2. **权限问题**
   ```bash
   # 检查数据目录权限
   sudo chown -R $USER:$USER ./data
   ```

3. **服务启动失败**
   ```bash
   # 查看详细日志
   docker-compose logs webserver
   docker-compose logs agent
   ```

4. **停止服务**
    ```bash
    # 停止服务
    docker-compose down

    # 停止服务并删除数据卷（谨慎使用）
    docker-compose down -v
    ```

5. **模型API错误**

6. **加载指纹失败**


## 更新升级

```bash
# 拉取最新代码
git pull

# 重新构建并启动
docker-compose up -d --build

# 清理旧镜像
docker image prune -f
``` 



