#!/usr/bin/env bash

OUTPUT=$(
  cat "$1" \
  | jq '.definitions.Node["x-nullable"] = false' \
)

[[ $? == 0 ]] && echo "${OUTPUT}" >| $1
