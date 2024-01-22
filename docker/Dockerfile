FROM alpine:3.6 as alpine
RUN apk add -U --no-cache ca-certificates

FROM alpine:3.6
ENV GODEBUG netdns=go

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

RUN apk update

RUN apk add gcompat

ADD release/linux/amd64/drone-helm /bin/
ENTRYPOINT ["/bin/drone-helm"]