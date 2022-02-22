# Open Swoole GRPC code generator plugin

```bash

# install openswoole-grpc code generator plugin

cp ./protoc-gen-openswoole-grpc /usr/local/bin/

protoc --php_out=./src --openswoole-grpc_out=./src helloworld.proto

# or

protoc --php_out=./src --openswoole-grpc_out=./src --plugin=protoc-gen-openswoole-grpc=protoc-gen-openswoole-grpc helloworld.proto
```

