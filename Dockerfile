#FROM golang:alpine
FROM golang:latest

ENV ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=arm64dock

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o stratton

CMD [ "./stratton", "-name=depth5", "-asset=ETH,BTC/BUSD" ]

