#kubectl delete -f namespace.yaml
kubectl delete -f 2_mcm-app/1-app.yaml
kubectl delete -f 2_mcm-app/2-app-channel.yaml
kubectl delete -f 2_mcm-app/3-app-subscription.yaml
kubectl delete -f 2_mcm-app/4-app-deployables.yaml