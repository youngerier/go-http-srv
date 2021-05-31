## 清理 <none>镜像
docker image prune

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



