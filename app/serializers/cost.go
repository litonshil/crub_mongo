package serializers

import "gorm.io/datatypes"

type PropertiesObj struct {
	PropName  string `json:"prop_name"`
	PropPrice int    `json:"prop_price"`
}

type CostPayload struct {
	ID         int             `json:"id"`
	Name       string          `json:"name"`
	Properties []PropertiesObj `json:"properties"`
	Total      int             `json:"total"`
	Date       datatypes.Date  `json:"date"`
}

type CostSummary struct {
	TotalCost int
	IndvPay   int
}
