FROM ethereum/client-go:latest

RUN apk add --no-cache jq

COPY entrypoint.sh /entrypoint.sh

VOLUME ["/db"]

ENTRYPOINT ["/bin/sh", "/entrypoint.sh"]