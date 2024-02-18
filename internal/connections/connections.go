package connections

import (
	"docker-dummy/internal/config"
	"docker-dummy/internal/rabbitmq"
	"docker-dummy/internal/redis"
	"log"
)

type Connections struct {
	Redis    *redis.ClientRedis
	RabbitMQ *rabbitmq.ClientRabbit
}

var Con *Connections

func InitAllConnections(config *config.Config) {
	Con = getAllConnections(config)
}

func getAllConnections(config *config.Config) *Connections {
	var rc *redis.ClientRedis
	var rabbitc *rabbitmq.ClientRabbit

	if config.Redis.Enable {
		rc = redis.NewRedisClient(config.Redis.Host, config.Redis.Port, config.Redis.Pass, config.Redis.Db)
	} else {
		log.Println("i can't connect to the redis host because the enable parameter is false")
	}

	if config.RabbitMQ.Enable {
		rabbitc = rabbitmq.NewRabbitClient(config)
	} else {
		log.Println("i can't connect to the redis host because the enable parameter is false")
	}

	return &Connections{Redis: rc, RabbitMQ: rabbitc}
}
