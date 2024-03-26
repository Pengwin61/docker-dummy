package web

import (
	"docker-dummy/internal/connections"
	"log"
)

var msg []string

func ReadInQueue() {

	if connections.Con.RabbitMQ != nil {
		messages := connections.Con.RabbitMQ.Get(webParams.rabbitQueue)
		forever := make(chan bool)

		go func() {
			for message := range messages {
				if len(msg) < 500 {
					msg = append(msg, string(message.Body))
				} else {
					log.Println("переполнился")
					msg = nil
				}

			}
		}()
		<-forever
	}

}
