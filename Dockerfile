FROM golang:1.14-alpine
WORKDIR /order_demo
ADD . /order_demo
RUN cd /order_demo \
    && cp .env.example .env \
    && go build
ENTRYPOINT ["./order_demo"]
