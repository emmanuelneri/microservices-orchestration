#!/usr/bin/env bash

namespace=kafka
kafka=confluent-oss-cp-kafka-0
zookeeper=confluent-oss-cp-zookeeper:2181

echo "-------------- Kafka Create Topics --------------"
echo "kafka selected: ${kafka}"
echo "zookeeper: ${zookeeper}"

kubectl exec -c cp-kafka-broker -it ${kafka} -n ${namespace} -- /bin/bash /usr/bin/kafka-topics --create --topic ApiRequested --partitions 6 --replication-factor 3 --if-not-exists --zookeeper ${zookeeper}