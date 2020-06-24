
# Install Script for IBM Cloud Pack for Multicloud Management on OCP 4.3 in TEC




## Preparing for installation

### 1) Prerequisites

You need the following installed and running:

* Docker
* kubectl (with kubernetes context set to the Cluster where you want to install)
* helm (the helm2 command line tool)  - **Helm 3 won't work!**
* cloudctl


### 2) Update Nodes

From the Bastion Host run:

```bash
for i in $(grep -E '(master|worker)' /etc/hosts | awk '{print $1}');do     ssh core@$i 'hostname;sudo sysctl -w vm.max_map_count=262144'; done
```

### 3) Adapt config file  (0_config.sh)

Please select the components that you want to install.

```bash
# ---------------------------------------------------------------------------------------------------------------------------------------------------"
# ---------------------------------------------------------------------------------------------------------------------------------------------------"
# Adapt Values
# ---------------------------------------------------------------------------------------------------------------------------------------------------"
# ---------------------------------------------------------------------------------------------------------------------------------------------------"
export MCM_VERSION=3.2.8
export CAM_VERSION=4.2.1
export APM_VERSION=1.7.1



# Install MultiCloud Manager (MCM) (only for testing, do not disable this as the other components depend on this)?
export INSTALL_MCM=true
# Register MCM CIS Controller ?
export INSTALL_MCMCIS=true
# Register MCM Mutation Advisor ?
export INSTALL_MCMMUT=true
# Register MCM Notary ?
export INSTALL_MCMNOT=false


# Register MCM Hub Cluster ?
export INSTALL_MCMREG=true
# Install Terraform & Service Automation Module (CAM) ?
export INSTALL_CAM=true
# Install Monitoring Module (APM) ?
export INSTALL_APM=true
# Install and integrate OPENLDAP ?
export INSTALL_LDAP=true
# Install Ansible Tower ?
export INSTALL_ANSIBLE=false
# Install ManageIQ Instance ?
export INSTALL_MIQ=false
# Install Demo Apps and Policies ?
export INSTALL_DEMO=true





```


### 4) Adapt hosts file

You have to adapt your hosts file

```
<BASTION_IP> 	console-openshift-console.apps.ocp43.tec.uk.ibm.com oauth-openshift.apps.ocp43.tec.uk.ibm.com default-route-openshift-image-registry.apps.ocp43.tec.uk.ibm.com api.ocp43.tec.uk.ibm.com icp-console.apps.ocp43.tec.uk.ibm.com cam.apps.ocp43.tec.uk.ibm.com grafana-openshift-monitoring.apps.ocp43.tec.uk.ibm.com openldap-default.apps.ocp43.tec.uk.ibm.com openldap-admin-default.apps.ocp43.tec.uk.ibm.com icp-proxy.apps.ocp43.tec.uk.ibm.com multicloud-console.apps.ocp43.tec.uk.ibm.com kubetoy-default.apps.ocp43.tec.uk.ibm.com
```

When installing with additional clusters (you probably won't need this) 

```
<OKD1_IP>	    okd1.tec.hur.cdn console.apps.okd1.tec.hur.cdn console.okd1.tec.hur.cdn grpc-web-route-grpcdemo-app.apps.okd1.tec.hur.cdn modresort-app-web-route-modresort-app.apps.okd1.tec.hur.cdn docker-registry-default.apps.okd1.tec.hur.cdn
<OKD2_IP>   	okd2.tec.hur.cdn console.apps.okd2.tec.hur.cdn console.okd2.tec.hur.cdn grpc-web-route-grpcdemo-app.apps.okd2.tec.hur.cdn modresort-app-web-route-modresort-app.apps.okd2.tec.hur.cdn docker-registry-default.apps.okd2.tec.hur.cdn
```





### 5) Adapt Docker Daemon configuration

```bash
"default-route-openshift-image-registry.apps.ocp43.tec.uk.ibm.com",
```

When installing with additional clusters (you probably won't need this) 

```bash
"docker-registry-default.apps.okd1.tec.hur.cdn",
"docker-registry-default.apps.okd2.tec.hur.cdn"
```

## Installing


> USAGE 
> 
> 1\_install\_all.sh 
> 
> **-t** \<REGISTRY\_TOKEN\> 
> 
> **-x** \<OCP\_CONSOLE\_PREFIX\> 
> 
> **-s** \<STORAGE\_CLASS\> 
> 
> **-p** \<MCM\_PASSWORD\> 
> 
> **-l** \<LDAP\_ADMIN\_PASSWORD\> 
> 
> [**-h** \<CLUSTER\_NAME\>] 
> 
> [**-d** \<TEMP\_DIRECTORY\>]


Example:

```bash
./1_install_all.sh -t MY_TOKEN -x console-openshift-console -s nfs-client -p passw0rd -l passw0rd -d /tmp/mcm-install

```


___


## Troubleshooting

### Separate installations

```bash
./2_install_mcm.sh -t MY_TOKEN -d /tmp/mcm-install -p passw0rd

./3_install_cam.sh -t MY_TOKEN -d /tmp/mcm-install -x console-openshift-console -p passw0rd

./4_install_apm.sh -t MY_TOKEN -d /tmp/mcm-install -x console-openshift-console -p passw0rd

./5_install_ansible.sh  -d /tmp/mcm-install -x console-openshift-console -p passw0rd

./6_integrate_cloudforms.sh -d /tmp/mcm-install -x console-openshift-console -p passw0rd -i <CF_IP>

./7_register_cluster_u.sh -d /tmp/mcm-install -x $OCP_CONSOLE_PREFIX -p $MCM_PWD  -n mcm-hub -h "https://icp-console.<CLUSTER_NAME>"

./8_install_ldap.sh -d /tmp/mcm-install -x console-openshift-console -p passw0rd

./9_register_k8_monitor.sh -d niklaushirt -n mcm-hub -f /Users/nhirt/ibm-cloud-apm-dc-configpack_CP4MCM002.tar

./11_demo_menu.sh -d /tmp/mcm-install

kubectl apply -f tools/apm/kubetoy_all_in_one.yaml -n default

```