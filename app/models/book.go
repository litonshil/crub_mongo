package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty" validate:"required"`
	Author      string             `json:"author,omitempty" validate:"required"`
	Publication string             `json:"publication,omitempty" validate:"required"`
}
