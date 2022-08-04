### Compile GRPC files

```
protoc datafiles/transaction.proto --go_out=datafiles --go-grpc_out=require_unimplemented_servers=false:datafiles
```
