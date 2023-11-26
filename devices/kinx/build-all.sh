#!/bin/bash

for t in `ls tinygo.target.*`; do tinygo build -target=$t -size=short -o "$(echo $t | sed -r 's/(tinygo.target.+).json/\1/g').uf2"; done