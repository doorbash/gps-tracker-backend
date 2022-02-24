FROM golang:1.17.1-alpine3.13 as builder
WORKDIR /go/src/app
COPY . .
RUN go build -o /app

FROM alpine
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true
COPY --from=builder /app /app
ENTRYPOINT [ "/app" ]