package repository

import (
	"crud_mongo/app/domain"
	"crud_mongo/app/models"
	"crud_mongo/app/utils/errutil"
	logger "crud_mongo/infra/logger"
	"encoding/json"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type costs struct {
	DB *gorm.DB
}

func NewCostRepository(dbc *gorm.DB) domain.ICostRepo {
	return &costs{
		DB: dbc,
	}
}

type propObj struct {
	Name  string `json:"prop_name"`
	Price int    `json:"prop_price"`
}

func (cr *costs) InsertDailyCost(req *models.Cost) error {
	// [{"prop_name":"sfs","prop_price":526},{"prop_name":"sfdfds","prop_price":56}]
	properties := req.Properties
	var props []propObj
	err := json.Unmarshal(properties, &props)
	if err != nil {
		logger.Error(err)
		return err
	}
	totalPrice := 0
	for _, item := range props {
		totalPrice += item.Price
	}

	req.Total = totalPrice
	res := cr.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		UpdateAll: true,
	}).Create(&req)

	if res.Error != nil {
		logger.Error(res.Error)
		return errutil.ErrUserCreate
	}
	return nil
}

func (cr *costs) GetDailyCost() ([]models.Cost, error) {
	var data []models.Cost
	if err := cr.DB.Model(&models.Cost{}).Find(&data).Error; err != nil {
		logger.Error("error occurred while getting daily cost", err)
		return nil, err
	}
	return data, nil
}
