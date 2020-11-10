#!/usr/bin/env bash

OUTPUT=$(jq '.definitions.Node["x-nullable"] = false' $1)

[[ $? == 0 ]] && echo "${OUTPUT}" >| $1
