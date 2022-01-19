package main

import (
	"BookShop/app_book/handlers"
	"BookShop/app_book/migrate"
	routerBook "BookShop/app_book/router"
	routerCategory "BookShop/app_category/router"
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Start..")
	setDataMemory := flag.Bool("setdatamemory", false, "Set data in memory hear !")
	// restApi := flag.Bool("restapi", false, "Use API to connect to many databases !")
	setupDatabaseMysql := flag.Bool("setupmysql", false, "Set table for mysql database !")
	setupDatabasePostGres := flag.Bool("setuppostgres", false, "Set table for postgres database !")
	restApiBook := flag.Bool("apibook", false, "Rest API for book")
	restApiCategory := flag.Bool("apicategory", false, "Rest API for category")
	flag.Parse()
	if *setDataMemory {
		handlers.SetDataMemory()
	}
	// if *restApi {
	// 	router.Run()
	// }
	if *setupDatabaseMysql {
		migrate.CreateTableMySql()
	}
	if *setupDatabasePostGres {
		migrate.CreateTablePostgres()
	}
	if *restApiBook {
		routerBook.Run()
	}
	if *restApiCategory {
		routerCategory.Run()
	}

}
