package bts

import (
	"btx/core"
	"log"
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

func (c *CellTower) GetAll() []CellTower {

	var towers []CellTower

	db, ok := core.GetDB["psql-0"].GetConnection().(*gorm.DB)

	if !ok {
		log.Fatalf("convert type to gorm db  failed")
	}

	result := db.Find(&towers)

	if result.Error != nil {
		log.Fatalf("failed to query cell towers: %v", result.Error)
	}

	return towers
}
