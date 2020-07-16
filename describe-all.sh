#!/usr/bin/env bash

echo "-------------- Kafka --------------"
kubectl get services -n kafka
kubectl get pods -n kafka

echo "-------------- Ingress --------------"
kubectl get ingress ingress -n apps

echo "-------------- Configmap --------------"
kubectl describe configmaps default -n apps

echo "-------------- Autoscaling --------------"
kubectl get deployment metrics-server -n kube-system
kubectl describe hpa -n apps

echo "-------------- Describe --------------"
kubectl get services -n apps
kubectl get pods -n apps
