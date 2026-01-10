#!/bin/bash

cd images

for dir in */; do
  cd "$dir"
  dir="${dir%/}"
  podman build --label app=codeexec -t "$dir" . 
  cd ..
done

