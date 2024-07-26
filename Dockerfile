FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o stress-test cmd/main.go

FROM scratch AS runner
COPY --from=builder /app/stress-test .
ENTRYPOINT ["./stress-test"]