package core

func Init() {

	dbSetup()
}

// RANGE OVER DATABSE OBJECT LIST AND MAKE CONNECTION AVAILABLE

func dbSetup() {

	for _, db := range dbs {
		db.MakeConnection()
	}
}

// TAKES CARE DATBASE  MIGRATIONS

func migrate() {

	migrateBTX()

}

func migrateBTX() {

}
