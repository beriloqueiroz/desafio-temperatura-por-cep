FROM  golang:1.22.3-alpine3.19 as build

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -o /handle ./cmd/main.go

# FROM scratch

# COPY --from=build /handle /handle
# COPY .env .env

ENTRYPOINT ["/handle"]