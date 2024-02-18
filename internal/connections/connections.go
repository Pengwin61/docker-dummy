package connections

import (
	"docker-dummy/internal/config"
	"docker-dummy/internal/rabbitmq"
	"docker-dummy/internal/redis"
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

	if config.Redis.Enable {
		rc := redis.NewRedisClient(config.Redis.Host, config.Redis.Port, config.Redis.Pass, config.Redis.Db)
		return &Connections{Redis: rc, RabbitMQ: nil}
	}
	if config.RabbitMQ.Enable {
		rabbitc := rabbitmq.NewRabbitClient(config)
		return &Connections{Redis: nil, RabbitMQ: rabbitc}
	}

	return nil
}
