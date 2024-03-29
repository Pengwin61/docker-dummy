# docker-dummy


## Приложение для тестирования отказоустойчивости компонентов в kubernetes

### Build через Makefile

```bash
make build
```

### Docker Hub
```bash
docker pull pengwin61/dummy:latest
```
### Docker-compose
```bash
cd ./deployment && docker compose up -d --force-recreate
```

### Flugs application
``` ./dummy --help```
```
      --server                 Run server app
      --client                 Run client app

      --rabbit-enable          Run rabbit
      --rabbit-host string     RabbitMQ host (default "127.0.0.1")
      --rabbit-port string     RabbitMQ port (default "5672")
      --rabbit-user string     RabbitMQ user (default "guest")
      --rabbit-pass string     RabbitMQ password (default "guest")
      --rabbit-queue string    Name Queue (default "q1")

      --redis-enable           Run redis
      --redis-host string      Redis host (default "127.0.0.1")
      --redis-port string      Redis port (default "6379")
      --redis-pass string      Redis password
      --redis-db int           Redis database

      --db-host string         Database host (default "127.0.0.1")
      --db-port string         Database port (default "5432")
      --db-user string         Database user (default "pengwin")
      --db-pass string         Database password (default "password")
      --db-name string         Database name  (default "dummy")

```



## Server
Сервер читает данные

## WEB
```
http://localhost
```

### Run CMD
```
./dummy --server=true \
--rabbit-host="localhost" --rabbit-port="5672" \
--redis-host="localhost" \

```


## Client
Клиент отсылает сообщения в redis
