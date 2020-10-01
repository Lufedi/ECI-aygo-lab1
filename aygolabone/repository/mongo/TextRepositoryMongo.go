package mongo

import (
	"aygolabone/model"
	"aygolabone/repository"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type TextRepositoryMongo struct {
	
}

func NewTextRepositoryMongo() repository.TextRepository {
	return &TextRepositoryMongo{}
}

func (repo TextRepositoryMongo) GetRecent() ([]*model.Text, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	 queryOptions := options.Find().SetLimit(10).SetSort(bson.D{{"createddate", -1}})

	cursor, err := Collection.Find(context.Background(), bson.D{}, queryOptions)
	if err != nil {
		log.Fatal("An error ocurrer obtaining the last texts", err)
	}
	var result []*model.Text
	for cursor.Next(ctx){
		var text *model.Text
		if err := cursor.Decode(&text); err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		result = append(result, text)
	}
	return result, nil
}


func (repo TextRepositoryMongo) Save(text *model.Text) (*model.Text, error) {
	text.CreatedDate = time.Now();
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Collection.InsertOne(ctx, &text)
	if err != nil {
		log.Fatal("Error saving the text", err)
	}
	return text, nil
}
