# Open Swoole GRPC PHP code generator

[![.github/workflows/release.yaml](https://github.com/openswoole/protoc-gen-openswoole-grpc/actions/workflows/release.yaml/badge.svg)](https://github.com/openswoole/protoc-gen-openswoole-grpc/actions/workflows/release.yaml)

## Prerequisites

* PHP and Composer
* Open Swoole GRPC
* Protocol Buffers Compiler <https://github.com/protocolbuffers/protobuf>

## Install openswoole-grpc code generator plugin

Download package from [releases page](https://github.com/openswoole/protoc-gen-openswoole-grpc/releases).

```bash
cp ./protoc-gen-openswoole-grpc /usr/local/bin/
```

## Generate PHP stub codes from GRPC proto files

```bash
protoc --php_out=./src --openswoole-grpc_out=./src helloworld.proto
# or
protoc --php_out=./src --openswoole-grpc_out=./src --plugin=protoc-gen-grpc=protoc-gen-openswoole-grpc helloworld.proto
```
