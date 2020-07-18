# PROCESSOR 

### Build 
- build application `go build -o api-main`
- build image `docker build -t microservices-orchestration/processor .`
- login `az acr login --name microservicesOrchestrationRegistry`
- tag image `docker tag microservices-orchestration/processor microservicesorchestrationregistry.azurecr.io/processor:0.0.1`
- push image `docker push microservicesorchestrationregistry.azurecr.io/processor:0.0.1`