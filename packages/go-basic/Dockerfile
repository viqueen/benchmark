FROM golang:1.23.6-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/main.go

FROM alpine:3.19

WORKDIR /root/
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]