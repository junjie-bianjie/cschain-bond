package dao

import (
	"cschain-bond/models"
	"cschain-bond/utils"
)

type RepurchaseVarietyDao struct {
}

func (r RepurchaseVarietyDao) FindAll() []models.BondRepurchaseVarietyDimension {
	db := utils.GetConnection()
	defer db.Close()

	var RepurchaseVarietys []models.BondRepurchaseVarietyDimension
	db.Find(&RepurchaseVarietys)
	return RepurchaseVarietys
}
