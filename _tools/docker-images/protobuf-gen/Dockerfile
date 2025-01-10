FROM golang:1.23.1 AS go

RUN go install github.com/viqueen/protoc-gen-sqlc/protoc-gen-sqlc@v0.0.6 \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.2 \
    && go install connectrpc.com/connect/cmd/protoc-gen-connect-go@v1.18.1

FROM bufbuild/buf:1.49.0 AS buf

COPY --from=go /go/bin/protoc-gen-sqlc /usr/local/bin/protoc-gen-sqlc
COPY --from=go /go/bin/protoc-gen-go /usr/local/bin/protoc-gen-go
COPY --from=go /go/bin/protoc-gen-connect-go /usr/local/bin/protoc-gen-connect-go
ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["buf"]