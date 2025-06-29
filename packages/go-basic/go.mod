module packages/go-basic

go 1.23.6

require (
	api/go-sdk v0.0.0-00010101000000-000000000000
	connectrpc.com/connect v1.18.1
)

require google.golang.org/protobuf v1.34.2 // indirect

replace api/go-sdk => ../../api/go-sdk
