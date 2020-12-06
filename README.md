# dzai-mp3-services

`protoc -I./api/ -I./pkg/pb -I/usr/local/include -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:./pkg/pb api/**/*.proto`
