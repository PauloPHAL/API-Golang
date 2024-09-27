FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server cmd/gobook/main.go

FROM scratch
COPY --from=builder /app/server .
CMD ["./server"]

#DB_HOST=localhost
#DB_PORT=5432
#DB_NAME=achadosEperdidos
#DB_USER=postgres
#DB_PASSWORD=achadoseperdidos123