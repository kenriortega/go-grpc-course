generate:
	./generate.sh

greetserver:
	go run greet/greet_server/server.go
greetclient:
	go run greet/greet_client/client.go

calcserver:
	go run calculator/calc_server/server.go
calcclient:
	go run calculator/calc_client/client.go

blogserver:
	go run blog/blog_server/server.go
blogclient:
	go run blog/blog_client/client.go