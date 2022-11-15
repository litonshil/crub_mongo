package repository

import (
	"context"
	"crud_mongo/app/domain"
	"crud_mongo/app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type books struct {
	DB *mongo.Client
}

func NewBookRepository(dbc *mongo.Client) domain.IBookRepo {
	return &books{
		DB: dbc,
	}
}

func (cr *books) InsertBook(req *models.Book) error {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := cr.DB.Database("test").Collection("book")
	newBook := models.Book{
		Id:          primitive.NewObjectID(),
		Title:       req.Title,
		Author:      req.Author,
		Publication: req.Publication,
	}
	_, err := collection.InsertOne(ctx, newBook)
	if err != nil {
		return err
	}
	return nil
}
