package models

type BondRepurchaseVarietyDimension struct {
	ID       int `gorm:"primary_key"`
	Name     string
	ParentId uint
	Level    int
}
