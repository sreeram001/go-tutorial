package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sreeram001/go-tutorial/model"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

type SuccessResponse struct {
	Success bool             `json:"success"`
	Status  int              `json:"status"`
	Data    []model.Customer `json:"data"`
}

func init() {
	db, err := gorm.Open(mysql.Open("root:root@(localhost:3306)/guvi?parseTime=true"))
	if err != nil {
		fmt.Println("Db not connected")
		panic(err)
	}
	fmt.Println("Db connected Successfully")
	Db = db
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/port", helloWorld).Methods("GET")
	fmt.Println("port listening : http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	var customer []model.Customer
	Db.Find(&customer)
	SuccessRes(w, customer)
}

func SuccessRes(w http.ResponseWriter, data []model.Customer) {
	w.Header().Add("content-type", "application/json")
	resp := &SuccessResponse{
		Success: true,
		Status:  200,
		Data:    data,
	}
	json, _ := json.Marshal(resp)
	w.Write([]byte(json))
}
