# microservices-orchestration

## Requirements
- Docker
- Kubernetes
- Helm

## Environment

1. Deploy Kafka, Zookeeper and Ingress
    - Run```./environment/kubernetes/infra/infra-start.sh```
    - Create topics ``./environment/kubernetes/infra/configure-kafka.sh``
2. Update Kafka Bootstrap servers
    - Run to get address ``kubectl get ep -n kafka -l 'app=cp-kafka'``
    - Update env `KAFKA_BOOTSTRAP_SERVERS` variable at ``/api/Dockerfile``
3. Deploy Applications
    - Build images
        - Run ``/api/build.sh``
    - Run ```./environment/kubernetes/apps/apps-start.sh```
    
## Execute   
1. Execute test
    - Run ``/send.sh`` to simulate call to ``http://localhost/api``
    
#### Useful Commands     

-  Force deploy apps ``./environment/kubernetes/apps/deploy-apps.sh``
-  Describe all topics ``./environment/kubernetes/infra/describe-kafka-topics.sh``
-  Consume Kafka topic ``./environment/kubernetes/infra/consumer-kafka-topic.sh``
