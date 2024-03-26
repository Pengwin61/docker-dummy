package web

import "docker-dummy/internal/connections"

func checkRedis() string {

	if connections.Con.Redis != nil {
		return connections.Con.Redis.Check()
	}
	return ""
}

func getInRedis(entity string) string {
	if connections.Con.Redis != nil {
		return connections.Con.Redis.Get(entity)
	}
	return ""
}

func checkDb() string {
	if connections.Con.Postgres != nil {
		return "connected"
	}
	return ""
}
