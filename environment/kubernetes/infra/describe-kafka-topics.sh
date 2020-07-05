#!/usr/bin/env bash

namespace=kafka
kafka=($(kubectl get pods -n ${namespace} -l 'app=cp-kafka' \
    -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')*/)

zookeeper=($(kubectl get pods -n ${namespace} -l 'app=cp-zookeeper' \
    -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')*/)

zookeepers=$(sed 's/.\{2\}$//'<<< $zookeeper)":2181"

echo "-------------- Kafka Describe Topic --------------"
echo "kafka selected: ${kafka}"
echo "zookeeper: ${zookeepers}"

kubectl exec -c cp-kafka-broker -it ${kafka} -n ${namespace} -- /bin/bash /usr/bin/kafka-topics --describe --zookeeper ${zookeepers}