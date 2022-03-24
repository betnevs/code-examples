.PHONY: gen
gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative proto/*.proto

.PHONY: server
server:
	go run cmd/server/main.go -port 8080

.PHONY: client
client:
	go run cmd/client/main.go -address 0.0.0.0:8080

.PHONY: client2
client2:
	go run cmd/client2/main.go -address 0.0.0.0:8080

.PHONY: client3
client3:
	go run cmd/client3/main.go -address 0.0.0.0:8080
