#!/bin/bash

unset GOSUMDB
unset GOPROXY

go fmt ./...
make docker-build IMG=ns-operator:v0.0.1

kind --name v118 load docker-image ns-operator:v0.0.1
kind --name v119 load docker-image ns-operator:v0.0.1

# Remove old CRD
make uninstall
make install


make undeploy IMG=ns-operator:v0.0.1
make deploy IMG=ns-operator:v0.0.1
