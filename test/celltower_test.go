package test

import (
	"btx/bts"
	"btx/core"
	"testing"
)

func TestCellTower_GetAll(t *testing.T) {

	core.Init()

	c := bts.CellTower{}

	c.GetAll()

}
