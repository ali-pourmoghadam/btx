package migration

import (
	"btx/bts"
	"btx/core"
	"log"

	"gorm.io/gorm"
)

func Migrate_Up_Cell_Tower_table() {

	conn, ok := core.GetDB["psql-0"].GetConnection().(*gorm.DB)

	if !ok {
		log.Fatal("GetConnection did not return *gorm.DB")
	}

	err := conn.AutoMigrate(&bts.CellTower{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

}

func Migrate_Down_Cell_Tower_table() {}

func Migrate_Up_Cell_Tower_data() {

}
