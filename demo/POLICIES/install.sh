
kubectl delete LimitRange -n default mem-limit-range

kubectl get LimitRange --all-namespaces



kubectl apply -n default -f policy-deployment-nginx.yaml
kubectl apply -n default -f policy-dev-endpoint.yaml
kubectl apply -n default -f policy-namespace.yaml
kubectl apply -n default -f policy-pod-security.yaml
kubectl apply -n default -f policy-prod-network.yaml
kubectl apply -n default -f policy-prod.yaml
kubectl apply -n default -f policy-rhocp-sdn.yaml
kubectl apply -n default -f policy-mutation.yaml
kubectl apply -n default -f policy-cis.yaml

kubectl get pp --all-namespaces