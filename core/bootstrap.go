package core

import "github.com/gin-gonic/gin"

/* ========================INITIAL GLOBAL OBJECTS======================= */

var R = gin.Default()

/* ========================INITIAL NECESSARY DEPENDENCIES======================= */

func Init() {

	dbSetup()
}

// RANGE OVER DATABSE OBJECT LIST AND MAKE CONNECTION AVAILABLE

func dbSetup() {

	for _, db := range dbs {
		db.MakeConnection()
	}

}
