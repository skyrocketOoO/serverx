FROM golang:latest
WORKDIR /builder
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server app/server.go


FROM alpine:latest
WORKDIR /app
COPY --from=0 /builder/server .
COPY config/config.yaml ./config/config.yaml
CMD ["./server"]