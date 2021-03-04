FROM alpine
RUN apk add --no-cache ca-certificates curl
COPY build/cloud-runner /usr/local/bin
