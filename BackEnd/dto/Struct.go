package dto

import "time"

type User struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Mail        string    `json:"mail" bson:"mail,omitempty"`
	Role        string    `json:"role" bson:"role,omitempty"`
	Password    string    `json:"password" bson:"password,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" bson:"updatedDate,omitempty"`
}

type Customer struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Mail        string    `json:"mail" bson:"mail,omitempty"`
	Role        string    `json:"role" bson:"role,omitempty"`
	Address     Address   `json:"address" bson:"address,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" bson:"updatedDate,omitempty"`
}

type Address struct {
	Area    string `json:"area" bson:"line,omitempty"`
	City    string `json:"city" bson:"city,omitempty"`
	Zipcode string `json:"pincode" bson:"zipcode,omitempty"`
	Contact string `json:"phone" bson:"contact,omitempty"`
}
type Response struct {
	Result       interface{}
	Error        string
	Verification bool
}
