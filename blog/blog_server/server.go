package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kenriortega/go-grpc-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	DB    = "blogDb"
	BLOGS = "blogs"
)

var (
	collection *mongo.Collection
)

type server struct{}

// DTO

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// db
	fmt.Println("Conecting to mongodb")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

		panic(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)

		panic(err)
	}
	collection = client.Database(DB).Collection(BLOGS)
	// server grpc

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	creds, sslErr := credentials.NewServerTLSFromFile("./ssl/server.crt", "./ssl/server.pem")
	if sslErr != nil {
		log.Fatalf("Failed to parse credentials: %v", sslErr)
		return
	}
	opts := grpc.Creds(creds)
	s := grpc.NewServer(opts)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Blog service started")

		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to server: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// block until a signal is received
	<-ch
	fmt.Println("Stoping server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Closing the mongoDB connection")
	client.Disconnect(context.TODO())
	fmt.Println("program was close")
}
