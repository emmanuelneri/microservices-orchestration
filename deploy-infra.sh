#!/usr/bin/env bash

echo "-------------- Kafka/Zookeeper Deploying... --------------"
kubectl apply -f kubernetes/infra/kafka-namespace.yml
helm repo add confluentinc https://confluentinc.github.io/cp-helm-charts/
helm repo update
helm install -n kafka --set cp-schema-registry.enabled=false,cp-kafka-rest.enabled=false,cp-kafka-connect.enabled=false,cp-ksql-server.enabled=false --generate-name confluentinc/cp-helm-charts

echo "-------------- Ingress Controller... --------------"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud/deploy.yaml

echo "-------------- Describe --------------"
kubectl get services -n kafka
kubectl get pods -n kafka

echo "-------------- Kafka Endpoints --------------"
kubectl get ep -n kafka -l 'app=cp-kafka'
