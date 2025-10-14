package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
	
	var req ReqLogin  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req Body")
		return
	}

  usr, err := h.svc.Find(req.Email, req.Password)

	if err != nil{
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if usr == nil{
		util.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}


	accessToken, err := util.CeateJwt(h.cnf.JwtSecretKey,util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	 })

	 if err != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	 }

	util.SendData(w, http.StatusCreated, accessToken)
	
}