#!/usr/bin/env bash

namespace=kafka
kafka=my-cluster-kafka-0

echo "-------------- Kafka Describe Topic --------------"
kubectl exec -it -c kafka ${kafka} -n ${namespace} -- bin/kafka-topics.sh --list --zookeeper localhost:2181