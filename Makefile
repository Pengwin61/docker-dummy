.PHONY: build

help:
	echo "Hello, help"

build:
	go build -o dummy cmd/main.go

build-docker:
    # make build-docker version=0.0.3
	docker build -t pengwin61/dummy:$(version) .
build-docker-arm64:
	docker buildx build --platform linux/arm64 -t pengwin61/dummy:$(version)-arm64 .
build-docker-latest:
	docker build -t pengwin61/dummy:latest .
build-docker-latest-arm64:
	docker build -t pengwin61/dummy:latest-arm64 .

push-docker:
	docker push pengwin61/dummy:$(version)
push-docker-arm64:
	docker push pengwin61/dummy:$(version)-arm64
push-docker-latest:
	docker push pengwin61/dummy:latest



