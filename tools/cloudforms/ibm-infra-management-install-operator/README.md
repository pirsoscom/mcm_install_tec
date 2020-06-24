<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [IBM Infrastructure Management Install Operator](#ibm-infrastructure-management-install-operator)
  - [PodSecurityPolicy Requirements](#podsecuritypolicy-requirements)
  - [SecurityContextConstraints Requirements](#securitycontextconstraints-requirements)
  - [GVK](#gvk)
    - [Group](#group)
    - [Version](#version)
    - [Kind](#kind)
    - [Sample CR](#sample-cr)
  - [Run Operator](#run-operator)
    - [Run Inside Cluster](#run-inside-cluster)
    - [Set up a default storageclass](#set-up-a-default-storageclass)
    - [Setup RBAC and deploy the operator](#setup-rbac-and-deploy-the-operator)
    - [Create the CR to deploy IBM Infrastructure Management](#create-the-cr-to-deploy-ibm-infrastructure-management)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# IBM Infrastructure Management Install Operator

This operator manages the install of IBM Infrastructure Management on a OCP4 cluster.

## PodSecurityPolicy Requirements

## SecurityContextConstraints Requirements

## GVK

### Group

```
iminstall.management.ibm.com
```

### Version

```
v1alpha1
```

### Kind

```
IMInstall
```

### Sample CR

```yaml
apiVersion: iminstall.management.ibm.com/v1alpha1
kind: IMInstall
metadata:
  name: example-iminstall
spec:
  applicationDomain: "iminstall.apps.gyliu-ocp-1.os.fyre.ibm.com"
```

## Run Operator

Deploy the IBM Infrastructure Management CRD

```bash
$ oc create -f deploy/crds/infra.management.ibm.com_iminstalls_crd.yaml
```
### Run Inside Cluster

1. Build and push the operator image (**Use your own dockerhub for test**):

```bash
$ operator-sdk build quay.io/example/iminstall-operator:latest
$ docker push quay.io/example/iminstall-operator:latest
```

2. Update the operator deployment with the new image:

```bash
$ sed -i 's|quay.io/multicloudlab/iminstall-operator:0.0.1|quay.io/example/iminstall-operator:latest|g' deploy/operator.yaml
```

3. Replace the namespace in operator resources by the running the following commands:

```sh
$ export IMINSTALL_NAMESPACE=<YOUR_EXPECTED_NAMESPACE>
$ sed -i 's|iminstall-namespace|'"${IMINSTALL_NAMESPACE}"'|g' deploy/namespace.yaml
$ sed -i 's|REPLACE_NAMESPACE|'"${IMINSTALL_NAMESPACE}"'|g' deploy/role_binding.yaml
```

### Set up a default storageclass

If you are using OCP, run the following command to set up the default storageclass.

```bash
oc patch storageclass rook-ceph-cephfs-internal -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

### Setup RBAC and deploy the operator

```bash
$ oc create -f deploy/namespace.yaml
$ oc create -f deploy/role.yaml -n ${IMINSTALL_NAMESPACE}
$ oc create -f deploy/role_binding.yaml
$ oc create -f deploy/service_account.yaml -n ${IMINSTALL_NAMESPACE}
$ oc create -f deploy/operator.yaml -n ${IMINSTALL_NAMESPACE}
```

### Create the CR to deploy IBM Infrastructure Management

```bash
$ oc create -f deploy/crds/infra.management.ibm.com_v1alpha1_iminstall_cr.yaml
```

**NOTE**: If your orchestor pod keeps crashing, please see https://github.com/ManageIQ/manageiq-pods/issues/487 for detail, and use following CR instead.

```yaml
kind: IMInstall
metadata:
  name: example-iminstall
spec:
  applicationDomain: "iminstall.apps.gyliu-ocp-1.os.fyre.ibm.com"
  orchestratorInitialDelay: "2400"
  license:
    accept: true
```

**IBM Infrastructure Management Instance Example**

> The domain here will work for a Code Ready Containers cluster. **Change it to one that will work for your environment**.

> Additional parameters are available and documented in the Custom Resource Definition

```yaml
apiVersion: iminstall.management.ibm.com/v1alpha1
kind: IMInstall
metadata:
  name: example-iminstall
spec:
  applicationDomain: "iminstall.apps.gyliu-ocp-1.os.fyre.ibm.com"
  license:
    accept: true
```
