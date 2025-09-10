package handlers

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request){
	
	var reqLogin ReqLogin  //Create struct instance 

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil{
		fmt.Println(err)
		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	 usr := database.Find(reqLogin.Email, reqLogin.Password)
   
	 if usr == nil{
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	 }
  
	cnf :=  config.GetConfig()

	accessToken, err := util.CeateJwt(cnf.JwtSecretKey,util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
	 })

	 if err != nil{
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	 }

	util.SendData(w, accessToken, http.StatusCreated)
	
}