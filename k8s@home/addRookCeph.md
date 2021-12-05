## Add persistant storage using Rook-Ceph
Make sure the disks that should be used are wiped.

```
git clone https://github.com/rook/rook.git
cd rook/deploy/examples
kubectl create -f crds.yaml
kubectl create -f common.yaml
```

Change "cluster.yaml" to allow multiple Mons on the same node since we only have two workers.
```
allowMultiplePerNode: true
```

Create the Ceph cluster
```
kubectl create -f operator.yaml
kubectl create -f cluster.yaml
```



Thanks to:   
https://dev.to/itminds/deploying-a-ceph-cluster-with-kubernetes-and-rook-1291
https://www.adaltas.com/en/2019/09/09/rook-ceph-k8s/
