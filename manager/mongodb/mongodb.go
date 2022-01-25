package mongodb

import (
	"BookShop/connect"
	"BookShop/model"
	"BookShop/pkg"
	"BookShop/pkg/cache"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	Bookconnection      *mongo.Collection
	CatergoryCollection *mongo.Collection
	GroupBook           *mongo.Collection
	Ctx                 context.Context
	Cancel              context.CancelFunc
	Client              *mongo.Client
}

func (mg *Mongo) Connect() error {
	client, ctx, cancel, err := connect.Connect()
	if err != nil {
		log.Errorf("Cannot connect to mongoDB: %v", err)
		return err
	}
	database := viper.GetString("mongodb.database")
	connectionBook := viper.GetString("mongodb.connection_book")
	connectionCategory := viper.GetString("mongodb.connection_category")
	connectionGroupBook := viper.GetString("mongodb.connection_group_book")
	bookDatabase := client.Database(database)
	bookCollection := bookDatabase.Collection(connectionBook)
	catergoryCollection := bookDatabase.Collection(connectionCategory)
	groupConliection := bookDatabase.Collection(connectionGroupBook)
	mg.Cancel = cancel
	mg.Ctx = ctx
	mg.Bookconnection = bookCollection
	mg.CatergoryCollection = catergoryCollection
	mg.GroupBook = groupConliection
	mg.Client = client

	return nil
}
func (mg *Mongo) Close() {
	defer connect.Close(mg.Client, mg.Ctx, mg.Cancel)
}
func (mg *Mongo) CreateBook(book *model.Book) error {
	mg.Connect()
	id := pkg.GenID()
	book.ID = id
	result, err := mg.Bookconnection.InsertOne(mg.Ctx, book)
	if err != nil {
		log.Errorf("Cannot insert book to mongoDB: %v, result: %v", err, result)
		return err
	}
	mg.Close()
	return nil
}

func (mg *Mongo) UpdateBook(book model.Book) error {
	mg.Connect()
	data, err := ConvertBson(book)
	if err != nil {
		log.Errorf("Cannot update book to mongoDB: %v", err)
		return err
	}
	result, err := mg.Bookconnection.UpdateMany(
		mg.Ctx,
		bson.M{"_id": book.ID},
		bson.D{{Key: "$set", Value: data}},
	)
	if err != nil {
		log.Errorf("Cannot update book to mongoDB: %v, result: %v", err, result)
		return err
	}
	return nil
}

func (mg *Mongo) DeleteBook(book model.Book) error {
	mg.Connect()

	result, err := mg.Bookconnection.DeleteOne(context.TODO(), bson.D{{"_id", book.ID}})
	if err != nil {
		log.Errorf("Cannot detele book to mongoDB: %v, result: %v", err, result)
		return err
	}
	mg.Close()
	return nil

}

func (mg *Mongo) GetAllBook() ([]model.Book, error) {
	mg.Connect()
	var books []model.Book
	result, err := mg.Bookconnection.Find(mg.Ctx, bson.D{})
	if err != nil {
		log.Errorf("Cannot update book to mongoDB: %v, result: %v", err, result)
		return nil, err
	}
	for result.Next(mg.Ctx) {
		var book model.Book
		err := result.Decode(&book)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	mg.Close()
	return books, nil
}

func (mg *Mongo) GetBook(ID string) (book model.Book, err error) {
	mg.Connect()
	err = mg.Bookconnection.FindOne(mg.Ctx, bson.D{{"_id", ID}}).Decode(&book)
	if err != nil {
		return book, err
	}

	mg.Close()
	return book, nil
}

func (mg *Mongo) GetTopBook() string {
	mg.Connect()
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		var books []model.Book
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{"rate", -1.1}})
		findOptions.SetLimit(8)
		cursor, err := mg.Bookconnection.Find(mg.Ctx, bson.D{}, findOptions)
		if err = cursor.All(context.TODO(), &books); err != nil {
			log.Errorf("Error when get group book %v, Error: %v", cursor, err)
			return ""
		}
		b, _ := json.Marshal(books)
		err = cache.InsertData(key, string(b))
		if err != nil {
			log.Error("Error when add data in remote cache")
		}
		return string(b)
	}
	mg.Close()
	return data
}

func (mg *Mongo) GetGroupBookByID(id string) (groupBook []model.GroupBook, err error) {
	mg.Connect()
	cursor, err := mg.GroupBook.Find(context.TODO(), bson.D{{"bookid", id}})
	if err != nil {
		log.Errorf("Error when get group book %v, Error: %v", cursor, err)
		return groupBook, err
	}
	if err = cursor.All(context.TODO(), &groupBook); err != nil {
		log.Errorf("Error when get group book %v, Error: %v", cursor, err)
		return groupBook, err
	}
	mg.Close()
	return groupBook, nil
}

// handler Catergory

func (mg *Mongo) CreateCategory(category model.Category) error {
	mg.Connect()
	id := pkg.GenID()
	category.ID = id
	result, err := mg.CatergoryCollection.InsertOne(mg.Ctx, category)
	if err != nil {
		log.Errorf("Cannot insert category to mongoDB: %v, result: %v", err, result)
		return err
	}
	mg.Close()
	return nil
}

func (mg *Mongo) UpdateCategory(category model.Category) error {
	mg.Connect()
	data, err := ConvertBson(category)
	if err != nil {
		log.Errorf("Cannot update catergory to mongoDB: %v", err)
		return err
	}
	result, err := mg.CatergoryCollection.UpdateMany(
		mg.Ctx,
		bson.M{"_id": category.ID},
		bson.D{{Key: "$set", Value: data}},
	)
	if err != nil {
		log.Errorf("Cannot update category to mongoDB: %v, result: %v", err, result)
		return err
	}
	return nil
}

func (mg *Mongo) DeleteCategory(category model.Category) error {
	mg.Connect()
	result, err := mg.CatergoryCollection.DeleteOne(context.TODO(), bson.D{{"_id", category.ID}})
	if err != nil {
		log.Errorf("Cannot detele book to mongoDB: %v, result: %v", err, result)
		return err
	}
	mg.Close()
	return nil
}

func (mg *Mongo) GetAllCategory() ([]model.Category, error) {
	mg.Connect()
	var categorys []model.Category
	result, err := mg.CatergoryCollection.Find(mg.Ctx, bson.D{})
	if err != nil {
		log.Errorf("Cannot update book to mongoDB: %v, result: %v", err, result)
		return nil, err
	}
	for result.Next(mg.Ctx) {
		var category model.Category
		err := result.Decode(&category)
		if err != nil {
			return nil, err
		}
		categorys = append(categorys, category)
	}
	mg.Close()
	return categorys, nil
}

func (mg *Mongo) GetCategory(ID string) (category model.Category) {
	mg.Connect()
	err := mg.CatergoryCollection.FindOne(mg.Ctx, bson.D{{"_id", ID}}).Decode(&category)
	if err != nil {
		log.Errorf("Error Get category: %v", err)
		return category
	}
	mg.Close()
	return category
}

func (mg *Mongo) UpdateGroupBook(groupBook model.GroupBook) error {
	mg.Connect()

	data, err := ConvertBson(groupBook)
	if err != nil {
		log.Errorf("Cannot update group book to mongoDB: %v", err)
		return err
	}

	result, err := mg.GroupBook.UpdateMany(mg.Ctx,
		bson.M{"_id": groupBook.ID},
		bson.D{{Key: "$set", Value: data}})
	if err != nil {
		log.Errorf("Cannot Update group book to mongoDB: %v, result: %v", err, result)
		return err
	}

	mg.Close()
	return nil

}

func (mg *Mongo) CreateGoupBook(groupBook model.GroupBook) error {
	mg.Connect()
	id := pkg.GenID()
	groupBook.ID = id
	result, err := mg.GroupBook.InsertOne(mg.Ctx, groupBook)
	if err != nil {
		log.Errorf("Cannot insert Group to mongoDB: %v, result: %v", err, result)
		return err
	}
	mg.Close()
	return nil
}

func (mg *Mongo) GetBookWithCatergory(ID string) ([]model.Book, error) {
	mg.Connect()
	var groupBook []model.GroupBook
	var books []model.Book
	cursor, err := mg.GroupBook.Find(context.TODO(), bson.D{{"category_id", ID}})
	if err != nil {
		log.Errorf("Error when get group book %v, Error: %v", cursor, err)
		return books, err
	}
	if err = cursor.All(context.TODO(), &groupBook); err != nil {
		log.Errorf("Error when get group book %v, Error: %v", cursor, err)
		return books, err
	}

	for _, group := range groupBook {
		book, _ := mg.GetBook(group.BookID)
		books = append(books, book)
	}

	return books, nil
}

// convert data to bson
func ConvertBson(v interface{}) (bson.M, error) {
	pByte, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return nil, err
	}
	return update, err
}
