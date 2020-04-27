FROM golang:latest

RUN mkdir /go/src/y_ara

WORKDIR /go/src/y_ara

RUN go get github.com/gorilla/mux
 
RUN go get github.com/gorilla/handlers

RUN go get github.com/caarlos0/env \
 && go get xorm.io/xorm \
 && go get github.com/lib/pq