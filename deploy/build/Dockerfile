#build stage
FROM golang:alpine AS builder

ARG APP_RELATIVE_PATH

RUN apk add --no-cache git

COPY ddns_server /src
WORKDIR /src/app/${APP_RELATIVE_PATH}

RUN go env -w GOPROXY=https://goproxy.cn,direct && go get -d -v ./...
RUN make build

#final stage
FROM alpine:latest
RUN apk --no-cache add tzdata ca-certificates
COPY --from=builder /src/app/${APP_RELATIVE_PATH}/bin /app
ENTRYPOINT /app
LABEL Name=hominsu/ali-ddns-server Version=1.1
EXPOSE 50001 50002
