kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/crds/k8sdc_cr.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/agentoperator.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/icam-reloader.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/operator.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/role.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/role_binding.yaml
kubectl delete -f ./tools/apm/deploy/service_account.yaml
kubectl delete -f ./tools/apm/deploy/crds/k8sdc_crd.yaml

oc delete secret dc-secret -n multicluster-endpoint
oc delete secret ibm-agent-https-secret -n multicluster-endpoint

oc delete clusterrolebinding icamklust-binding -n multicluster-endpoint
oc delete clusterrolebinding icamklust-binding_default -n multicluster-endpoint
