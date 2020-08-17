package models

type BondVarietyDimension struct {
	ID       int `gorm:"primary_key"`
	Name     string
	ParentId uint
	Level    int
}
