# Kubernetes installation at home with kubeadm

## 1. Overview

|Hostname|Role|IP|
|---|---|---|
|k8s-master|master|192.168.3.100|
|k8s-worker1|worker|192.168.3.101|
|k8s-worker2|worker|192.168.3.102|


## 2. Install and configure the OS

### Install
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

### Configure
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


## 3. Install Kubernetes
### Initialize the cluster
This only needs to perform on the control plane node only. (If you have multiple control plane nodes, do the same)
```
sudo kubeadm init --pod-network-cidr 172.20.0.0/16

mkdir -p $HOME/.kube
sudo cp -I /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

kubectl version
```

### Install Calico - CNI
On the control plane node
```
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

kubectl get pods -n kube-system
```

### Join the worker nodes to the Cluster
```
kubeadm token create --print-join-command

sudo kubeadm join <join command from previous command>

kubectl get nodes
```


Thanks to   
https://medium.com/platformer-blog/building-a-kubernetes-1-20-cluster-with-kubeadm-4b745eb5c697   
https://medium.com/platformer-blog/kubernetes-highly-available-cluster-upgrade-10f709bb357a   





## 4. Install a baremetal loadbalancer
Using Metallb with ip range 192.168.3.240-250   

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


## 5. Install Helm
   
On the master node install Helm   
```
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh
```

Thanks to   
https://helm.sh/docs/intro/install/



## 6. Install ingress

Add nginx repository to helm
```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
```

Install nginx ingress controller
```
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

Thanks to   
https://www.debontonline.com/2020/10/kubernetes-part-5-install-and-configure.html
https://kubernetes.io/docs/concepts/services-networking/ingress/



## 7. Setup dynamic persistant storage - CSI
NFS to my Synology NAS with "nfs-subdir-external-provisioner"   
Install NFS commons on the worker nodes
```
sudo apt install nfs-common
```

Verify that the share is available
```
showmount -e 192.168.1.200
```

Add the NFS driver
```
$ helm repo add nfs-subdir-external-provisioner https://kubernetes-sigs.github.io/nfs-subdir-external-provisioner/
$ helm install nfs-subdir-external-provisioner nfs-subdir-external-provisioner/nfs-subdir-external-provisioner  \
  --set nfs.server=192.168.1.200  \
  --set nfs.path=/volume1/NFS
```
Thanks to   
https://www.debontonline.com/2020/11/kubernetes-part-11-how-to-configure.html
https://github.com/kubernetes-sigs/nfs-subdir-external-provisioner/blob/master/charts/nfs-subdir-external-provisioner/README.md


## 8. Ceph storage - CSI
See the "addRockCeph" document


## 9. KubeView
Nice GUI for showing how everything is connected.

```
git clone https://github.com/benc-uk/kubeview
cd kubeview/charts/
cp example-values.yaml myvalues.yaml 

edit myvalues.yaml  (set to false) 

helm install kubeview kubeview -f myvalues.yaml --namespace kubeview --create-namespace
(helm uninstall kubeview kubeview  --namespace kubeview)
```

Setup the ingress rules
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-kubeview
  namespace: kubeview
  annotations:
spec:
  ingressClassName: nginx
  rules:
    - host: "kubeview.k8s.home.ip"      
      http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: kubeview
              port:
                number: 80
```


Thanks to   
https://github.com/benc-uk/kubeview


## 10. Gitea

<< Need to document this!  >>


Setup the ingress rules
```
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-gitea
  namespace: gitea
  annotations:
spec:
  ingressClassName: nginx
  rules:
    - host: "gitea.k8s.home.ip"      
      http:
        paths:
        - path: /
          pathType: Prefix
          backend:
            service:
              name: gitea-service
              port:
                number: 3000
```


## 11. Add DNS entries to DNS
Login to the EdgeRouter and add the entries. 
(Was not able to add any wildcard DNS entries)

```
$ configure
# set system static-host-mapping host-name kubeview.k8s.home.ip inet 192.168.x.x
# set system static-host-mapping host-name gitea.k8s.home.ip inet 192.168.x.x
# commit
# save
# exit
```

Thanks to   
https://gist.github.com/plembo/6bb4491ebbfbce049c7efce0634d57f0




## 12. PiHole
See the "pi-hole.yaml" file   



## Things to do...
1. Flux
2. ~~Perssistant volumes with (OpenEBS, Longhorn...)~~
3. Local persistant volumes
4. Backup
5. Dashboard
6. Keycloak

## Applications to add...
1. ~~Gitea~~
2. ~~PiHole~~
3. MeshController
4. Home Assistant
   1. Deconz
   2. MariaDB
5. Minio
6. Prometheus
7. Grafana
8. Photoprism
9. Wireguard
10. ~~KubeView~~


## Roadmap
For the next iteration:
1. PXE boot the servers
2. Ansible configuration
  - OS settings
  - kubeadm
  - calico
  - flux

3. Flux syncs all applications
-  storage setup
-  backup setup
-  apps..


