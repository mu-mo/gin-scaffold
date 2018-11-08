FROM golang:1.10

LABEL maintainer="yun_tofar@qq.com"
LABEL version="1.0"
LABEL description="gin scaffold"

WORKDIR /app
RUN mkdir /app/log && mkdir /app/src
ENV GOPATH "/app"
ENTRYPOINT [ "go", "run", "src/main.go" ]