generate:
	./generate.sh

greetserver:
	go run handOne/greet/greet_server/server.go
greetclient:
	go run handOne/greet/greet_client/client.go

calcserver:
	go run handOne/calculator/calc_server/server.go
calcclient:
	go run handOne/calculator/calc_client/client.go