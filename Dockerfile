# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY geoip/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"'

# server image
FROM debian:buster-slim
LABEL org.opencontainers.image.authors="tekmanic"
COPY geoip/public/ /srv/geoip/public/
COPY geoip/internal/views/ /srv/geoip/internal/views/
COPY --from=builder /src/geoip /usr/local/bin/
ENV GEOIP_DIR /srv/geoip/
EXPOSE 3000
CMD ["/usr/local/bin/geoip"]
# ENTRYPOINT [ "tail", "-f", "/dev/null" ]