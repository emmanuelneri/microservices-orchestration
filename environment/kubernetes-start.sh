#!/usr/bin/env bash

KUBERNETES_DIRECTORY="/"

echo "-------------- Kafka/Zookeeper Deploying... --------------"
kubectl apply -f kubernetes/kafka-namespace.yml
kubectl apply -f kubernetes/zookeeper.yaml  --namespace=kafka
kubectl apply -f kubernetes/kafka.yaml  --namespace=kafka

echo "-------------- Apps Deploying... --------------"
kubectl apply -f kubernetes/apps-namespace.yml
kubectl apply -f kubernetes/app-api.yaml  --namespace=apps

echo "-------------- Infra Deploying... --------------"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud/deploy.yaml
kubectl apply -f kubernetes/ingress.yaml  --namespace=apps

./describe-environment.sh