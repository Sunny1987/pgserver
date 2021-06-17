FROM golang:latest

RUN mkdir /pgserver

COPY . /pgserver

WORKDIR /pgserver

RUN go build -o main .

CMD ["/pgserver/main"]