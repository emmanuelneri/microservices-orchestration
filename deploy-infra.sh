#!/usr/bin/env bash

echo "-------------- Prepare Helm --------------"
helm repo add confluentinc https://confluentinc.github.io/cp-helm-charts/
helm repo update

echo "-------------- Kafka/Zookeeper Deploying... --------------"
kubectl apply -f kubernetes/infra/kafka-namespace.yml
helm install confluent-oss -n kafka \
    --set cp-control-center.enabled=false,cp-schema-registry.enabled=false,cp-kafka-rest.enabled=false,cp-kafka-connect.enabled=false,cp-ksql-server.enabled=false \
    --set cp-kafka.prometheus.jmx.enabled=false,cp-zookeeper.prometheus.jmx.enabled=false \
    confluentinc/cp-helm-charts

echo "-------------- Metrics-server... --------------"
kubectl apply -f kubernetes/infra/metrics-server.yaml

echo "-------------- Ingress Controller... --------------"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/provider/cloud/deploy.yaml

echo "-------------- Describe --------------"
kubectl get services -n kafka
kubectl get pods -n kafka
