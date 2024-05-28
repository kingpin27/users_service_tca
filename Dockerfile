FROM golang:1.22

WORKDIR /app

RUN apt-get update && apt-get install -y \
    unzip \
    wget  \
    protobuf-compiler

COPY go.mod go.sum ./
RUN go mod download

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY *.go ./
COPY ./model/* ./model/
COPY ./server/* ./server/
COPY *.proto ./

RUN protoc *.proto --go_out=./ --go-grpc_out=./

RUN CGO_ENABLED=0 GOOS=linux go build -o /user_service_tca

EXPOSE 8080

CMD ["/user_service_tca"]