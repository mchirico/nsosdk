#!/bin/bash

unset GOSUMDB
unset GOPROXY

go fmt ./...
make docker-build docker-push IMG=quay.io/mchirico/ns-operator:v0.0.1


# kind --name v118 load docker-image ns-operator:v0.0.1
# kind --name v119 load docker-image ns-operator:v0.0.1

# Remove old CRD
make uninstall
make install


make undeploy IMG=quay.io/mchirico/ns-operator:v0.0.1
make deploy IMG=quay.io/mchirico/ns-operator:v0.0.1
