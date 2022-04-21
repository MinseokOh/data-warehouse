#!/bin/bash

proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
  echo "$dir"
done

for dir in $proto_dirs; do
  protoc -I="proto" \
    --go_out=. \
    --go-grpc_out=. \
    $(find "${dir}" -maxdepth 1 -name '*.proto')
done

cp -r github.com/MinseokOh/data-warehouse/* ./
rm -rf github.com
