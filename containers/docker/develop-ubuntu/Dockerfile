FROM ubuntu:bionic

ENV PATH=/usr/lib/go-1.10/bin:$PATH

RUN \
  apt-get update && apt-get upgrade -q -y && \
  apt-get install -y --no-install-recommends golang-1.10 git make gcc libc-dev ca-certificates && \
  git clone --depth 1 https://github.com/nextyio/gonex && \
  (cd gonex && make gonex) && \
  cp gonex/build/bin/gonex /nexty && \
  apt-get remove -y golang-1.10 git make gcc libc-dev && apt autoremove -y && apt-get clean && \
  rm -rf /gonex

EXPOSE 8545
EXPOSE 30303

ENTRYPOINT ["/nexty"]
