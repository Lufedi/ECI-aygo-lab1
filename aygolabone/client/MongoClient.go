package client
import (
	collections "aygolabone/repository/mongo"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func Connect() *mongo.Database {
	host := "dbmongo" //docker
	//host:= "127.0.0.1" // local
	port := 27017

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://admin:admin@%s:%d/aygo", host, port))
	client,err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	defer cancel()
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Could not connect to the database", err)
	} else {
		log.Println("Connected")
	}

	db := client.Database("aygo")
	injectDependency(db)

	return db
}

func injectDependency(database *mongo.Database) {
	collections.TextCollection(database)
}
