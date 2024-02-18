.PHONY: build

help:
	echo "Hello, help"

build:
	go build -o dummy cmd/main.go

build-docker:
    # make build-docker-version version=0.0.3
	docker build -t pengwin61/dummy:$(version) .
push-docker:
	docker push pengwin61/dummy:$(version)


build-docker-latest:
	docker build -t pengwin61/dummy:latest .

push-docker-latest:
	docker push pengwin61/dummy:latest
