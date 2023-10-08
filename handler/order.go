package handler

import (
	"fmt"
	"net/http"
)


type Order struct{}

func(o *Order) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("create an order")
}

func(o *Order) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("listing all orders")
}

func(o *Order) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("retrieving specific order")
}

func(o *Order) UpdateByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("updating order")
}

func(o *Order) DeleteByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("deleting specific order")
}
