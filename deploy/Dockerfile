FROM golang:latest
WORKDIR /builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=0 /builder/server .
# COPY .env .
CMD ["./server"]