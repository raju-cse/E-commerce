package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request){
	
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "please give me valid product id")
		return
		
	}

  product, err := h.productRepo.Get(pId)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

  if product == nil{
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}


	util.SendData(w, http.StatusOK, product)
}