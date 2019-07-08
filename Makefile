GOSERVER=server/main.go
GOCLIENT=client/main.go

.PHONY: all server client

all: server client

api:
	protoc -I api/ api/mtgcard.proto --go_out=plugins=grpc:api

server: api
	go build -i -v -o bin/server ${GOSERVER}

client: api
	go build -i -v -o bin/client ${GOCLIENT}