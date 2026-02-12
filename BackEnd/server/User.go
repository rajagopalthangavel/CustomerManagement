package server

import (
	"encoding/json"
	"io"
	"net/http"
	"run/db"
	"run/dto"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var body []byte
var err error

func CreateUser(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("CreateUser bodyReadAll Error:", err)
	}
	var user dto.User
	var ress dto.Response
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error("CreateUser Unmarshal Error:", err)
	}

	user.Id = db.GenerateRandomID(db.USER_PREFIX)
	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()
	val, err := db.Insert(user)
	if err != nil {
		ress.Verification = false
		ress.Error = err.Error()
	} else {
		ress.Verification = true
		ress.Result = val
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func ListUser(res http.ResponseWriter, req *http.Request) {
	var user []dto.User
	filter := bson.M{}
	db.FindAll(&user, filter)
	if user == nil {
		log.Error("NOT FOUND")
	}
	json.NewEncoder(res).Encode(user)
}

func ListOneUser(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("ListOneUser bodyReadAll Error:", err)
	}
	var user dto.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error("ListOneUser Unmarshal Error:", err)
	}
	filter := bson.M{"_id": user.Id}
	err = db.FindOne(&user, filter)
	if err != nil {
		log.Error("ListOneUser FindOne Error:", err)
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(user)
}

func UserCount(res http.ResponseWriter, req *http.Request) {
	var user []dto.User
	var count int64
	var ress dto.Response
	filter := bson.M{}
	count, err = db.CountDocuments(user, filter)
	if err != nil {
		ress.Error = err.Error()
	} else {
		ress.Result = count
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func UpdateUser(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdateUser bodyReadAll Error:", err)
	}
	var user dto.User
	var ress dto.Response

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error("UpdateUser Unmarshal Error:", err)
	}
	filter := bson.M{"_id": user.Id}
	query := bson.M{"$set": bson.M{"name": user.Name, "mail": user.Mail, "password": user.Password}}
	var val *mongo.UpdateResult
	val, err = db.UpdateOne(user, filter, query)
	if err != nil {
		ress.Verification = false
		ress.Error = err.Error()
	} else {
		ress.Verification = true
		ress.Result = val
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func DeleteUser(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("DeleteUser bodyReadAll Error:", err)
	}
	var user dto.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error("DeleteUser Unmarshal Error:", err)
	}
	filter := bson.M{"_id": user.Id}
	var val *mongo.DeleteResult
	var ress dto.Response
	val, err = db.DeleteOne(user, filter)
	if err != nil {
		ress.Verification = false
		ress.Error = err.Error()
	} else {
		ress.Verification = true
		ress.Result = val
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}
