# Install Scripts for IBM Cloud Pack for Multicloud Management on TEC


## Preparation of the Nodes

### Set correct vm.max_map_count

1. Login to Bation Host

```shell
ssh root@10.99.96.220
```

2. Update the Nodes

```shell
ssh core@192.168.1.101 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.102 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.103 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.104 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.105 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.106 sudo sysctl -w vm.max_map_count=262144
ssh core@192.168.1.101 echo “vm.max_map_count=262144”>>/etc/sysctl.conf
ssh core@192.168.1.102 echo “vm.max_map_count=262144">>/etc/sysctl.conf
ssh core@192.168.1.103 echo “vm.max_map_count=262144”>>/etc/sysctl.conf
ssh core@192.168.1.104 echo “vm.max_map_count=262144">>/etc/sysctl.conf
ssh core@192.168.1.105 echo “vm.max_map_count=262144”>>/etc/sysctl.conf
ssh core@192.168.1.106 echo “vm.max_map_count=262144">>/etc/sysctl.conf
```



## Install MCM

```
./2_install_mcm_online.sh -t YOUR-KEY -d /tmp/mcm-install -p Your-Pa$$w0rd
```

## Install CAM

```
./3_install_cam_online.sh -t YOUR-KEY -d /tmp/mcm-install -p Your-Pa$$w0rd
```

## Install ICAM

```
./4_install_apm_online.sh -t YOUR-KEY -d /tmp/mcm-install -p Your-Pa$$w0rd

```