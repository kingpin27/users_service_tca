FROM golang:1.22

WORKDIR /app

RUN apt-get update && apt-get install -y \
    unzip \
    wget
RUN sudo apt-get install protobuf-compiler

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /user_service_tca

EXPOSE 8080

CMD ["/user_service_tca"]