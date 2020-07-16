#!/usr/bin/env bash

URL=http://localhost/api
TOTAL=10000

START=$(date +%s)

echo "send starting... ${TOTAL} registers"
for i in $(seq 1 $TOTAL);
do
  identifier="${i}"
  customer="Customer ${i}"
  json="{\"identifier\": \"${identifier}\",\"customer\": \"${customer}\"}"

  curl -s -d "${json}" -H "Content-Type: application/json" -X POST ${URL} &
done