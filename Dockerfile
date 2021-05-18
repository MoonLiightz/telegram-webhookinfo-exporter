FROM golang:alpine AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
    GOPROXY=https://goproxy.io

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -v -o app cmd/tgwe/main.go


FROM scratch

ARG BUILD_DATE
ARG VCS_REF
ARG VERSION

LABEL maintainer="MoonLiightz" \
  org.label-schema.build-date=$BUILD_DATE \
  org.label-schema.name="telegram-webhookinfo-exporter" \
  org.label-schema.description="Docker image for telegram-webhookinfo-exporter" \
  org.label-schema.version=$VERSION \
  org.label-schema.url="https://github.com/MoonLiightz/telegram-webhookinfo-exporter" \
  org.label-schema.vcs-ref=$VCS_REF \
  org.label-schema.vcs-url="https://github.com/MoonLiightz/telegram-webhookinfo-exporter" \
  org.label-schema.vendor="MoonLiightz" \
  org.label-schema.schema-version="1.0"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/app /app

EXPOSE 2112

ENTRYPOINT ["/app"]
