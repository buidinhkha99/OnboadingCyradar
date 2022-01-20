package program

import (
	"BookShop/connect"
	"BookShop/model"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var memory = make(map[string][]interface{})

func SetDataMemory() {
	for {
		fmt.Println(`Start program...
		1. Create new book.
		2. Create new category.
		3. Get data from memory.
		4. Save data just input to MongoDb
		5. Quit.
		`)
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter your chose: ")
		var choose int
		fmt.Scan(&choose)
		switch choose {
		case 1:
			fmt.Println("Create new book !")
			fmt.Println("Enter name book: ")
			str, _ := reader.ReadString('\n')
			name := strings.Split(str, "\n")
			fmt.Println(name[0])

			fmt.Println("Enter quantily: ")
			var quantily int
			fmt.Scanln(&quantily)

			fmt.Println("Enter description: ")
			str, _ = reader.ReadString('\n')
			description := strings.Split(str, "\n")

			fmt.Println("Enter price: ")
			var price float32
			fmt.Scanln(&price)

			fmt.Println("Enter rate: ")
			var rate float32
			fmt.Scanln(&rate)

			fmt.Println("Enter image: ")
			str, _ = reader.ReadString('\n')
			image := strings.Split(str, "\n")

			SetBookMer(name[0], quantily, description[0], price, rate, image[0])
			fmt.Println("Completed")

		case 2:
			fmt.Println("Create new category!")
			fmt.Println("Enter name category: ")
			str, _ := reader.ReadString('\n')
			name := strings.Split(str, "\n")
			SetCategory(name[0])

		case 3:
			if len(memory) == 0 {
				fmt.Println("No data in memory..")
				break
			}
			data, _ := json.Marshal(memory)
			fmt.Println(string(data))

		case 4:
			err := SaveDataMongodb()
			if err != nil {
				fmt.Println("Error save data !")
				break
			}
			fmt.Println("Completed")

		case 5:
			return
		}

		fmt.Println("--------------------------------------------------------")
	}
}

func SetBookMer(name string, quantily int, description string, price float32, rate float32, image string) {
	memory["book"] = append(memory["book"], model.Book{
		Name:        name,
		Quantily:    quantily,
		Description: description,
		Price:       price,
		Rate:        rate,
		Image:       image,
	})
}

func SetCategory(name string) {
	memory["category"] = append(memory["category"], model.Category{
		Name: name,
	})
}

func SaveDataMongodb() error {
	client, ctx, cancel, err := connect.Connect()
	if err != nil {
		log.Errorf("Cannot connect to mongoDB: %v", err)
		return err
	}
	database := viper.GetString("mongodb.Database")
	connectionBook := viper.GetString("mongodb.ConnectionBook")
	connectionCategory := viper.GetString("mongodb.ConnectionCategory")
	bookDatabase := client.Database(database)
	bookCollection := bookDatabase.Collection(connectionBook)
	catergoryCollection := bookDatabase.Collection(connectionCategory)

	if len(memory) == 0 {
		return errors.New("no data in memory")
	}
	Result, err := bookCollection.InsertMany(ctx, memory["book"])
	if err != nil {
		log.Errorf("Cannot insert book to mongoDB: %v, Result: %v", err, Result)
		return err
	}
	Result, err = catergoryCollection.InsertMany(ctx, memory["category"])
	if err != nil {
		log.Errorf("Cannot insert category to mongoDB: %v, Result: %v", err, Result)
		return err
	}
	defer connect.Close(client, ctx, cancel)
	return nil
}
