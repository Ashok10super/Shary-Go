package main

import (
	"context"
	"fmt"
	"log"

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