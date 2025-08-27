# builder
FROM golang:1.21-alpine AS builder
WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compilar la app (usa main.go de la raíz)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /bin/robotapp main.go

# runtime
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /bin/robotapp /bin/robotapp
WORKDIR /app
EXPOSE 8080
ENV APP_PORT=8080
# CMD ["/bin/robotapp"]
