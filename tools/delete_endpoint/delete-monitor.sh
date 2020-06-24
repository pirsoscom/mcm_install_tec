kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/crds/k8sdc_cr.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/agentoperator.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/icam-reloader.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/operator.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/role.yaml
kubectl delete -n multicluster-endpoint -f ./tools/apm/deploy/role_binding.yaml
kubectl delete -f ./tools/apm/deploy/service_account.yaml
oc delete secret dc-secret -n multicluster-endpoint
oc delete secret ibm-agent-https-secret -n multicluster-endpoint
kubectl patch k8sdcs.ibmcloudappmgmt.com -p '{"metadata":{"finalizers":[]}}' --type=merge k8sdc-cr -n multicluster-endpoint
kubectl delete -f ./tools/apm/deploy/crds/k8sdc_crd.yaml

