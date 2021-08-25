package main

import (
	"context"
	"fmt"
	"io"
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
	fmt.Println("Create new Blog")
	// createBlog(c)
	fmt.Println("Read Blog")
	// Read Blog by ID
	// readBlog(c)
	fmt.Println("Update Blog")
	// update blog
	// updateBlog(c)

	listStreamBlog(c)

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

func readBlog(c blogpb.BlogServiceClient) {
	blogId := "612667705c5c154d44f5560b"
	res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: blogId,
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(res.Blog)
}

func updateBlog(c blogpb.BlogServiceClient) {
	blogId := "612667705c5c154d44f5560b"
	blog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "SAs",
		Title:    "Update title",
		Content:  "QSd dfe dfdf e",
	}
	res, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: blog,
	})

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(res.Blog)
}
func listStreamBlog(c blogpb.BlogServiceClient) {

	stream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.GetBlog())
	}
}
