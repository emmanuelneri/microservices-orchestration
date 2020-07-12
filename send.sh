#!/usr/bin/env bash

URL=http://localhost/api
TOTAL=100

START=$(date +%s)

echo "send starting... ${TOTAL} registers"
for i in $(seq 1 $TOTAL);
do
  customer="Customer ${i}"
  json="{\"customer\": \"${customer}\"}"

  curl -s -d "${json}" -H "Content-Type: application/json" -X POST ${URL} &
done