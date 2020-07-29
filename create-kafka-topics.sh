#!/usr/bin/env bash

namespace=kafka
kafka=my-cluster-kafka-0

echo "-------------- Kafka Create Topics --------------"
kubectl exec -c kafka -it ${kafka} -n ${namespace} -- bin/kafka-topics.sh --create --topic ApiRequested --partitions 6 --replication-factor 3 --if-not-exists --zookeeper localhost:2181