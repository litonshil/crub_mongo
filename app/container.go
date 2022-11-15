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
	// _ = conn.Db()

	bookRepo := repo.NewCostRepository(db)

	bookSvc := svc.NewCostService(bookRepo)

	bookCr := controllers.NewCostController(bookSvc)

	routes.Init(e, bookCr)

}
