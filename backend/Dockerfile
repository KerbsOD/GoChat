FROM golang:1.21

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /gochat

EXPOSE 8080

CMD [ "/gochat" ]
