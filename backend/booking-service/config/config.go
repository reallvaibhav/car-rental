package config

type Config struct {
	MongoURI      string
	RedisAddr     string
	NatsURL       string
	ServicePort   string
	MongoDatabase string
}

func NewConfig() *Config {
	return &Config{
		MongoURI:      "mongodb://localhost:27017",
		RedisAddr:     "localhost:6379",
		NatsURL:       "nats://localhost:4222",
		ServicePort:   ":50051",
		MongoDatabase: "bookingdb",
	}
}
