# step 1

FROM golang:1.10.4 as step1

RUN mkdir -p /go/src/go-api-example

COPY . /go/src/go-api-example

WORKDIR /go/src/go-api-example

RUN go get -v && CGO_ENABLED=0 GOOS=linux go build main.go 
# step 2

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

RUN apk --no-cache add curl

EXPOSE 90

WORKDIR /root/

COPY --from=step1 /go/src/go-api-example/main .

CMD ["./main"]