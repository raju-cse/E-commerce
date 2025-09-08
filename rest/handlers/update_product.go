package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)



func UpdateProduct(w http.ResponseWriter, r *http.Request){	
 
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil{
		http.Error(w, "please give me valid product id", 400)
		return
	}

		var newProduct database.Product  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = pId
	
	database.Update(newProduct)
	
	util.SendData(w, "Successfully updated Product", 201)

}