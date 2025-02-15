# fullcycle_gRPC

## Preparar ambiente

1- Instalar compilador do protoc:
```bash
brew install protobuf-compiler
```
2-Instalar o plugin do protoc-gen-go:
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```
3- Instalar o plugin do protoc-gen-go-grpc:
```bash
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Projeto

Gerar entidades e interfaces - chamada do plugin para o arquivo proto
```
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
```