package gorm_test

import (
	"cschain-bond/entity"
	"cschain-bond/gorm"
	"cschain-bond/utils"
	"fmt"
	"testing"
)

func TestFindAll(t *testing.T) {
	var b gorm.BondVariety
	bs := b.FindAll()
	fmt.Print(bs)

	fmt.Println("=================")
	var r gorm.RepurchaseVariety
	rs := r.FindAll()
	fmt.Print(rs)
}

func TestInsert(t *testing.T) {
	db := utils.GetConnection()
	user := entity.BondVarietyDimension{
		Name:     "cjj",
		ParentId: 1,
		Level:    1,
	}
	db.NewRecord(user)
	db.Create(&user)
}

func TestCreateTable(t *testing.T) {
	db := utils.GetConnection()
	db.CreateTable(&entity.BondTransaction{})
}
