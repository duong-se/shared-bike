FROM golang:1.18.3-alpine3.16 as builder
RUN apk add --no-cache tzdata

ADD ./ /shared-bike
WORKDIR /shared-bike
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -o cmd/main -v *.go

FROM alpine:3.15.0
WORKDIR /
RUN apk add --no-cache tzdata
COPY --from=builder /shared-bike/cmd/main main

EXPOSE 8000

ENTRYPOINT ["/main"]
