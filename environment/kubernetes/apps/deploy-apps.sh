#!/usr/bin/env bash


echo "--------------Deploying... --------------"

kubectl apply -f kubernetes/app-api.yaml  --namespace=apps
kubectl delete pods -l "app=api" -n apps
kubectl get services -n apps
kubectl get pods -n apps