#!/bin/bash

cd images

for dir in */; do
  cd "$dir"
  dir="${dir%/}"
  docker build -t "$dir" . 
  cd ..
done

