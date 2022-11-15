package serializers

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookPayload struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Author      string             `json:"author"`
	Publication string             `json:"publication"`
}
