package dao_test

import (
	"cschain-bond/dao"
	"cschain-bond/models"
	"cschain-bond/utils"
	"fmt"
	"testing"
)

func TestFindAll(t *testing.T) {
	var b dao.BondVarietyDao
	bs := b.FindAll()
	fmt.Print(bs)

	fmt.Println("=================")
	var r dao.RepurchaseVarietyDao
	rs := r.FindAll()
	fmt.Print(rs)
}

func TestInsert(t *testing.T) {
	db := utils.GetConnection()
	user := models.BondVarietyDimension{
		Name:     "cjj",
		ParentId: 1,
		Level:    1,
	}
	db.NewRecord(user)
	db.Create(&user)
}

func TestCreateTable(t *testing.T) {
	db := utils.GetConnection()
	db.CreateTable(&models.BondTransaction{})
}
