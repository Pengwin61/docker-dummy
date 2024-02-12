package connections

import (
	"docker-dummy/internal/rabbitmq"
	"docker-dummy/internal/redis"
)

type Connections struct {
	Redis    *redis.ClientRedis
	RabbitMQ *rabbitmq.ClientRabbit
}

var Con *Connections

func InitAllConnections(ip string, pass string, db int) {
	Con = getAllConnections(ip, pass, db)
}

func getAllConnections(ip string, pass string, db int) *Connections {

	rc := redis.NewRedisClient(ip, pass, db)

	rabbitc := rabbitmq.NewRabbitClient()

	return &Connections{Redis: rc, RabbitMQ: rabbitc}
}
