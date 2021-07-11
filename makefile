generate:
	./generate.sh

server:
	go run handOne/greet/greet_server/server.go
client:
	go run handOne/greet/greet_client/client.go