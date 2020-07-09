#!/usr/bin/env bash

KUBERNETES_DIRECTORY="kubernetes/apps/"

echo "-------------- Apps Deploying... --------------"
kubectl apply -f ${KUBERNETES_DIRECTORY}/apps-namespace.yml
kubectl apply -f ${KUBERNETES_DIRECTORY}/app-api.yaml  --namespace=apps
kubectl apply -f ${KUBERNETES_DIRECTORY}/app-processor.yaml  --namespace=apps
kubectl apply -f ${KUBERNETES_DIRECTORY}/ingress.yaml  --namespace=apps

echo "-------------- APPs --------------"
kubectl get ingress ingress -n apps
kubectl get services -n apps
kubectl get pods -n apps