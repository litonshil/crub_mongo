package svc

import (
	"crud_mongo/app/domain"
	"crud_mongo/app/models"
	"crud_mongo/app/serializers"
	"crud_mongo/app/utils/errutil"
	methods "crud_mongo/app/utils/methodutil"
	logger "crud_mongo/infra/logger"
	"fmt"
)

type books struct {
	bookRepo domain.IBookRepo
}

func NewCostService(bookRepo domain.IBookRepo) domain.IBookSvc {
	return &books{
		bookRepo: bookRepo,
	}
}

func (cs *books) InsertBook(req *serializers.BookPayload) error {
	book := &models.Book{}
	respErr := methods.CopyStruct(req, &book)
	if respErr != nil {
		return respErr
	}
	fmt.Println("test")
	if err := cs.bookRepo.InsertBook(book); err != nil {
		logger.Error(err)
		return errutil.ErrBookCreate
	}

	return nil
}
