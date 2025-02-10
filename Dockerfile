FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o transaction_x ./cmd/api/main.go

FROM scratch

COPY  --from=builder /app/transaction_x /transaction_x

COPY --from=builder /app/static /static

EXPOSE 8080

CMD ["/transaction_x"]
