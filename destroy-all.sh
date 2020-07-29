#!/usr/bin/env bash

echo "-------------- Starting stop Cluster --------------"
kubectl -n apps delete pod,service,deployment,ingress,hpa --all
kubectl -n ingress-nginx delete pod,service,deployment --all

kubectl delete statefulset --all -n kafka
kubectl delete pvc --all -n kafka
kind delete cluster