package routes

import (
	"crud_mongo/app/http/controllers"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo, cc *controllers.CostCnt) {
	// Cost
	e.POST("cost/v1/daily-cost", cc.DailyCost)
	e.GET("cost/v1/daily-cost", cc.GetDailyCost)
}
