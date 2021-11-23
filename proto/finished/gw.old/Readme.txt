[test_gw.proto 컴파일 과정]
- 절차가 중요하니 아래의 순서를 따를것

1. gateway 옵션, proto 소스를 뺀 일반적인 .proto파일을 컴파일한다(-l 제거)
	-> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative test_gw.proto
2. gateway 옵션, proto 소스를 넣고 수정한 다음에 .proto파일을 컴파일한다
	-> protoc -I . -I googleapis/ --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative test_gw.proto

