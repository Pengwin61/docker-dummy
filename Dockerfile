FROM golang:alpine AS builder
WORKDIR /build
ADD go.mod .
COPY . .
RUN go build -o dummy cmd/main.go



FROM alpine
WORKDIR /build
COPY --from=builder /build/dummy /build/dummy
CMD [". /dummy"]