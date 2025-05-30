package config

import "os"

type Config struct {
	MongoURI  string
	JWTSecret string
	NATSURL   string
	GRPCPort  string
}

func Load() Config {
	return Config{
		MongoURI:  getEnv("MONGO_URI", "mongodb://localhost:27017"),
		JWTSecret: getEnv("JWT_SECRET", "secret"),
		NATSURL:   getEnv("NATS_URL", "nats://localhost:4222"),
		GRPCPort:  getEnv("GRPC_PORT", "50052"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
