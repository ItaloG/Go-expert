apt install -y protobuf-compiler

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# https://github.com/ktr0731/evans
go install github.com/ktr0731/evans@latest

protoc --go_out=. --go-grpc_out=. proto/course_category.proto