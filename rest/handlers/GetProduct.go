package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request){
	
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil{
		http.Error(w, "please give me valid product id", 400)
		return
	}

  product := database.Get(pId)

  if product == nil{
		util.SendError(w, 404, "Product not found")
		return
	}

  // util.SendData(w, "Data not found", 404)

	util.SendData(w, product, 200)
}