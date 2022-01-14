package main

import (
	"BookShop/app/handlers"
	"BookShop/app/migrate"
	"BookShop/app/router"
	"flag"
)

func main() {
	setDataMemory := flag.Bool("setdatamemory", false, "Set data in memory hear !")
	restApi := flag.Bool("restapi", false, "Use API to connect to many databases !")
	setupDatabaseMysql := flag.Bool("setupmysql", false, "Set table for mysql database !")
	flag.Parse()
	if *setDataMemory {
		handlers.SetDataMemory()
	}
	if *restApi {
		router.Run()
	}
	if *setupDatabaseMysql {
		migrate.CreateTable()
	}

}
