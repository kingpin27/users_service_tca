# User Service for Totality Corp Assignment

## Steps to run


## Generate GRPC go files
Navigate to root of the project. And then run the command below in your terminal.
```bash
protoc *.proto --go_out=./ --go-grpc_out=./
```

## Running tests
Navigate to root of the project. And then run the command below in your terminal.
```bash
go test -v ./...
```