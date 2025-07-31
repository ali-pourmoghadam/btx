package core

import (
	crossdomain "btx/cross-domain"
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type DATABASE interface {
	MakeConnection()
	GetConnection() interface{}
}

// YIELDS CONNECTION POOLS OF DATABASES

var dbs = []DATABASE{&Psql{}}

// MKAE DATBASES GLOBAL

var GetDB = map[string]DATABASE{
	"psql-0": dbs[0],
}

type Psql struct {
	db *gorm.DB
}

/* ========================PSQL SETUP======================= */

// INITIATE CONECTION INTO PSQL

func (p *Psql) MakeConnection() {

	cfg, err := crossdomain.LoadConfigFile()

	if err != nil {
		log.Fatal(err)
	}

	section := cfg.Section("POSTGRESQL")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		section.Key("USERNAME").String(),
		section.Key("PASSWORD").String(),
		section.Key("HOST").String(),
		section.Key("PORT").String(),
		section.Key("DBNAME").String())

	p.db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	log.Println("postgress connected successfully")

}

func (p *Psql) GetConnection() interface{} {
	return p.db
}

/* ========================X DB SETUP======================= */
