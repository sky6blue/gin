run:
	go run ./cmd/main.go

build:
	go build ./cmd/main.go

test:
	go test -cover ./internal/...