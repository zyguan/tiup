FROM golang:1.13.8-alpine3.11 AS builder

RUN apk add --no-cache git build-base
WORKDIR /project
COPY go.mod go.sum ./
RUN go mod download
ADD . .
RUN go build -o fbproxy


FROM alpine:3.11

RUN apk add --no-cache ca-certificates
COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /project/fbproxy /usr/bin/fbproxy

ENTRYPOINT [ "/usr/bin/fbproxy" ]

# hub.pingcap.net/zyguan/fbproxy
