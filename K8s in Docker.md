## K8s in Docker

1.	Setup a docker server with 8GB mem, 16GB disk

2.	Login to the server and verify that ’docker ps’ works   
```
docker ps
```

3.	Start ’Portainer’ to get some visibility   
```
docker volume create portainer_data
docker run -d -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce
```

4.	Download ”Kind”
```
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.9.0/kind-linux-amd64
chmod +x ./kind
```

5.	Download ”kubectl”
```
curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x ./kubectl
```

6.	Create K8s clusters
- Create a minimal cluster with name "cluster1"
```
kind create cluster --name cluster1
```

- Create a 3 master + 3 worker + 1 loadbalancer cluster with the name "cluster2"
Create a configfile ’3m3w.yaml’, change IP to match the server IP.

```
# A cluster with 3 control-plane nodes and 3 workers
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: 192.168.x.x
  apiServerPort: 10001
nodes:
- role: control-plane
- role: control-plane
- role: control-plane
- role: worker
- role: worker
- role: worker
```

- Create cluster
```
kind create cluster --config 3m3w.yaml --name cluster2
```	

7.	Access the K8s cluster
```
kubectl cluster-info --context kind-cluster1
kubectl get nodes --context kind-cluster1
kubectl get pods --all-namespaces --context kind-cluster1
```
