# Open Swoole GRPC Compiler

[![.github/workflows/release.yaml](https://github.com/openswoole/protoc-gen-openswoole-grpc/actions/workflows/release.yaml/badge.svg)](https://github.com/openswoole/protoc-gen-openswoole-grpc/actions/workflows/release.yaml)

Stub code generation and GRPC `.proto` compiler for [OpenSwoole GRPC](https://github.com/openswoole/grpc). You can use the `openswoole/protoc` docker images to compile and generate PHP codes or install the GRPC protoc plugin at your local environment.

## Install openswoole-grpc code generator plugin

Download package from [releases page](https://github.com/openswoole/protoc-gen-openswoole-grpc/releases).

```bash
cp ./protoc-gen-openswoole-grpc /usr/local/bin/
```

## Generate PHP stub codes from GRPC proto files

```bash
protoc --php_out=./src \
       --openswoole-grpc_out=./src helloworld.proto
# or
protoc --php_out=./src \
       --openswoole-grpc_out=./src \
       --plugin=protoc-gen-grpc=protoc-gen-openswoole-grpc \
       helloworld.proto
```

## Use openswoole/protoc Docker image to generate codes

```bash
docker run -v $APP_DIR:/app openswoole/protoc
```

## License

OpenSwoole GRPC code generator is open-sourced software licensed under the [Apache 2.0 license](https://github.com/openswoole/protoc-gen-openswoole-grpc/blob/main/LICENSE).
