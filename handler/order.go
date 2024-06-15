package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ArjunMalhotra07/apis.git/structs"
)

type Order struct {
}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create an order")
	response := structs.Response{Message: "Created an order"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List all order")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("List an order by iD")
}

func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update an order by id")
}

func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete an order by id")
}
