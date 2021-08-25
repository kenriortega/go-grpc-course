package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/kenriortega/go-grpc-course/blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	DB    = "blogDb"
	BLOGS = "blogs"
)

var (
	collection *mongo.Collection
)

type server struct{}

func (*server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()
	data := blogItem{
		ID:       primitive.NewObjectID(),
		AuthorID: blog.GetAuthorId(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}
	result, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		log.Fatal(err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert oid error: %v", err),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (*server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogIdHex := req.GetBlogId()
	data := &blogItem{}
	blogIdFromHex, err := primitive.ObjectIDFromHex(blogIdHex)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert blogid error: %v", err),
		)
	}

	filter := bson.D{primitive.E{Key: "_id", Value: blogIdFromHex}}
	err = collection.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog by OID error: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (*server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	blog := req.GetBlog()
	blogIdFromHex, err := primitive.ObjectIDFromHex(blog.GetId())
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert blogid error: %v", err),
		)
	}
	filter := bson.D{primitive.E{Key: "_id", Value: blogIdFromHex}}
	updater := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "title", Value: blog.GetTitle()},
	}}}
	result, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot find blog by OID error: %v", err),
		)
	}
	if result.ModifiedCount == 1 {

		return &blogpb.UpdateBlogResponse{
			Blog: &blogpb.Blog{
				Id:       blog.GetId(),
				AuthorId: blog.GetAuthorId(),
				Title:    blog.GetTitle(),
				Content:  blog.GetContent(),
			},
		}, nil
	} else {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot update blog by OID error: %v", blog),
		)
	}

}

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
