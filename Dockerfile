# build
FROM golang:1.13.4-alpine3.10 as builder
WORKDIR /app
COPY . /app
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o pim cmd/product/main.go
# runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app
USER app
WORKDIR /app
COPY --from=builder /app/pim .
ENTRYPOINT ["./pim"]