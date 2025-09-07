package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request){
	
	productID := r.PathValue("productId")

	pId, err := strconv.Atoi(productID)

	if err != nil{
		http.Error(w, "please give me valid product id", 400)
		return
	}

	for _, product := range database.ProductList{
		if product.ID == pId{
			util.SendData(w, product, 200)
			return
		}
	}

  util.SendData(w, "Data not found", 404)
}