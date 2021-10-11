Kubernetes installation at home with kubeadm

Nodes:
-k8s-master
-k8s-worker0
-k8s-worker1

##Steps:
###OS install
Download Ubuntu 20.04 LTS .iso
Create VM with
    • 8GB mem
    • 4vcpu
    • 200GB thin disk
    • vlan3

Entire disk, no lvm
Enable sshd
User: user/...



###Prepare servers
sudo apt install -y kubelet kubeadm kubectl


###Install k8s
sudo kubeadm init --pod-network-cidr 172.20.0.0/16



Thanks to
https://medium.com/platformer-blog/building-a-kubernetes-1-20-cluster-with-kubeadm-4b745eb5c697
