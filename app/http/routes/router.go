package routes

import (
	"crud_mongo/app/http/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, bc *controllers.BookCnt) {
	// Book
	e.POST("/book", bc.InsertBook)
}
