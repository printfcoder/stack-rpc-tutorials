FROM alpine

ENV MICRO_REGISTRY consul
ENV MICRO_REGISTRY_ADDRESS 172.17.0.3:8500

RUN apk update && apk add tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

ADD go-micro-demo-client /go-micro-demo-client

WORKDIR /
ENTRYPOINT [ "/go-micro-demo-client" ]