FROM golang:1.23.6-alpine AS builder

WORKDIR /go/src/packages/go-basic

RUN apk add --no-cache git ca-certificates

COPY packages/go-basic/go.mod packages/go-basic/go.sum ./
COPY packages/go-basic/ .
COPY api/go-sdk ../../api/go-sdk

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./cmd/main.go

FROM alpine:3.19

WORKDIR /root/
COPY --from=builder /go/src/packages/go-basic/server .

EXPOSE 8080
CMD ["./server"]