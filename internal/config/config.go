package config

import (
	"fmt"

	"github.com/spf13/pflag"
)

type Config struct {
	App      App
	Redis    RedisConfig
	RabbitMQ RabbitConfig
	Database DatabaseConfig
}

type App struct {
	Server  bool
	Client  bool
	Version string
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
type DatabaseConfig struct {
	Enable bool
	Host   string
	Port   string
	User   string
	Pass   string
	DbName string
}

func GetConfig() *Config {
	cfg := initConfig()
	checkConfigs(cfg)

	return cfg
}

func initConfig() *Config {

	server := pflag.Bool("server", false, "Run server app")
	client := pflag.Bool("client", false, "Run client app")
	redis := pflag.Bool("redis-enable", false, "Run redis")
	rabbit := pflag.Bool("rabbit-enable", false, "Run rabbit")
	database := pflag.Bool("database-enable", false, "Run database")

	redisHost := pflag.String("redis-host", "127.0.0.1", "Redis host")
	redisPort := pflag.String("redis-port", "6379", "Redis port")
	redisPass := pflag.String("redis-pass", "", "Redis password")
	redisDb := pflag.Int("redis-db", 0, "Redis database")

	rabbitHost := pflag.String("rabbit-host", "127.0.0.1", "RabbitMQ host")
	rabbitPort := pflag.String("rabbit-port", "5672", "RabbitMQ port")
	rabbitUser := pflag.String("rabbit-user", "guest", "RabbitMQ user")
	rabbitPass := pflag.String("rabbit-pass", "guest", "RabbitMQ password")
	rabbitQueue := pflag.String("rabbit-queue", "q1", "Name Queue")

	dbHost := pflag.String("db-host", "127.0.0.1", "Database host")
	dbPort := pflag.String("db-port", "5432", "Database port")
	dbUser := pflag.String("db-user", "pengwin", "Database user")
	dbPass := pflag.String("db-pass", "password", "Database password")
	dbName := pflag.String("db-name", "dummy", "Database name")

	pflag.Parse()

	return &Config{
		App: App{
			Server:  *server,
			Client:  *client,
			Version: "0.0.4",
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
		},
		Database: DatabaseConfig{
			Enable: *database,
			Host:   *dbHost,
			Port:   *dbPort,
			User:   *dbUser,
			Pass:   *dbPass,
			DbName: *dbName,
		}}
}

func checkConfigs(cfg *Config) {
	checkRedisConfig(cfg)
	checkRabbitConfig(cfg)
	checkDbConfig(cfg)
}

func checkRedisConfig(cfg *Config) {
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

func checkDbConfig(cfg *Config) {
	if cfg.Database.Host != "127.0.0.1" {
		fmt.Println("database host is change")
		cfg.Database.Enable = true
	}
	if cfg.Database.Port != "5432" {
		fmt.Println("database port is change")
		cfg.Database.Enable = true
	}
	if cfg.Database.User != "pengwin" {
		fmt.Println("database user is change")
		cfg.Database.Enable = true
	}
	if cfg.Database.Pass != "password" {
		fmt.Println("database pass is change")
		cfg.Database.Enable = true
	}
	if cfg.Database.DbName != "dummy" {
		fmt.Println("database name is change")
		cfg.Database.Enable = true
	}

}

func checkRabbitConfig(cfg *Config) {

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
