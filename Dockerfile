FROM golang:1.8
RUN go get -v gopkg.in/olivere/elastic.v6
ADD . /go/src/elastigo
RUN ls -l /go/src/elastigo
RUN ls -l /go/src/elastigo/elgoclient
RUN go install elastigo/elgoclient
ENTRYPOINT /go/bin/elgoclient
EXPOSE 8080
