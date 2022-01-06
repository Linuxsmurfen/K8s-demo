## Howto setup FluxCD + SOPS

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


flux create kustomization my-secrets \
--source=my-secrets \
--path=./clusters/cluster0 \
--prune=true \
--interval=10m \
--decryption-provider=sops \
--decryption-secret=sops-gpg



### Uninstall
How to uninstall flux from the k8s cluster
```
flux uninstall --namespace=flux-system
```
   
### Check
```
flux check
```


### Watch for sync 
```
flux get kustomizations --watch
```



## SOPS
Install gpg and sops
Check for latest version on: https://github.com/mozilla/sops/releases

```
wget https://github.com/mozilla/sops/releases/download/v3.7.1/sops-3.7.1-1.x86_64.rpm
sudo apt install ./sops_3.7.1_amd64.deb
sops --version
```

Generate GPG key
```
export KEY_NAME="cluster0.yourdomain.com"
export KEY_COMMENT="flux secrets"

gpg --batch --full-generate-key <<EOF
%no-protection
Key-Type: 1
Key-Length: 4096
Subkey-Type: 1
Subkey-Length: 4096
Expire-Date: 0
Name-Comment: ${KEY_COMMENT}
Name-Real: ${KEY_NAME}
EOF
```

Export the key and import it into K8s
```
gpg --export-secret-keys --armor "${KEY_FP}" |
kubectl create secret generic sops-gpg \
--namespace=flux-system \
--from-file=sops.asc=/dev/stdin
```



Thanks to:   
https://fluxcd.io/docs/get-started/   
https://blog.sldk.de/2021/02/introduction-to-gitops-on-kubernetes-with-flux-v2/  
   
https://fluxcd.io/docs/guides/mozilla-sops/   
https://blog.sldk.de/2021/03/handling-secrets-in-flux-v2-repositories-with-sops/  




