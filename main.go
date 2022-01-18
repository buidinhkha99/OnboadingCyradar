package main

import (
	"BookShop/app/handlers"
	"BookShop/app/migrate"
	"BookShop/app/router"
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Start..")
	setDataMemory := flag.Bool("setdatamemory", false, "Set data in memory hear !")
	restApi := flag.Bool("restapi", false, "Use API to connect to many databases !")
	setupDatabaseMysql := flag.Bool("setupmysql", false, "Set table for mysql database !")
	setupDatabasePostGres := flag.Bool("setuppostgres", false, "Set table for postgres database !")
	flag.Parse()
	if *setDataMemory {
		handlers.SetDataMemory()
	}
	if *restApi {
		router.Run()
	}
	if *setupDatabaseMysql {
		migrate.CreateTableMySql()
	}
	if *setupDatabasePostGres {
		migrate.CreateTablePostgres()
	}

}
