package server

import (
	"encoding/json"
	"io"
	"net/http"
	"run/db"
	"run/dto"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(res http.ResponseWriter, req *http.Request) {
	body, err = io.ReadAll(req.Body)
	if err != nil {
		log.Error("Login ReadAll Error:", err)
	}
	var user dto.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Error("Login Unmarshal Error:", err)
	}
	var resUser dto.User
	var ress dto.Response
	filter := bson.M{"mail": user.Mail}
	err = db.FindOne(&resUser, filter)
	if err != nil {
		ress.Error = "User not found"
		ress.Verification = false
	} else if user.Password != resUser.Password {
		ress.Error = "Incorrect Password"
		ress.Verification = false
	} else if resUser.Role != user.Role {
		ress.Error = "Role Mismatch"
		ress.Verification = false
	} else {
		ress.Verification = true
		ress.Result = "Loged in SuccesFull"
	}
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ress)
}
