FROM golang:latest

ENV GOPROXY https://goproxy.cn

WORKDIR /build

COPY . /build

RUN go build -o main

EXPOSE 8848

ENTRYPOINT [ "/build/main" ]