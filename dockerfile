FROM golang:1.18 AS build

WORKDIR /go/src/github.com/correavitor4/concorrentes2
COPY . .

RUN go build -o app .

FROM alpine:latest

WORKDIR /root/

COPY --from=build /go/src/github.com/correavitor4/concorrentes2 .

ENTRYPOINT [ "./app" ]