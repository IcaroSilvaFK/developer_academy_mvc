build:
	GOOS=linux GOARCH=arm64 go build -o build/main cmd/main.go
test:
	go test ./...
