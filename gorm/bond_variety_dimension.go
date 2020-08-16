package gorm

import (
	"cschain-bond/entity"
	"cschain-bond/utils"
)

type BondVariety struct {
}

func (b BondVariety) FindAll() []entity.BondVarietyDimension {
	db := utils.GetConnection()
	defer db.Close()

	var bondVarietys []entity.BondVarietyDimension
	db.Find(&bondVarietys)
	return bondVarietys
}
