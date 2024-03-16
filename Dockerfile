FROM golang:1.21 as builder

WORKDIR /app

COPY ./ ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o aggregate cmd/aggregate/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/aggregate /aggregate

ENTRYPOINT ["/aggregate"]
