kubectl apply -f namespace.yaml
kubectl apply -f 2_mcm-app/1-app.yaml
kubectl apply -f 2_mcm-app/2-app-channel.yaml
kubectl apply -f 2_mcm-app/3-app-subscription.yaml
kubectl apply -f 2_mcm-app/4-app-deployables.yaml
