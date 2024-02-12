package connections

import "docker-dummy/internal/redis"

type Connections struct {
	Redis *redis.ClientRedis
}

var Con *Connections

func InitAllConnections(ip string, pass string, db int) {
	Con = getAllConnections(ip, pass, db)
}

func getAllConnections(ip string, pass string, db int) *Connections {

	rc := redis.NewRedisClient(ip, pass, db)

	return &Connections{Redis: rc}
}
