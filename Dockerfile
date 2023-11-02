FROM alpine:3.18

WORKDIR /app

ENV HELM_EXPERIMENTAL_OCI=1

RUN apk --no-cache add ca-certificates curl && \
    curl -LO "https://get.helm.sh/helm-v3.7.0-linux-amd64.tar.gz" && \
    tar -zxvf helm-v3.7.0-linux-amd64.tar.gz && \
    mv linux-amd64/helm /usr/local/bin/helm && \
    rm -rf linux-amd64 helm-v3.7.0-linux-amd64.tar.gz && \
    chmod +x /usr/local/bin/helm

RUN apk --no-cache add python3

COPY main.py /app/

CMD ["python3", "/app/main.py"]
