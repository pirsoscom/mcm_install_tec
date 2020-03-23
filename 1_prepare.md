
# Preparation of the Nodes

## Set correct vm.max_map_count

### Login to Bation Host
ssh root@10.99.96.220


### Update the Nodes
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


