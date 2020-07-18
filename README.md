# microservices-orchestration

## Local deployment

#### Requirements
- Docker
- Kubernetes
- Helm

#### Environment

1. Deploy Kafka, Zookeeper and Ingress
    - Run```./deploy-infra.sh```
    - Create topics ``../create-kafka-topics``
2. Deploy Applications
    - Build images
        - Run ``/api/build.sh``
    - Run ```./deploy-apps.sh```
    
#### Execute   
1. Execute test
    - Run ``/send.sh`` to simulate call to ``http://localhost/api``
    
#### Useful Commands     

-  Force deploy apps ``./deploy-apps-force.sh``
-  Describe all topics ``./describe-kafka-topics.sh``
-  Consume Kafka topic ``kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic ApiRequested --from-beginning``
-  List Kafka Consumer Group``kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-consumer-groups --bootstrap-server localhost:9092 --list``
-  Describe Kafka Consumer Group``kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group processor-consumer-group``
-  Repartition topic``kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-topics --alter --topic ApiRequested --partitions 6 --replication-factor 3 --zookeeper confluent-oss-cp-zookeeper:2181``



## Azure Deployment

- Resource Group
```
az group create \
    --name microservices-orchestration \
        --location eastus
```
```
az aks create \
    --resource-group microservices-orchestration \
    --name microservices-orchestration-cluster \
    --location eastus \
    --node-count 2 \
    --generate-ssh-key
```

- Kubernetes (AKS)
```
az aks create \
    --resource-group microservices-orchestration \
    --name microservices-orchestration-cluster \
    --location eastus \
    --node-count 2 \
    --generate-ssh-key
```
```
az acr create --resource-group microservices-orchestration \
  --name microservicesOrchestrationRegistry --sku Basic
```

- Kafka (HD Insight)
```
az storage account create \
    --name kafkastorageaccount \
    --resource-group microservices-orchestration \
    --https-only true \
    --kind StorageV2 \
    --location eastus \
    --sku Standard_LRS
```

Get key1 value to populate account-key on container creation
```az storage account keys list -g microservices-orchestration -n kafkastorageaccount```

```
az storage container create \
    --name kafkastoragecontainer \
    --account-key ******** \
    --account-name kafkastorageaccount
```
```
az hdinsight create \
    --name kafka-hdinsight-cluster \
    --resource-group microservices-orchestration \
    --type kafka \
    --component-version kafka=2.1 \
    --http-password ****** \
    --http-user admin \
    --location eastus \
    --ssh-password ****** ******  \
    --ssh-user sshuser \
    --storage-account kafkastorageaccount \
    --storage-account-key ******  \
    --storage-container kafkastoragecontainer \
    --version 4.0 \
    --workernode-count 3 \
    --workernode-data-disks-per-node 2 \
    --kafka-management-node-size "Standard_D4_v2" \
    --kafka-client-group-id kafkagroupid \
    --kafka-client-group-name "kafkagroup"
```

