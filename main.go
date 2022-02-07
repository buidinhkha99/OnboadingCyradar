package main

import (
	handlersBook "BookShop/cmd/microservice/book/app/handlers"
	routerBook "BookShop/cmd/microservice/book/app/router"
	routerCategory "BookShop/cmd/microservice/category/app/router"
	"BookShop/migrate"
	"BookShop/program"
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
	restApiBook := flag.Bool("apibook", true, "Rest API for book")
	restApiCategory := flag.Bool("apicategory", false, "Rest API for category")
	flag.Parse()

	if *setDataMemory {
		program.SetDataMemory()
	}

	if *setupDatabaseMysql {
		migrate.CreateTableMySql()
	}

	if *setupDatabasePostGres {
		migrate.CreateTablePostgres()
	}

	if *restApiBook {
		go handlersBook.CheckPubSub()
		routerBook.Run()
	}

	if *restApiCategory {
		// go handlersCat.CheckPubSub()
		routerCategory.Run()
	}

}
