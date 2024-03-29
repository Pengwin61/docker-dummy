FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o dummy cmd/main.go



FROM alpine
WORKDIR /build
ENV GIN_MODE=release
COPY --from=builder /build/dummy /build/dummy
CMD [". /dummy"]