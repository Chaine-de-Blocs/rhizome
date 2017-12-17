FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN apt-get update
RUN apt-get install -y vim

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
