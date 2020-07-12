#!/usr/bin/env bash


echo "--------------Deploying... --------------"

# TODO deploy rolling update
kubectl apply -f kubernetes/apps/app-api.yaml  --namespace=apps
kubectl apply -f kubernetes/apps/app-processor.yaml  --namespace=apps
kubectl delete pods -l "app=api" -n apps
kubectl delete pods -l "app=processor" -n apps
kubectl get services -n apps
kubectl get pods -n apps