
kubectl delete LimitRange -n default mem-limit-range

kubectl delete -n default -f policy-deployment-nginx.yaml
kubectl delete -n default -f policy-dev-endpoint.yaml
kubectl delete -n default -f policy-namespace.yaml
kubectl delete -n default -f policy-pod-security.yaml
kubectl delete -n default -f policy-prod-network.yaml
kubectl delete -n default -f policy-prod.yaml
kubectl delete -n default -f policy-rhocp-sdn.yaml
kubectl delete -n default -f policy-mutation.yaml
kubectl delete -n default -f policy-cis.yaml


kubectl get pp --all-namespaces