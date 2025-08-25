package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request){
	
	var newProduct database.Product  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList)+1
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
	
}