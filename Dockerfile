FROM golang:1.10-alpine

RUN apk update && apk add --no-cache \
  gcc \
  git \
  libc-dev \
  make \
  sqlite

WORKDIR /opt

RUN git clone https://github.com/chadwickbureau/retrosheet.git \
  && git clone https://github.com/chadwickbureau/baseballdatabank.git

RUN go get -u github.com/golang/lint/golint \
  && rm -rf $GOPATH/src/* \
  && rm -rf $GOPATH/pkg/*

COPY . /go/src/github.com/rippinrobr/baseball-stats-db/

WORKDIR /go/src/github.com/rippinrobr/baseball-stats-db

RUN make get \
  && make all

