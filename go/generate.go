package main

//go:generate rm -rf build/protoc
//go:generate mkdir -p build/protoc
//go:generate protoc --proto_path ../protobuf --go_out=build/protoc --twirp_out=build/protoc ../protobuf/service.proto
//go:generate rm -rf rpc
//go:generate cp -r build/protoc/github.com/ngyewch/twirp-playground/rpc .
//go:generate rm -rf build/protoc
