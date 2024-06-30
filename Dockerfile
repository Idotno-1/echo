FROM golang:latest

WORKDIR /srv/echo

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
