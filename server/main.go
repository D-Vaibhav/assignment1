package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	blogpb "github.com/vaibhav/assignment1/proto"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var db *mongo.Client
var blogdb *mongo.Collection
var mongoCtx context.Context

type BlogServiceServer struct{}

// In the function bodies we’ll generally use the following workflow:
// Protbuf Message (Request) → Regular Go Struct → Convert to BSON + Mongo Action → Protobuf Message (Response)

func (s *BlogServiceServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogReq) (*blogpb.CreateBlogRes, error) {
	//  First we’ll extract the Blog message from our request message and convert it to a regular go struct
	blog := req.GetBlog()

	// convert it into BlogItem type to convert into BSON
	data := BlogItem{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	// insert data into db,  result contains the newly generated Object ID for the new document
	result, err := blogdb.InsertOne(mongoCtx, data)
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	// adding id to the blog, , first cast the "generic type" (go doesn't have real generics yet) to an Object ID.
	oid := result.InsertedID.(primitive.ObjectID)

	// convert the id to its string counterpart
	blog.Id = oid.Hex()

	return &blogpb.CreateBlogRes{Blog: blog}, nil
}

// The MongoDB FindOne() methods takes in a context and a filter, which is a BSON document for which to filter by its keys
func (s *BlogServiceServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogReq) (*blogpb.ReadBlogRes, error) {
	blogId := req.GetId()

	// converting string blogId from pb to mongoDB ObjectId
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	blogFromDB := blogdb.FindOne(ctx, bson.M{"_id": oid})

	// creating an empty BlogItem, then write our decoded blogFromDB to emptyBlog
	data := BlogItem{}

	err = blogFromDB.Decode(&data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object Id %s:%v", blogId, err))
	}

	// cast it to ReadBlogRes type
	response := &blogpb.ReadBlogRes{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}

	return response, nil
}

func (s *BlogServiceServer) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogReq) (*blogpb.UpdateBlogRes, error) {
	blog := req.GetBlog()

	id := blog.GetId() // as string

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert the supplied blog id to a MongoDB ObjectId: %v", err))
	}

	// convert the data to be updated into an unordered Bson document
	update := bson.M{
		"author_id": blog.GetAuthorId(),
		"title":     blog.GetTitle(),
		"content":   blog.GetContent(),
	}

	// convert oid also to an unordered bson data, it'll be same as previous
	filter := bson.M{"_id": oid}

	// result is the BSON encoded result
	// To return the updated document instead of original we have to add options.
	result := blogdb.FindOneAndUpdate(ctx, filter, bson.M{"$set": update}, options.FindOneAndUpdate().SetReturnDocument(1))

	// decode the result
	data := BlogItem{}
	err = result.Decode(&data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Supplied ID %s: %v", id, err))
	}

	response := &blogpb.UpdateBlogRes{
		Blog: &blogpb.Blog{Id: data.ID.Hex(), AuthorId: data.AuthorID, Title: data.Title, Content: data.Content},
	}
	return response, nil
}

// For DeleteBlog we’ll use the mongodb DeleteOne() method, which takes in an Object ID of the document to remove.
func (s *BlogServiceServer) DeleteBlog(ctx context.Context, req *blogpb.DeleteBlogReq) (*blogpb.DeleteBlogRes, error) {
	idAsString := req.GetId()

	oid, err := primitive.ObjectIDFromHex(idAsString)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}

	// we're returning boolean not BlogItem
	_, err = blogdb.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Couldn't find/delete blog with id %s : %v", idAsString, err))
	}

	return &blogpb.DeleteBlogRes{Success: true}, nil
}

// func (s *BlogServiceServer) ListBlogs(req *blogpb.ListBlogsReq, stream blogpb.BlogService_ListBlogsServer) error {
// 	// Initiate a BlogItem type to write decoded data to
// 	data := &BlogItem{}
// 	// collection.Find returns a cursor for our (empty) query
// 	cursor, err := blogdb.Find(context.Background(), bson.M{})
// 	if err != nil {
// 		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
// 	}
// 	// An expression with defer will be called at the end of the function
// 	defer cursor.Close(context.Background())
// 	// cursor.Next() returns a boolean, if false there are no more items and loop will break
// 	for cursor.Next(context.Background()) {
// 		// Decode the data at the current pointer and write it to data
// 		err := cursor.Decode(data)
// 		// check error
// 		if err != nil {
// 			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
// 		}
// 		// If no error is found send blog over stream
// 		stream.Send(&blogpb.ListBlogsRes{
// 			Blog: &blogpb.Blog{
// 				Id:       data.ID.Hex(),
// 				AuthorId: data.AuthorID,
// 				Content:  data.Content,
// 				Title:    data.Title,
// 			},
// 		})
// 	}
// 	// Check if the cursor has any errors
// 	if err := cursor.Err(); err != nil {
// 		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
// 	}
// 	return nil
// }

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {
	// configure log package to produce line number if in case og log.Fatalf(), (log.LstdFLags = log.Ldate | log.Ltime)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Starting server on port : 4000...")

	// start our listner
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatalf("Failed to listen to the server port: 4000 , %v", err)
	}

	// creating a new grpcServer with blank opts
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	srv := &BlogServiceServer{}

	// registering the microservice with grpc server
	blogpb.RegisterBlogServiceServer(grpcServer, srv)

	// INITIALIZE MONGODB CLIENT
	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background() // non nil empty context

	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("Failed to connect to mongo: ", err)
	}

	// check for successful conection by pinging to MongoDB server
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	log.Printf("Connected to MongoDB.!")

	// binding our connection to our Global variable so to be used in other method
	blogdb = db.Database("mydb").Collection("blog")

	// STARTING SERVER IN CHILD GOROUTE
	go func() {
		err := grpcServer.Serve(listener)
		if err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("Server  started successfully on port:4000")

	// server SHUTDOWN hook to stop server properly
	shutdownSignalChannel := make(chan os.Signal)
	signal.Notify(shutdownSignalChannel, os.Kill)
	signal.Notify(shutdownSignalChannel, os.Interrupt)

	_ = <-shutdownSignalChannel

	// after recieveing shutdownSignal
	fmt.Println("\n Stopping the server...")
	grpcServer.Stop()
	err = listener.Close()
	if err != nil {
		log.Println("Listener not close properly")
	}

	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("All command executed, done.!")
}
