FROM golang:1.19-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct


WORKDIR /build

COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o /build/im_server


FROM alpine:3.14

ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /build/im_server /app
COPY --from=builder /build/configs/server.yaml /app/configs/
EXPOSE 8080
CMD ["./im_server"]