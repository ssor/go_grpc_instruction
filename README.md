
# Instruction

### 1. install protoc
```
brew install protobuf
```
> This will install **protoc** to **$PATH**

### 2. install go plugin

```
cd $GOPATH/src/github.com/golang/protobuf

make
```
> This will build **protoc-gen-go** to **$GOBIN**

### 3. compile .proto file

```
 protoc --go_out=plugins=grpc:. hellogrpc.proto
 ```

> This will generate **hellogrpc.pb.go**

