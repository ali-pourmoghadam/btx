package migration

var dataMigration = map[string]func(){
	"cell_tower": Migrate_Cell_Tower_data,
}

func Migration_data() {

}

// REGISTER ALL MIGRATION HERE (i need script that handle migration entirely)

func Migration_table() {

	Migrate_Up_Cell_Tower_table()

}
