package web

import "net"

type response struct {
	TimeStamp       string          `json:"timestamp"`
	System          System          `json:"system"`
	ExternalService ExternalService `json:"externalService"`
}

type ExternalService struct {
	Redis    Redis    `json:"redis"`
	Database Database `json:"database"`
	Rabbit   Rabbit   `json:"rabbit"`
}
type System struct {
	Hostname string `json:"hostname"`
	Ip       net.IP `json:"ip"`
}

type Redis struct {
	Hostname      string        `json:"redishostname"`
	Status        string        `json:"redisstatus"`
	RedisResponse RedisResponse `json:"redisresponse"`
}
type RedisResponse struct {
	Name    string `json:"redisname"`
	Surname string `json:"redissurname"`
}

type Database struct {
	Hostname   string     `json:"dbname"`
	Status     string     `json:"dbstatus"`
	DbResponse DbResponse `json:"dbresponse"`
}
type DbResponse struct {
	Name    string `json:"dbname"`
	Surname string `json:"dbsurname"`
}

type Rabbit struct {
	Hostname       string         `json:"rabbithostname"`
	Status         string         `json:"rabbitstatus"`
	RabbitResponse RabbitResponse `json:"rabbitresponse"`
}
type RabbitResponse struct {
	Msg []string `json:"rabbitmsg"`
}
