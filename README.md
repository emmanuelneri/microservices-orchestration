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

`az aks get-credentials --resource-group microservices-orchestration --name microservices-orchestration-cluster --admin`



