module github.com/Car-Rental/backend/user-service

go 1.24.1

require (
	github.com/Car-Rental/proto v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.5
	github.com/golang-jwt/jwt/v4 v4.5.2
	github.com/gorilla/mux v1.8.1
	github.com/nats-io/nats.go v1.42.0
	go.mongodb.org/mongo-driver v1.17.3
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.38.0
	google.golang.org/grpc v1.72.1
	proto v0.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)

replace proto => ./proto

replace github.com/Car-Rental/proto => ../proto
