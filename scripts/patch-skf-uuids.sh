#!/usr/bin/env bash

NEEDLE='"format": "uuid"'
read -r -d '' REPLACE << EOM
"x-go-type": {
  "type": "UUID",
  "import": {
    "package": "github.com/SKF/go-utility/v2/uuid"
  },
  "hints": {
    "kind": "object"
  }
},
"format": "uuid"
EOM

# Escape replacement for sed
REPLACE=$(sed 's/[&/\]/\\&/g' <<<"$(echo ${REPLACE})")

sed -i "s/${NEEDLE}/${REPLACE}/g" $1