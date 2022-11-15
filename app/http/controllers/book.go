package controllers

import (
	"crud_mongo/app/domain"
	"crud_mongo/app/serializers"
	"crud_mongo/app/utils/msgutil"
	logger "crud_mongo/infra/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookCnt struct {
	bookSvc domain.IBookSvc
}

func NewBookController(bookSvc domain.IBookSvc) *BookCnt {
	costc := &BookCnt{
		bookSvc: bookSvc,
	}
	return costc
}

func (bc *BookCnt) InsertBook(c echo.Context) error {
	var req serializers.BookPayload
	var err error

	if err = c.Bind(&req); err != nil {
		logger.Error(err)
		return c.JSON(http.StatusBadRequest, msgutil.RequestBodyParseErrorResponseMsg())
	}

	err = bc.bookSvc.InsertBook(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("User"))
	}

	return c.NoContent(http.StatusCreated)
}
