package db

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"reflect"
	"run/dto"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx context.Context
var client *mongo.Client
var err error

const CUST_PREFIX = "CUST"
const USER_PREFIX = "USER"

func MongoInit() {
	ctx = context.TODO()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("MongoUrl")))
	if err != nil {
		log.Error("db connection Error:", err)
	} else {
		log.Println("db connected")
	}
	if client != nil {
		CreateIndexUser()
		CreateIndexForCustomer()
		DefaultUser()
	}
}
func CreateIndexUser() {
	coll := client.Database(viper.GetString("Datebase")).Collection("User")
	ctx := context.Background()
	mod := mongo.IndexModel{
		Keys:    bson.M{"mail": +1},
		Options: options.Index().SetUnique(true),
	}
	str, err := coll.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Error(err)
	} else {
		log.Println("Index Created for User:", str)
	}

}

func CreateIndexForCustomer() {
	coll := client.Database(viper.GetString("Datebase")).Collection("Customer")
	ctx := context.Background()
	mod := mongo.IndexModel{
		Keys:    bson.M{"mail": +1},
		Options: options.Index().SetUnique(true),
	}
	str, err := coll.Indexes().CreateOne(ctx, mod)
	if err != nil {
		log.Error(err)
	} else {
		log.Println("Index Created for Customer:", str)
	}

}

func Coll(i interface{}) *mongo.Collection {
	value := reflect.TypeOf(i)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}
	}
	if value.Kind() == reflect.Slice {
		value = value.Elem()
		if value.Kind() == reflect.Ptr {
			value = value.Elem()
			if value.Kind() == reflect.Ptr {
				value.Elem()
			}
		}
		if value.Kind() == reflect.Slice {
			value = value.Elem()
		}
	}
	valuName := value.Name()
	coll := client.Database(viper.GetString("Datebase")).Collection(valuName)
	return coll
}

func DefaultUser() {
	data := dto.User{
		Id:          GenerateRandomID(USER_PREFIX),
		Name:        "Admin",
		Mail:        "admin@gmail.com",
		Password:    "Test@123",
		Role:        "ROLE_ADMIN",
		CreatedDate: time.Now(),
		UpdatedDate: time.Now(),
	}
	coll := Coll(dto.User{})
	coll.InsertOne(ctx, data)
}

func Insert(i interface{}) (*mongo.InsertOneResult, error) {
	coll := Coll(i)
	var res *mongo.InsertOneResult
	res, err = coll.InsertOne(ctx, i)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindAll(i interface{}, filter bson.M) error {
	coll := Coll(i)
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	err = cur.All(ctx, i)
	if err != nil {
		return err
	}
	return nil
}

func FindOne(i interface{}, filter bson.M) error {
	coll := Coll(i)
	cur := coll.FindOne(ctx, filter)
	err = cur.Decode(i)
	if err != nil {
		return err
	}
	return nil
}

func CountDocuments(i interface{}, filter bson.M) (int64, error) {
	coll := Coll(i)
	var res int64
	res, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res, nil
}

func UpdateOne(i interface{}, filter bson.M, query bson.M) (*mongo.UpdateResult, error) {
	coll := Coll(i)
	var res *mongo.UpdateResult
	res, err = coll.UpdateOne(ctx, filter, query)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DeleteOne(i interface{}, filter bson.M) (*mongo.DeleteResult, error) {
	coll := Coll(i)
	var res *mongo.DeleteResult
	res, err = coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GenerateRandomID(prefif string) string {
	b := make([]byte, 3)
	rand.Read(b)
	return prefif + "_" + hex.EncodeToString(b)
}
