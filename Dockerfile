FROM docker.io/golang:1.21

WORKDIR /usr/src/app

COPY ./backend .
RUN go mod download && go mod verify

RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
