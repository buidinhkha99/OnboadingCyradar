package main

import (
	"SavingBooks/app/handlers"
	"SavingBooks/app/pkg"
	"flag"

	log "github.com/sirupsen/logrus"
)

func main() {
	err := pkg.LoadConfig()
	if err != nil {
		log.Error("Error loading cofig")

	}
	setDataMemory := flag.Bool("setdatamemory", false, "Set data in memory hear !")
	saveDataMongodb := flag.Bool("setdatamongodb", false, "Set data in mongo hear !")
	flag.Parse()
	if *setDataMemory {
		handlers.SetDataMemory()
	}
	if *saveDataMongodb {
		handlers.SaveDataMongodb()
	}

}
