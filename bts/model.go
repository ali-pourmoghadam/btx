package bts

import "time"

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
