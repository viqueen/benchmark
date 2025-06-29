#!/usr/bin/env bash

_clean() {
  find ./api/go-sdk -name "*.go" -exec rm -f {} \;
}


_protobuf_gen() {
    docker run --rm \
    		--volume "${PWD}"/_schema:/workspace/_schema \
    		--volume "${PWD}"/api:/workspace/api \
    		--workdir /workspace/_schema \
    		bufbuild/buf:1.51.0 generate --verbose
}

_clean
_protobuf_gen
