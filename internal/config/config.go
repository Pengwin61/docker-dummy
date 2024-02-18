package config

import (
	"fmt"

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
	Enable    bool
	Host      string
	Port      string
	User      string
	Pass      string
	QueueName string
}
type RedisConfig struct {
	Enable bool
	Host   string
	Port   string
	Pass   string
	Db     int
}

func GetConfig() *Config {
	cfg := initConfig()
	checkConfig(cfg)

	return cfg
}

func initConfig() *Config {

	server := pflag.Bool("server", false, "Run server app")
	client := pflag.Bool("client", false, "Run client app")
	redis := pflag.Bool("redis-enable", false, "Run redis")
	rabbit := pflag.Bool("rabbit-enable", false, "Run rabbit")

	redisHost := pflag.String("redis-host", "127.0.0.1", "Redis host")
	redisPort := pflag.String("redis-port", "6379", "Redis port")
	redisPass := pflag.String("redis-pass", "", "Redis password")
	redisDb := pflag.Int("redis-db", 0, "Redis database")

	rabbitHost := pflag.String("rabbit-host", "127.0.0.1", "RabbitMQ host")
	rabbitPort := pflag.String("rabbit-port", "5672", "RabbitMQ port")
	rabbitUser := pflag.String("rabbit-user", "guest", "RabbitMQ user")
	rabbitPass := pflag.String("rabbit-pass", "guest", "RabbitMQ password")
	rabbitQueue := pflag.String("rabbit-queue", "q1", "Name Queue")

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
			Enable:    *rabbit,
			Host:      *rabbitHost,
			Port:      *rabbitPort,
			User:      *rabbitUser,
			Pass:      *rabbitPass,
			QueueName: *rabbitQueue,
		}}
}

func checkConfig(cfg *Config) {
	checkRedis(cfg)
	checkRabbit(cfg)
}

func checkRedis(cfg *Config) {
	if cfg.Redis.Host != "127.0.0.1" {
		fmt.Println("redis host is change")
		cfg.Redis.Enable = true
	}
	if cfg.Redis.Port != "6379" {
		fmt.Println("redis port is change")
		cfg.Redis.Enable = true
	}
	if cfg.Redis.Pass != "" {
		fmt.Println("redis pass is change")
		cfg.Redis.Enable = true
	}
	if cfg.Redis.Db != 0 {
		fmt.Println("redis db is change")
		cfg.Redis.Enable = true
	}
}

func checkRabbit(cfg *Config) {

	if cfg.RabbitMQ.Host != "127.0.0.1" {
		fmt.Println("rabbit host is change")
		cfg.RabbitMQ.Enable = true
	}
	if cfg.RabbitMQ.Port != "5672" {
		fmt.Println("rabbit port is change")
		cfg.RabbitMQ.Enable = true
	}
	if cfg.RabbitMQ.User != "guest" {
		fmt.Println("rabbit user is change")
		cfg.RabbitMQ.Enable = true
	}
	if cfg.RabbitMQ.Pass != "guest" {
		fmt.Println("rabbit pass is change")
		cfg.RabbitMQ.Enable = true
	}
	if cfg.RabbitMQ.QueueName != "q1" {
		fmt.Println("rabbit queue is change")
		cfg.RabbitMQ.Enable = true
	}
}
