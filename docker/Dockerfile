FROM alpine:3.6 as alpine

RUN apk add -U --no-cache ca-certificates && \
    apk add gcompat

ENV GODEBUG netdns=go

ADD release/linux/amd64/drone-helm /bin/

ENTRYPOINT ["/bin/drone-helm"]