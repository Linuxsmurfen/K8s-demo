# Kubernetes installation at home with kubeadm

## Overview

|Hostname|Role|
|---|---|
|k8s-master|master|
|k8s-worker1|worker|
|k8s-worker2|worker|

## Prepare the servers
Download Ubuntu 20.04 LTS "ubuntu-20.04.3-live-server-amd64.iso"   
- 8GB mem
- 4vcpu
- 200GB thin disk
- vlan3

Use entire disk    
no lvm   
Enable sshd   
Create a user: user/...   
   
Set hostname
```
sudo hostnamectl set-hostname k8s-master
sudo hostnamectl set-hostname k8s-worker1
sudo hostnamectl set-hostname k8s-worker2
```

## Install Packages
```
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf 
overlay 
br_netfilter 
EOF

sudo modprobe overlay 
sudo modprobe br_netfilter

cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf 
net.bridge.bridge-nf-call-iptables = 1 
net.ipv4.ip_forward = 1 
net.bridge.bridge-nf-call-ip6tables = 1 
EOF

sudo sysctl --system

sudo apt-get update && sudo apt-get install -y containerd

sudo mkdir -p /etc/containerd
sudo containerd config default | sudo tee /etc/containerd/config.toml

sudo systemctl restart containerd

sudo swapoff -a
sudo sed -i '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab

sudo apt-get update && sudo apt-get install -y apt-transport-https curl

curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -

cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF

sudo apt-get update

sudo apt-get install -y kubelet=1.22.2-00 kubeadm=1.22.2-00 kubectl=1.22.2-00

sudo apt-mark hold kubelet kubeadm kubectl
```


## Initialize the Cluster
This only needs to perform on the control plane node only. (If you have multiple control plane nodes, do the same)
```
sudo kubeadm init --pod-network-cidr 172.20.0.0/16

mkdir -p $HOME/.kube
sudo cp -I /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl version
```

## Install Calico Network Add-On
On the control plane node
```
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

kubectl get pods -n kube-system
```

## Join the worker nodes to the Cluster
```
kubeadm token create --print-join-command

sudo kubeadm join <join command from previous command>

kubectl get nodes
```


Thanks to   
https://medium.com/platformer-blog/building-a-kubernetes-1-20-cluster-with-kubeadm-4b745eb5c697   
https://medium.com/platformer-blog/kubernetes-highly-available-cluster-upgrade-10f709bb357a   





## Install MetalLB Add-On
Used for baremetal loadbalancing, using ip range 192.168.3.240-250

Change to "strictARP: true"
```
kubectl edit configmap -n kube-system kube-proxy
```

Install by manifest
```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.3/manifests/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.10.3/manifests/metallb.yaml
```

Configure
```
kubectl apply -f MetalLB-ConfigMap.yaml -n metallb-system
------------
apiVersion: v1
kind: ConfigMap
metadata:
  namespace: metallb-system
  name: config
data:
  config: |
    address-pools:
    - name: default
      protocol: layer2
      addresses:
      - 192.168.3.240-192.168.3.250
      
```
Thanks to   
https://metallb.universe.tf/installation/   
https://metallb.universe.tf/configuration/   


## Install Helm
   
On the master node install Helm   
```
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh
```

Thanks to   
https://helm.sh/docs/intro/install/
