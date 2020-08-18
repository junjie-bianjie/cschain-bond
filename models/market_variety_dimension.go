package models

type MarketVarietyDimension struct {
	Id     int `gorm:"primary_key"`
	Market string
}
