#!/usr/bin/env bash

kafkaRest=https://kafka-hdinsight-cluster-kafkarest.azurehdinsight.net

curl -X POST -H "Content-Type: application/vnd.api+json" -H "Accept: application/vnd.api+json" \
          --data '{"data":{"attributes": {"topic_name": "ApiRequested", "partitions_count": 6, "replication_factor": 3}}}' \
          "${kafkaRest}/topics"
