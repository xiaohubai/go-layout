FROM golang:1.18.3-alpine AS builder

WORKDIR /go/src/github.com/xiaohubai/go-layout
# COPY 源路径 目标路径
COPY . .

RUN GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o server .

FROM alpine:latest
LABEL MAINTAINER="xiaohubai@outlook.com"
# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

WORKDIR /go/src/github.com/xiaohubai/go-layout

COPY --from=builder /go/src/github.com/xiaohubai/go-layout ./

EXPOSE 8888

ENTRYPOINT ./server -c config.docker.yaml
