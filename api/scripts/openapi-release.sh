#!/bin/bash

for f in openapi/*; do
  if [ -d "$f" ]; then
    echo "Preparing release ..."
    npm run openapi -- validate ./openapi/client/api.yaml
    npm run openapi -- bundle --output ./docs/release/client.yaml ./openapi/client/api.yaml
  fi
done
