

step1: run command on terminal
=> go get github.com/golang/protobuf/proto
=> go get -u google.golang.org/grpc

step2: download protoc and configure

step3: to generate
=> protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

step4: run 
=> go run .


