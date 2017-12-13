FROM golang:latest
MAINTAINER Ana Paula Gomes "apgomes88@gmail.com"

RUN mkdir /go-tunefy
WORKDIR /go-tunefy

ENV SRC_DIR=/go/src/github.com/anapaulagomes/go-tunefy
ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o gotunefy; cp gotunefy /go-tunefy/

ENTRYPOINT ["./gotunefy"]
