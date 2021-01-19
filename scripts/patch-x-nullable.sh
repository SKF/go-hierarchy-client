#!/usr/bin/env bash

OUTPUT=$(
  cat "$1" \
  | jq '.definitions.Node["x-nullable"] = false' \
  | jq '.definitions.Origin["x-nullable"] = false' \
  | jq '.definitions.NodeMetaData.additionalProperties["x-nullable"] = true' \
)

[[ $? == 0 ]] && echo "${OUTPUT}" >| $1
