package config

import (
	"github.com/spf13/pflag"
)

type Config struct {
	App      App
	Redis    RedisConfig
	RabbitMQ RabbitConfig
}
type App struct {
	Server bool
	Client bool
}

type RabbitConfig struct {
	Enable bool
	Host   string
	Port   string
	User   string
	Pass   string
}
type RedisConfig struct {
	Enable bool
	Host   string
	Port   string
	Pass   string
	Db     int
}

func InitConfig() *Config {

	server := pflag.Bool("server", false, "Run server app")
	client := pflag.Bool("client", false, "Run client app")
	redis := pflag.Bool("redis", false, "Run redis")
	rabbit := pflag.Bool("rabbit", false, "Run rabbit")

	redisHost := pflag.String("redis-host", "127.0.0.1", "Redis host")
	redisPort := pflag.String("redis-port", "6379", "Redis port")
	redisPass := pflag.String("redis-pass", "", "Redis password")
	redisDb := pflag.Int("redis-db", 0, "Redis database")

	rabbitHost := pflag.String("rabbit-host", "127.0.0.1", "RabbitMQ host")
	rabbitPort := pflag.String("rabbit-port", "5672", "RabbitMQ port")
	rabbitUser := pflag.String("rabbit-user", "guest", "RabbitMQ user")
	rabbitPass := pflag.String("rabbit-pass", "guest", "RabbitMQ password")

	pflag.Parse()

	return &Config{
		App: App{
			Server: *server,
			Client: *client,
		},
		Redis: RedisConfig{
			Enable: *redis,
			Host:   *redisHost,
			Port:   *redisPort,
			Pass:   *redisPass,
			Db:     *redisDb,
		},
		RabbitMQ: RabbitConfig{
			Enable: *rabbit,
			Host:   *rabbitHost,
			Port:   *rabbitPort,
			User:   *rabbitUser,
			Pass:   *rabbitPass,
		}}
}
