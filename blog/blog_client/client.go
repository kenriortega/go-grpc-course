package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kenriortega/go-grpc-course/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, sslErr := credentials.NewClientTLSFromFile("./ssl/ca.crt", "")
	if sslErr != nil {
		log.Fatalf("Failed to parse credentials: %v", sslErr)
		return
	}
	cc, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)
	createBlog(c)
}

func createBlog(c blogpb.BlogServiceClient) {
	blog := &blogpb.Blog{
		AuthorId: "SAs",
		Title:    "My first",
		Content:  "QSd dfe dfdf e",
	}
	res, err := c.CreateBlog(context.Background(),
		&blogpb.CreateBlogRequest{Blog: blog},
	)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(res.Blog.Id)
}
