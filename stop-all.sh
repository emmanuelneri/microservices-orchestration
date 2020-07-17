#!/usr/bin/env bash

echo "-------------- Starting stop Cluster --------------"
kubectl -n apps delete pod,service,deployment,ingress,hpa --all
kubectl -n ingress-nginx delete pod,service,deployment --all

helm delete -n kafka $(helm -n kafka ls --short)
kubectl delete statefulset --all -n kafka
kubectl delete pvc --all -n kafka