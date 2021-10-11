Kubernetes installation at home with kubeadm

Nodes:
-master1
-worker1
-worker2

Steps:
Donwload Ubunto 20.04
Install OS on the servers

Load overlay and br_netfilter kernal modules.
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf 
overlay 
br_netfilter 
EOF



