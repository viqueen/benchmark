FROM golang:1.23.1 AS go

RUN go install github.com/viqueen/protoc-gen-sqlc/protoc-gen-sqlc@v0.0.6

FROM bufbuild/buf:1.49.0 AS buf

COPY --from=go /go/bin/protoc-gen-sqlc /usr/local/bin/protoc-gen-sqlc
ENV PATH="/usr/local/bin:${PATH}"

ENTRYPOINT ["buf"]