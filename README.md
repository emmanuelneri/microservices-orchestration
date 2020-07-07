# microservices-orchestration

## Requirements
- Docker
- Kubernetes
- Helm

## Environment

1. Deploy Kafka, Zookeeper and Ingress
    - Run```./deploy-infra.sh```
    - Create topics ``../create-kafka-topics``
2. Update Kafka Bootstrap servers
    - Run to get address ``kubectl get ep -n kafka -l 'app=cp-kafka'``
    - Update env `KAFKA_BOOTSTRAP_SERVERS` variable at ``/api/Dockerfile``
3. Deploy Applications
    - Build images
        - Run ``/api/build.sh``
    - Run ```./deploy-apss.sh```
    
## Execute   
1. Execute test
    - Run ``/send.sh`` to simulate call to ``http://localhost/api``
    
#### Useful Commands     

-  Force deploy apps ``./deploy-apps-force.sh``
-  Describe all topics ``./describe-kafka-topics.sh``
-  Consume Kafka topic ``./consumer-kafka-topic.sh``
