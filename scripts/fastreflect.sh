#!/bin/sh

set -e

build() {
    echo finding protobuf files in "$1"
    proto_files=$(find "$1" -name "*.proto")
    for file in $proto_files; do
      echo "building proto file $file"
      protoc -I=. -I=./third_party/proto --plugin /usr/bin/protoc-gen-go-pulsar --go-pulsar_out=. --go-pulsar_opt=features=fast+protoc+interfaceservice "$file"
    done
}

for dir in "$@"
do
  build "$dir"
done

cp -r github.com/cosmos/cosmos-proto/* ./
rm -rf github.com