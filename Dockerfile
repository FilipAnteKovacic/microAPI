FROM golang:1.10.1

WORKDIR /go/src/app/
COPY . /go/src/app/

RUN go get github.com/globalsign/mgo
RUN go get github.com/globalsign/mgo/bson

RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/handlers

RUN go get github.com/FilipAnteKovacic/microapi/crud/

RUN go build
RUN go install

EXPOSE 80

CMD ["app"]