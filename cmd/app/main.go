package main

import (
	"btx/core"
	"btx/routes"
)

func main() {

	// SETUP  CORE DEPENDENCIES
	core.Init()

	// REGISTER ROUTES
	routes.GetCellTower()

	// START WEBSERVER
	core.R.Run(":8080")

}
