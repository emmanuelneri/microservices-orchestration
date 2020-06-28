#!/usr/bin/env bash

echo "-------------- Kafka namespace --------------"
kubectl get services -n kafka
kubectl get pods -n kafka

echo "-------------- APPs namespace --------------"
kubectl get services -n apps
kubectl get pods -n apps
kubectl get ingress ingress -n apps


echo "-------------- APP --------------"
echo "[API] http://localhost/api (POST)"