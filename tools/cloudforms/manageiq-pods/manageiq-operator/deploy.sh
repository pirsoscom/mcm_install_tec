oc create namespace manageiq


oc create -n manageiq -f deploy/crds/manageiq.org_manageiqs_crd.yaml
oc create -n manageiq -f deploy/role.yaml
oc create -n manageiq -f deploy/role_binding.yaml
oc create -n manageiq -f deploy/service_account.yaml
oc create -n manageiq -f deploy/operator.yaml
oc create -n manageiq -f deploy/kafka-secret.yaml

kubectl create secret generic -n manageiq miq-configuration --from-file=./manageiq-external-auth-openidc.conf



oc create -n manageiq -f deploy/crds/manageiq.org_v1alpha1_manageiq_cr.yaml


oc delete -n manageiq -f deploy/crds/manageiq.org_v1alpha1_manageiq_cr.yaml
oc delete -n manageiq -f deploy/crds/manageiq.org_manageiqs_crd.yaml
oc delete -n manageiq -f deploy/role.yaml
oc delete -n manageiq -f deploy/role_binding.yaml
oc delete -n manageiq -f deploy/service_account.yaml
oc delete -n manageiq -f deploy/operator.yaml


cat >> deploy/kafka-secret.yaml <<EOL
kind: Secret
apiVersion: v1
metadata:
  name: kafka-secrets
  namespace: manageiq
  selfLink: /api/v1/namespaces/manageiq/secrets/kafka-secrets
  uid: 7c1b541f-87d5-467f-bbe1-938f9e84c7bc
  resourceVersion: '20801086'
  creationTimestamp: '2020-06-12T06:42:35Z'
  labels:
    app: manageiq
data:
  hostname: a2Fma2E=
  password: NmMwYjQ0NjU5MjYwMjgzZA==
  username: cm9vdA==
type: Opaque
EOL





      echo " ${GREEN}Register OAUTH Client${NC}"
          cloudctl iam oauth-client-delete $CLIENT_ID
          cloudctl iam oauth-client-register -f $INSTALL_PATH/registration.json
        echo " ${GREEN}    OK${NC}"


        echo " ${GREEN}Get Cluster CA cert${NC}"
          kubectl get secret -n kube-public ibmcloud-cluster-ca-cert -o jsonpath='{.data.ca\.crt}' | base64 --decode | sed 's/CERTIFICATE/TRUSTED CERTIFICATE/' > $INSTALL_PATH/ibm-mcm-ca.crt
        echo " ${GREEN}    OK${NC}"

        echo " ${GREEN}Upload files to CF Instance${NC}"
        echo " ${ORANGE}   Enter your admin password for the CF Server if asked for${NC}"
          scp $INSTALL_PATH/ibm-mcm-ca.crt $INSTALL_PATH/manageiq-external-auth-openidc.conf $SCRIPT_PATH/tools/cloudforms/register_host.sh root@$CF_IP:/root
        echo " ${GREEN}    OK${NC}"


        echo " ${GREEN}Register on CF Instance${NC}"
        echo " ${ORANGE}   Enter your admin password for the CF Server if asked for${NC}"
          ssh root@$CF_IP /root/register_host.sh
        echo " ${GREEN}    OK${NC}"