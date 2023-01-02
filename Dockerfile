FROM golang:alpine

ENV GOPROXY https://goproxy.cn,direct \
    GO111MODULE on \
    
WORKDIR /build

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080

CMD ["./dist/main"]