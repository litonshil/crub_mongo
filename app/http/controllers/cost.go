package controllers

import (
	"crud_mongo/app/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CostCnt struct {
	costSvc domain.ICostSvc
}

func NewCostController(costSvc domain.ICostSvc) *CostCnt {
	costc := &CostCnt{
		costSvc: costSvc,
	}
	return costc
}

// swagger:route GET /cost get cost
// Get companies list
// responses:
//  401: errorResponse
//  200: genericSuccessResponse
func (cc *CostCnt) DailyCost(c echo.Context) error {
	// var req serializers.CostPayload
	// var err error

	// if err = c.Bind(&req); err != nil {
	// 	logger.Error(err)
	// 	return c.JSON(http.StatusBadRequest, msgutil.RequestBodyParseErrorResponseMsg())
	// }
	// err = cc.costSvc.InsertDailyCost(&req)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, msgutil.EntityCreationFailedMsg("User"))
	// }

	return c.NoContent(http.StatusCreated)
}

func (cc *CostCnt) GetDailyCost(c echo.Context) error {
	// res, err := cc.costSvc.GetDailyCost()
	// if err != nil {
	// 	logger.Error(err)
	// 	return err
	// }
	// return c.JSON(http.StatusOK, res)
	return nil
}

func (cc *CostCnt) GetCostSummary(c echo.Context) error {
	return nil
}
