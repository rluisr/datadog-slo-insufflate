# syntax = docker/dockerfile:1.3-labs

FROM golang:1-alpine as builder
ARG VERSION=0.0.0
WORKDIR /go/src/datadog-slo-insufflate
COPY . .
RUN apk --no-cache add git openssh build-base
RUN cd cmd && go build -ldflags "-X main.version=${VERSION}" -o app .

FROM alpine as production
LABEL maintainer="rluisr" \
  org.opencontainers.image.url="https://github.com/rluisr/datadog-slo-insufflate" \
  org.opencontainers.image.source="https://github.com/rluisr/datadog-slo-insufflate" \
  org.opencontainers.image.vendor="rluisr" \
  org.opencontainers.image.title="datadog-slo-insufflate" \
  org.opencontainers.image.description="TradingView webhook handler for Bybit." \
  org.opencontainers.image.licenses="AGPL"
RUN apk add --no-cache ca-certificates libc6-compat
COPY --from=builder /go/src/datadog-slo-insufflate/app /app
ENTRYPOINT ["/app"]