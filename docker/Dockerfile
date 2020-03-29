FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

RUN go get github.com/gorilla/mux \
 && go get github.com/gorilla/handlers

ADD . /go/src/work