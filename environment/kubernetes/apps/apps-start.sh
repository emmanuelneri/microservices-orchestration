#!/usr/bin/env bash

KUBERNETES_DIRECTORY="/"

echo "-------------- Apps Deploying... --------------"
kubectl apply -f apps-namespace.yml
kubectl apply -f app-api.yaml  --namespace=apps
kubectl apply -f ingress.yaml  --namespace=apps

echo "-------------- APPs --------------"
kubectl get ingress ingress -n apps
kubectl get services -n apps
kubectl get pods -n apps