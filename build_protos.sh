export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
protoc -I api/ api/mtgcard.proto --go_out=plugins=grpc:api