# User Service for Totality Corp Assignment

- The code is split into two packages i.e. server and model
- model packages all the code related to storage, retrival and management of user data. I here have followed **strategy pattern** to abstract hoow information is stored from behaviour using interfaces.
- I have written test using standard go testing library.
- server package contains all code related to exposing functionality of user model to outside world using GRPC.

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