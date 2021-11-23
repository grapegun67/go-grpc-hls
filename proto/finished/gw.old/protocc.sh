#protoc 컴파일 명령어
protoc -I . -I googleapis/ --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative test_gw.proto
