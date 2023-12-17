#!/bin/bash

for t in `ls tinygo.target.*.json`; do
  echo "Building $t..."
  tinygo build -target=$t -size=short -o "$(echo $t | sed -r 's/(tinygo.target.+).json/\1/g').uf2";
done
