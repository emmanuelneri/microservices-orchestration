#!/usr/bin/env bash

echo "-------------- Starting stop Cluster --------------"
kubectl -n kafka delete pod,service,deployment --all
kubectl -n apps delete pod,service,deployment,ingress --all