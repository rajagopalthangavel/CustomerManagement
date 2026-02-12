package main

import (
	"net/http"
	cof "run/config"
	"run/db"
	s "run/service"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	cof.LoadConfig()
	db.MongoInit()
	router := mux.NewRouter()
	router.HandleFunc("/login", s.Login)
	router.HandleFunc("/createUser", s.CreateUser)
	router.HandleFunc("/listUser", s.ListUser)
	router.HandleFunc("/listOneUser", s.ListOneUser)
	router.HandleFunc("/userCount", s.UserCount)
	router.HandleFunc("/updateUser", s.UpdateUser)
	router.HandleFunc("/deleteUser", s.DeleteUser)
	router.HandleFunc("/addCustomer", s.AddCustomer)
	router.HandleFunc("/listCustomer", s.ListCustomer)
	router.HandleFunc("/listOneCustomer", s.ListOneCustomer)
	router.HandleFunc("/customerCount", s.CustomerCount)
	router.HandleFunc("/updateCustomer", s.UpdateCustomer)
	router.HandleFunc("/deleteCustomer", s.DeleteCustomer)
	headersOK := handlers.AllowedHeaders([]string{"Content-Type",
		"access-control-allow-origin",
		"access-control-request-headers",
		"Access-Control-Allow-Headers",
		"Access-Control-Allow-Methods",
		"Authorization", "X-Requested-With, X-CSRF-Token"})
	originsOK := handlers.AllowedOrigins([]string{"*"})
	methodsOK := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	loggedRouter := handlers.CORS(headersOK, originsOK, methodsOK)(router)
	log.Info("Server is running on port :", viper.GetString("Port"))
	if err := http.ListenAndServe(":"+viper.GetString("Port"), loggedRouter); err != nil {
		log.Errorf("Server connection error: %v\n", err)
	}
}
