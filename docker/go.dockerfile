FROM golang:1.20.3

WORKDIR /source

RUN go install github.com/codegangsta/gin@latest 

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
