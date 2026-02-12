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

func AddCustomer(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("AddCustomer ReadAll Error:", err)
	}

	var cus dto.Customer
	var ress dto.Response
	if err = json.Unmarshal(body, &cus); err != nil {
		log.Error("AddCustomer Unmarshal Error:", err)
	} else {
		cus.Role = "Customer"
		cus.Id = db.GenerateRandomID(db.CUST_PREFIX)
		cus.CreatedDate = time.Now()
		cus.UpdatedDate = time.Now()
		var val *mongo.InsertOneResult
		val, err = db.Insert(cus)
		if err != nil {
			ress.Error = err.Error()
			ress.Verification = false
		} else {
			ress.Result = val
			ress.Verification = true
		}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(ress)
	}

}

func ListCustomer(res http.ResponseWriter, req *http.Request) {
	var cus []dto.Customer
	var ress dto.Response
	filter := bson.M{}
	err = db.FindAll(&cus, filter)
	if err != nil {
		ress.Error = err.Error()
		ress.Verification = false
	} else {
		ress.Result = cus
		ress.Verification = true
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func ListOneCustomer(res http.ResponseWriter, req *http.Request) {
	var cus dto.Customer
	var ress dto.Response
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("ListOneCustomer ReadAll Error:", err)
	}
	if err = json.Unmarshal(body, &cus); err != nil {
		log.Error("ListOneCustomer Unmarshal Error:", err)
	}
	filter := bson.M{"_id": cus.Id}
	err = db.FindOne(&cus, filter)
	if err != nil {
		ress.Verification = false
		ress.Error = err.Error()
	} else {
		ress.Result = cus
		ress.Verification = true
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func CustomerCount(res http.ResponseWriter, req *http.Request) {
	var cus dto.Customer
	var ress dto.Response
	filter := bson.M{}
	if val, err := db.CountDocuments(cus, filter); err != nil {
		ress.Error = err.Error()
		ress.Verification = false
	} else {
		ress.Result = val
		ress.Verification = true
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func UpdateCustomer(res http.ResponseWriter, req *http.Request) {
	var cus dto.Customer
	var ress dto.Response
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdateCustomer ReadAll Error:", err)
	}
	if err = json.Unmarshal(body, &cus); err != nil {
		log.Error("UpdateCustomer Unmarshal Error:", err)
	}
	filter := bson.M{"_id": cus.Id}
	query := bson.M{"$set": bson.M{"name": cus.Name, "mail": cus.Mail, "address.line": cus.Address.Area, "address.city": cus.Address.City, "address.zipcode": cus.Address.Zipcode, "address.contact": cus.Address.Contact}}
	val, err := db.UpdateOne(cus, filter, query)
	if err != nil {
		ress.Error = err.Error()
		ress.Verification = false
	} else {
		ress.Verification = true
		ress.Result = val
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}

func DeleteCustomer(res http.ResponseWriter, req *http.Request) {
	var cus dto.Customer
	var ress dto.Response
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("UpdateCustomer ReadAll Error:", err)
	}
	if err = json.Unmarshal(body, &cus); err != nil {
		log.Error("UpdateCustomer Unmarshal Error:", err)
	}
	filter := bson.M{"_id": cus.Id}
	val, err := db.DeleteOne(cus, filter)
	if err != nil {
		ress.Error = err.Error()
		ress.Verification = false
	} else {
		ress.Result = val
		ress.Verification = true
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}
