# nsosdk


### Requirements

-  [Operator SDK](https://sdk.operatorframework.io/docs/installation/)
-  [Go 1.15](https://golang.org/dl/)


Need KinD

```bash
GO111MODULE="on" go get sigs.k8s.io/kind@v0.10.0
```



```bash
operator-sdk init --domain cwxstat.com --repo github.com/mchirico/ns-operator
operator-sdk create api --group top --version v1alpha1 --kind Namespecial --resource --controller

```

## Steps to Build

```bash
make docker-build IMG=ns-operator:v0.0.1

kind load docker-image ns-operator:v0.0.1


make docker-build docker-push IMG=quay.io/mchirico/ns-operator:v0.0.1

# Outside cluster
make install run


# Inside cluster
make deploy IMG=quay.io/mchirico/ns-operator:v0.0.1
make undeploy IMG=quay.io/mchirico/ns-operator:v0.0.1

```


## Example how to use


```bash
k get po -A

NAMESPACE            NAME                                         READY   STATUS    RESTARTS   AGE
kube-system          coredns-66bff467f8-hnmqd                     1/1     Running   0          17h
kube-system          coredns-66bff467f8-mdbwl                     1/1     Running   0          17h
kube-system          etcd-v118-control-plane                      1/1     Running   0          17h
kube-system          kindnet-8b8ws                                1/1     Running   0          17h
kube-system          kindnet-jkkhl                                1/1     Running   0          17h
kube-system          kube-apiserver-v118-control-plane            1/1     Running   0          17h
kube-system          kube-controller-manager-v118-control-plane   1/1     Running   0          17h
kube-system          kube-proxy-98km7                             1/1     Running   0          17h
kube-system          kube-proxy-kgt5r                             1/1     Running   0          17h
kube-system          kube-scheduler-v118-control-plane            1/1     Running   0          17h
local-path-storage   local-path-provisioner-5b4b545c55-zh45g      1/1     Running   0          17h
nsosdk-system        nsosdk-controller-manager-5974f79c64-g7rhv   2/2     Running   0          2m54s

```



```bash
k logs nsosdk-controller-manager-5974f79c64-g7rhv -n nsosdk-system manager
```
