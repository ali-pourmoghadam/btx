package bts

import (
	"btx/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type CellTower struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Radio      string
	MCC        int
	MNC        int
	LAC        int
	CID        int
	Lat        float64
	Lon        float64
	Range      int
	Samples    int
	Changeable bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (c *CellTower) GetAll() ([]CellTower, error) {

	var towers []CellTower

	db, ok := core.GetDB["psql-0"].GetConnection().(*gorm.DB)

	if !ok {
		return nil, fmt.Errorf("convert type to gorm db failed")
	}

	result := db.Find(&towers)

	if result.Error != nil {
		return nil, result.Error
	}

	return towers, nil
}
