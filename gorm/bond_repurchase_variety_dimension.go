package gorm

import (
	"cschain-bond/entity"
	"cschain-bond/utils"
)

type RepurchaseVariety struct {
}

func (r RepurchaseVariety) FindAll() []entity.BondRepurchaseVarietyDimension {
	db := utils.GetConnection()
	defer db.Close()

	var RepurchaseVarietys []entity.BondRepurchaseVarietyDimension
	db.Find(&RepurchaseVarietys)
	return RepurchaseVarietys
}
