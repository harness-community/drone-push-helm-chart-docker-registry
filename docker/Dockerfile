FROM debian:11-slim

ENV GODEBUG netdns=go

ADD release/linux/amd64/drone-helm /bin/

ENTRYPOINT ["/bin/drone-helm"]