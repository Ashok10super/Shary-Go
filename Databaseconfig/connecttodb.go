package databaseconfig

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	post "shary/Model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connecttodb() *mongo.Client {
	//setting the client object for connection
	clientoptions:= options.Client().ApplyURI("mongodb://localhost:27017")

	//connect to the db using client object
	client,err:= mongo.Connect(context.TODO(),clientoptions)

	if err!=nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(),nil)

	if err!=nil{
		log.Fatal(err)
	}else{
	  fmt.Printf("connection is successfully initiated")
	}
	return client
}

func Disconnectdb(cli *mongo.Client) string{
	
//disconnect the db using client

err:= cli.Disconnect(context.TODO())
   
  if err!=nil {
	log.Fatal(err)
  }
  return "disconnection successfull"

}

func Getallposts(client *mongo.Client){

	collection := client.Database("Post").Collection("userposts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    defer cancel()

    // Fetch all posts
    var Post []post.Post

    filter := bson.M{} // Empty filter to fetch all documents
    cursor, err := collection.Find(ctx, filter)
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)
 for cursor.Next(ctx) {
        var post post
        if err := cursor.Decode(&post); err != nil {
            log.Fatal(err)
        }
        post = append(Post, post)
    }

    // Check for cursor iteration errors
    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    // Print all posts in JSON format
    postsJSON, err := json.MarshalIndent(Post, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(postsJSON))
}