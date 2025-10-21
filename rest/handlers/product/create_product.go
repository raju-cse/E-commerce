package product

import (
	"ecommerce/domain"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct{
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	IngUrl       string  `json:"imageUrl"`
}


func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request){
	
	var req ReqCreateProduct  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	
 createdProduct, err := h.svc.Create(domain.Product{
	Title: req.Title,
	Description: req.Description,
	Price: req.Price,
	IngUrl: req.IngUrl,

 }) 

 if err != nil{
	util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	return
 }

	util.SendData(w, http.StatusCreated, createdProduct)
	
}

func base64UrlEncode(data []byte) string{
	return  base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}