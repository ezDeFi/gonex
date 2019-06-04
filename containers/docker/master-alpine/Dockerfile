FROM alpine:3.7

RUN \
  apk add --update go git make gcc musl-dev linux-headers ca-certificates && \
  git clone --depth 1 --branch release/2.0 https://github.com/nextyio/gonex && \
  (cd gonex && make gonex) && \
  cp gonex/build/bin/gonex /nexty && \
  apk del go git make gcc musl-dev linux-headers && \
  rm -rf /gonex && rm -rf /var/cache/apk/*

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/nexty"]
