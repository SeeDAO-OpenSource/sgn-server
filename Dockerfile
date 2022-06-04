FROM golang:1.18 as builder

ENV GIN_MODE=release \
    CGO_ENABLED=0 \
    GOSUMDB=off \
    GOARCH=amd64

RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
WORKDIR /app

COPY . .
RUN go build .
RUN mkdir publish && cp sgn publish && cp -r app publish

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/publish .
# COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/cert

ENV GIN_MODE=release \
    SGN_SERVER_PORT=80
EXPOSE 80

ENTRYPOINT [ "./sgn","server" ]