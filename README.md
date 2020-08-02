# microservices-orchestration
------------------------------------------------------

![alt tag](https://github.com/emmanuelneri/microservices-orchestration/blob/master/architecture.png?style=centerme)

##### TODO
- schema registry
- commons file segregation review
- requestId (HTTP response - Kafka Header)
- project structure revision
- dockerfile revision
- migrate Ingress namespace
- request/limits configuration revision

## Local deployment

#### Requirements
- Docker
- Kubernetes
- Helm

#### Environment

1. Deploy Kafka, Zookeeper and Ingress
    - Run`./deploy-infra.sh`
    - Create topics `./create-kafka-topics`
2. Deploy Applications
    - Build images
        - Run `./api/build.sh`
    - Run `./deploy-apps.sh`
    
#### Execute   
1. To test
    - Run `./send.sh` to simulate call to `http://localhost/api`
    
#### Useful Commands     

-  Force deploy apps `./deploy-apps-force.sh`
-  Describe all topics `./describe-kafka-topics.sh`
-  Consume Kafka topic `kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-console-consumer --bootstrap-server localhost:9092 --topic ApiRequested --from-beginning`
-  List Kafka Consumer Group`kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-consumer-groups --bootstrap-server localhost:9092 --list`
-  Describe Kafka Consumer Group`kubectl exec -c cp-kafka-broker -it confluent-oss-cp-kafka-0 -n kafka -- /bin/bash /usr/bin/kafka-consumer-groups --bootstrap-server localhost:9092 --describe --group processor-consumer-group`

#### Stop all

-  Destroy all resources `./destroy-all.sh`


## Local development

#### Requirements
- Docker
- Go

#### Environment

1. Start Kafka `docker-compose up`
2. Run API application `go run main.go` (inside api directory)
3. Run Processor application `go run main.go` (inside processor directory)

#### Execute   

1. To test
    - Run endpoint request `curl -v -d "{\"identifier\": \"123\",\"customer\": \"Customer 1\"}" -H "Content-Type: application/json" -X POST http://localhost:8080`