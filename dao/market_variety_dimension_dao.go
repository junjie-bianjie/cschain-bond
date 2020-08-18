package dao

import (
	"cschain-bond/models"
	"cschain-bond/utils"
)

type MarketVarietyDao struct {
}

func (m MarketVarietyDao) FindAll() []models.MarketVarietyDimension {
	db := utils.GetConnection()
	defer db.Close()

	var markets []models.MarketVarietyDimension
	db.Find(&markets)
	return markets
}
