# docker-dummy


## Приложение для тестирования отказоустойчивости компонентов в kubernetes

### DockerHub

```bash
docker pull pengwin61/dummy:latest
```

### Build через Makefile

```bash
make build
```

### Flugs application
``` ./dummy --help```
```
      --client                Run client app
      --rabbit-enable         Run rabbit
      --rabbit-host string    RabbitMQ host (default "127.0.0.1")
      --rabbit-pass string    RabbitMQ password (default "guest")
      --rabbit-port string    RabbitMQ port (default "5672")
      --rabbit-queue string   Name Queue (default "q1")
      --rabbit-user string    RabbitMQ user (default "guest")
      --redis-db int          Redis database
      --redis-enable          Run redis
      --redis-host string     Redis host (default "127.0.0.1")
      --redis-pass string     Redis password
      --redis-port string     Redis port (default "6379")
      --server                Run server app
```
### Server
Сервер читает данные

### Client
Клиент отсылает сообщения в redis
