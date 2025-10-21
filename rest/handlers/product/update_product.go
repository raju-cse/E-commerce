package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct{
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
}


func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request){	
 
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil{
		util.SendError(w, http.StatusBadRequest, "please give me valid product id")
		return
	}

		var req ReqUpdateProduct  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "plz give me valid json")
		return
	}
  
	_, err = h.svc.Update(domain.Product{
		ID:          pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		IngUrl:      req.IngUrl,
	})

	if err != nil {
		fmt.Println(err)
		util.SendError(w,http.StatusInternalServerError, "Interanl server error" )
		return
	}
	
	// database.Update(newProduct)
	
	util.SendData(w, http.StatusOK, "Successfully updated Product")

}