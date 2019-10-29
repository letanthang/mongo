package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/letanthang/mongo/sequence"
)

const (
	DBName = "go0110"
	Col    = "test_counter"
)

var Client *mongo.Client

func main() {
	// test sequence
	connect()
	testSequence()
}

func testSequence() {
	col := Client.Database(DBName).Collection(Col)
	id, err := sequence.GetNextID(col, "student-seq")
	fmt.Println(id, err)
}

func connect() {
	uri := "mongodb://mongoadmin:secret@localhost:27017"
	fmt.Println("Connect MongoDb with uri", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Client = client
}
