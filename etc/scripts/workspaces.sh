#!/bin/bash

modules=$(go work edit -json | jq -r '.Use[].DiskPath' | grep -v './examples')

for m in $modules; do
  echo "> $m"
  (cd "$m" && $@)
done
