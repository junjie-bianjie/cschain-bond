package models

import "time"

type BondTransaction struct {
	ID                 int `gorm:"primary_key"`
	NftId              string
	SourceType         int
	DenomId            string
	Owner              string
	Uri                string
	Visible            bool
	Amount             float64
	Market             int
	StartDate          time.Time
	EndDate            time.Time
	PeriodCategory     string
	BondCategory       int
	RepurchaseCategory int
}
