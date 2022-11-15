package container

import (
	"crud_mongo/app/http/controllers"
	"crud_mongo/app/http/routes"
	repo "crud_mongo/app/repository"
	svc "crud_mongo/app/service"
	"crud_mongo/infra/conn"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	db := conn.Db()

	bookRepo := repo.NewBookRepository(db)

	bookSvc := svc.NewBookService(bookRepo)

	bookCr := controllers.NewBookController(bookSvc)

	routes.Init(e, bookCr)

}
