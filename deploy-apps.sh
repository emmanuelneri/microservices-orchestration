#!/usr/bin/env bash


echo "-------------- Apps Deploying... --------------"
kubectl apply -f kubernetes/apps/apps-namespace.yml
kubectl apply -f kubernetes/configmap/default-config.yaml

kubectl apply -f kubernetes/apps/app-api.yaml  --namespace=apps
kubectl apply -f kubernetes/apps/app-processor.yaml  --namespace=apps
kubectl apply -f kubernetes/apps/ingress.yaml  --namespace=apps

echo "-------------- APPs --------------"
kubectl get ingress ingress -n apps
kubectl get services -n apps
kubectl get pods -n apps