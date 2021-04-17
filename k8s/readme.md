kubectl cluster-info 

minikube addons enable ingress

kubectl get ingress


# 创建ingress失败

kubectl get validatingwebhookconfigurations
kubectl delete -A ValidatingWebhookConfiguration ingress-nginx-admission

## 
kubectl delete clusterrolebinding kubernetes-dashboard


kubectl proxy
```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
EOF
```

kubectl -n kubernetes-dashboard get secret $(kubectl -n kubernetes-dashboard get sa/admin-user -o jsonpath="{.secrets[0].name}") -o go-template="{{.data.token | base64decode}}"
