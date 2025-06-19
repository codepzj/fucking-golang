# 初识grpc

## 安装

```bash
go get -u google.golang.org/grpc
```

安装代码生成工具
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

创建一个简单的`hello.proto`文件

```proto
syntax = "proto3";

option go_package = ".;service";

service SayHello {
  rpc SayHi(HelloRequest) returns (HelloResponse){}
}

message HelloRequest {
  string RequestMsg = 1;
}

message HelloResponse {
  string ResponseMsg = 1;
}
```

使用以下命令生成对应的go代码

```bash
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
```
