module e2e-tests

go 1.23.6

require (
	api/go-sdk v0.0.0-00010101000000-000000000000
	connectrpc.com/connect v1.18.1
	github.com/go-faker/faker/v4 v4.6.1
	github.com/stretchr/testify v1.10.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace api/go-sdk => ../api/go-sdk
