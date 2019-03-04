package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gsafcik/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Geoff",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected Error: %v", err)
	}
	fmt.Printf("Blog created: %v\n", createBlogRes)
	blogID := createBlogRes.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")
	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "5c7da4123ddc9B90aa3076e6"})
	if err2 != nil {
		fmt.Printf("Error happended while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happended while reading: %v\n", readBlogErr)
	}

	fmt.Printf("Blog read: %v\n", readBlogRes)

	// update blog
	fmt.Println("Update the blog")
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Changed Author",
		Title:    "My First Blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if updateErr != nil {
		fmt.Printf("Error happended while updating: %v\n", updateErr)
	}
	fmt.Printf("Blog updated: %v\n", updateRes)

	// update blog
	fmt.Println("Delete blog")
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if deleteErr != nil {
		fmt.Printf("Error happended while deleting: %v\n", deleteErr)
	}
	fmt.Printf("Blog deleted: %v\n", deleteRes)

}
