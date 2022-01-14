package main

import (
	"SavingBooks/app/handlers"
	"flag"
)

func main() {
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
