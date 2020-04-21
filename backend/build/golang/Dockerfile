FROM golang:latest

RUN mkdir /go/src/work

WORKDIR /go/src/work

RUN go get github.com/gorilla/mux
 
RUN go get github.com/gorilla/handlers

RUN go get github.com/caarlos0/env \
 && go get xorm.io/xorm \
 && go get github.com/lib/pq

ADD . /go/src/work