FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

EXPOSE 9009

CMD ["./main"]