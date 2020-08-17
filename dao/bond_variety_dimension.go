package dao

import (
	"cschain-bond/models"
	"cschain-bond/utils"
)

type BondVarietyDao struct {
}

func (b BondVarietyDao) FindAll() []models.BondVarietyDimension {
	db := utils.GetConnection()
	defer db.Close()

	var bondVarietys []models.BondVarietyDimension
	db.Find(&bondVarietys)
	return bondVarietys
}
