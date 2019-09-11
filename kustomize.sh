#!/bin/bash -e

yamls=$(find config -maxdepth 1 -type f -name "*.yaml" | sort)

for y in ${yamls[@]}; do
  kustomize edit add resource $y
done
