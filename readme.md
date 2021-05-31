### 环境
  - macos 11.3
  - docker version
```
➜  simpleapp git:(master) docker -v    
Docker version 20.10.6, build 370c289
```
  - docker-compose version 
```
➜  simpleapp git:(master) ✗ docker-compose -v                       
docker-compose version 1.29.1, build c34c88b2
➜  simpleapp git:(master) ✗ 
```
  - minikube version
```
➜  simpleapp git:(master) ✗ minikube version
minikube version: v1.19.0
commit: 15cede53bdc5fe242228853e737333b09d4336b5
```
  - helm 
```
➜  simpleapp git:(master) ✗ helm version                            
version.BuildInfo{Version:"v3.5.3", GitCommit:"041ce5a2c17a58be0fcd5f5e16fb3e7e95fea622", GitTreeState:"dirty", GoVersion:"go1.16"}
```

### 清理 `<none>` 镜像
`docker image prune`

## minikube中 docker-daemon指向本地
    eval $(minikube docker-env)
### 构建镜像 到minikube
    docker compose build 
### 配置helm 
    cd simpleapp && \
    helm install -f values.yaml simpleapp . 
### 检查结果
    - 获取集群ip 
        minikube ip
      配置本地hosts
        vim /etc/hosts
            192.168.64.6 ziggle.io
      请求服务
        curl ziggle.io



