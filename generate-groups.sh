#!/bin/bash

scriptPath=${SCRIPT_PATH}

if [ -z ${scriptPath} ]; then
  echo "Tell me where is the code-generator script? (SCRIPT_PATH)"
  echo "Maybe 'export SCRIPT_PATH=~/go/src/github.com/kubernetes/code-generator/generate-groups.sh'?"
  exit 1
fi

${scriptPath} \
  all \
  github.com/spongeprojects/client-go/client \
  github.com/spongeprojects/client-go/api \
  spongeprojects.com:v1alpha1 \
  --go-header-file ./boilerplate.go.txt
