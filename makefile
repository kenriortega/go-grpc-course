generate:
	./generate.sh

greetserver:
	go run handOne/greet/greet_server/server.go
greetclient:
	go run handOne/greet/greet_client/client.go