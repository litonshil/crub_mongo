package container

import (
	"crud_mongo/infra/conn"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	// db := conn.Db()
	_ = conn.Db()

	// costRepo := repo.NewCostRepository(db)

	// costSvc := costsvc.NewCostService(costRepo)

	// costCr := controllers.NewCostController(costSvc)

	// routes.Init(e, costCr)

}
