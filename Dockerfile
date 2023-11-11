FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod. /
RUN go mod download && mod verify

COPY flat_code .
RUN go build -o flat./main.go

CMD [".flat"]