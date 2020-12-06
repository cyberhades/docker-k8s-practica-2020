FROM golang:1.15 as build

WORKDIR /go/src/app
COPY main.go .

RUN go get -d -v ./... 
RUN go install -v ./...

EXPOSE 8080

CMD ["/go/bin/app"]
