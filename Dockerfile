FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=1 go build -ldflags="-w -s" -o server cmd/gobook/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]

