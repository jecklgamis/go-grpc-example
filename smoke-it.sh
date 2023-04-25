#!/usr/bin/env bash
PAYLOAD=$(cat << 'EOF'
{ "key": "some-key", "value": "some-value" }
EOF)
echo "Storing KV"
curl -X PUT http://localhost:8080/v1 -d"${PAYLOAD}"
echo "Retrieving KV"
curl  http://localhost:8080/v1?key=some-key
