package mongo

import "go.mongodb.org/mongo-driver/mongo"

var Collection *mongo.Collection

func TextCollection(database *mongo.Database) {
	Collection = database.Collection("texts")
}

