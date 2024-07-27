FROM golang:1.20.13-alpine as builder
LABEL MAINTAINER="VinCent.hu@koalaplatform.com"

WORKDIR /go/src/webSite_statistics

COPY . /go/src/webSite_statistics

# 运行可执行文件
CMD ["./wynpay_webSite_linux"]
EXPOSE 5001
