GOSERVER=server/main.go
GOCLIENT=client/main.go

.PHONY: all server client

all: server client

deps:
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go get -u github.com/golang/protobuf/protoc-gen-go

# If this fails add $GOPATH/bin to path
api/api.pb.go: api/mtgcard.proto
	protoc -I api/ \
		-I${GOPATH}/src \
		-I/usr/local/include \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		--go_out=plugins=grpc:api \
		api/mtgcard.proto 

api/api.pb.gw.go: api/mtgcard.proto
	protoc -I api/ \
		-I${GOPATH}/src \
		-I/usr/local/include \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		--grpc-gateway_out=logtostderr=true:api \
		api/mtgcard.proto 

api/api.swagger.json: api/mtgcard.proto
	protoc -I api/ \
		-I${GOPATH}/src \
		-I/usr/local/include \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
		--swagger_out=logtostderr=true:api \
		api/mtgcard.proto

api: api/api.pb.go api/api.pb.gw.go api/api.swagger.json

server: api
	mkdir -p bin
	go build -i -v -o bin/server ${GOSERVER}

client: api
	mkdir -p bin
	go build -i -v -o bin/client ${GOCLIENT}

certs:
	mkdir -p cert
	openssl genrsa -out cert/server.key 2048
	openssl req -new -x509 -sha256 -key cert/server.key -out cert/server.crt -days 3650
	openssl req -new -sha256 -key cert/server.key -out cert/server.csr
	openssl x509 -req -sha256 -in cert/server.csr -signkey cert/server.key -out cert/server.crt -days 3650