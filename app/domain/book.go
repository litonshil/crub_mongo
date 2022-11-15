package domain

import (
	"crud_mongo/app/models"
	"crud_mongo/app/serializers"
)

type IBookRepo interface {
	InsertBook(req *models.Book) error
}

type IBookSvc interface {
	InsertBook(payload *serializers.BookPayload) error
}
