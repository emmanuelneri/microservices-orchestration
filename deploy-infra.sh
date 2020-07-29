#!/usr/bin/env bash

echo "-------------- create Cluster --------------"
kind create cluster

echo "-------------- Kafka/Zookeeper Deploying... --------------"
kubectl create namespace kafka
kubectl apply -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
kubectl apply -f https://strimzi.io/examples/latest/kafka/kafka-persistent.yaml -n kafka

echo "-------------- Metrics-server... --------------"
kubectl apply -f kubernetes/infra/metrics-server.yaml

echo "-------------- Ingress Controller... --------------"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud/deploy.yaml

echo "-------------- Describe --------------"
kubectl wait kafka/my-cluster --for=condition=Ready --timeout=300s -n kafka
kubectl get services -n kafka
kubectl get pods -n kafka
