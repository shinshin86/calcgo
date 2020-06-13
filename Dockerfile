FROM golang:latest

ENV SRC_DIR=/go/src/work
ENV GO_BIN=/go/bin

RUN mkdir $SRC_DIR
WORKDIR $SRC_DIR
ADD . $SRC_DIR

RUN go get -u

