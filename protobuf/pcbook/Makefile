.PHONY: gen
gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=openapi proto/*.proto

.PHONY: server
server:
	go run cmd/server/main.go -port 8080

.PHONY: rest
rest:
	go run cmd/server/main.go -port 8081 -type rest

server1:
	go run cmd/server/main.go -port 50051

server2:
	go run cmd/server/main.go -port 50052

.PHONY: client1
client1:
	go run cmd/client/main.go -address 0.0.0.0:8080

.PHONY: client2
client2:
	go run cmd/client2/main.go -address 0.0.0.0:8080

.PHONY: client
client:
	go run cmd/client3/main.go -address 0.0.0.0:8080

.PHONY: client-tls
client-tls:
	go run cmd/client3/main.go -address 0.0.0.0:8080 -tls

.PHONY: cert
cert:
	cd cert; ./gen.sh