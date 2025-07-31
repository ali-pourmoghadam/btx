package migration

import (
	"btx/bts"
	"btx/core"
	"log"

	"gorm.io/gorm"
)

func Migrate_Cell_Tower() {

	conn, ok := core.GetDB["psql-0"].GetConnection().(*gorm.DB)

	if !ok {
		log.Fatal("GetConnection did not return *gorm.DB")
	}

	err := conn.AutoMigrate(&bts.CellTower{})

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("cell_tower migration runs seccessfully")
}
