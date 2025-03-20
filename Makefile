build:
	GOOS=linux GOARCH=arm64 go build -o build/main cmd/main.go
test:
	go test ./...

sync:
	browser-sync start --proxy "http://localhost:8080" --files "**/*.gotmpl"