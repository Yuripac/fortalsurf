FROM golang:1.24.1-alpine3.21

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /fortalsurf

EXPOSE 8080

CMD ["/fortalsurf"]
