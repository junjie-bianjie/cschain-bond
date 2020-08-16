package entity

import "time"

type BondRepurchaseVarietyDimension struct {
	ID       int `gorm:"primary_key"`
	Name     string
	ParentId uint
	Level    int
}

type BondVarietyDimension struct {
	ID       int `gorm:"primary_key"`
	Name     string
	ParentId uint
	Level    int
}

type BondTransaction struct {
	ID                 int `gorm:"primary_key"`
	NftId              string
	SourceType         int
	DenomId            string
	Owner              string
	Uri                string
	Visible            bool
	Amount             float64
	Market             string
	StartDate          time.Time
	EndDate            time.Time
	PeriodCategory     string
	BondCategory       int
	RepurchaseCategory int
}
