FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux CGO_ENABLE=0 go build -ldflags="-w -s" -o main ./cmd/main.go

CMD ["./main"]

# FROM scratch

# COPY --from=builder /app/main .

# CMD ["./main"]
