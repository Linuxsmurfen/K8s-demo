## Howto setup FluxCD

### Installation

Install the flux binary on the system
```
curl -s https://fluxcd.io/install.sh | sudo bash
flux --version
```

Export git token
```
export GITHUB_TOKEN=<....>
```

Create the repository on git and set everything up in the K8s cluster.
```
flux bootstrap github \
  --owner=Linuxsmurfen \
  --repository=Mini-k8s-cluster \
  --path=flux \
  --personal
```


### Uninstall
How to uninstall flux from the k8s cluster
```
flux uninstall --namespace=flux-system
```
   
   

### Watch for sync 
```
flux get kustomizations --watch
```


Thanks to:   
https://fluxcd.io/docs/get-started/
