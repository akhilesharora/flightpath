### Prepare the base image
FROM golang:1.18 AS builder

ADD . /app
WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o flightpath cmd/main.go

### Build the target image
FROM scratch
EXPOSE 8080/tcp

COPY --from=builder /app/flightpath .

CMD ["./flightpath"]
