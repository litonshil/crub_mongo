package costsvc

import (
	"crud_mongo/app/domain"
	"crud_mongo/app/models"
	"crud_mongo/app/serializers"
	"crud_mongo/app/utils/errutil"
	methods "crud_mongo/app/utils/methodutil"
	logger "crud_mongo/infra/logger"
)

type costs struct {
	costRepo domain.ICostRepo
}

func NewCostService(costRepo domain.ICostRepo) domain.ICostSvc {
	return &costs{
		costRepo: costRepo,
	}
}

func (cs *costs) InsertDailyCost(req *serializers.CostPayload) error {
	cost := &models.Cost{}
	respErr := methods.CopyStruct(req, &cost)
	if respErr != nil {
		return respErr
	}

	if err := cs.costRepo.InsertDailyCost(cost); err != nil {
		logger.Error(err)
		return errutil.ErrCostCreate
	}

	return nil
}

func (cs *costs) GetDailyCost() ([]models.Cost, error) {
	res, err := cs.costRepo.GetDailyCost()
	return res, err
}
