package handlers

import (
	"log"
	"net/http"
)

 func Test(w http.ResponseWriter, r *http.Request)  {
		log.Println("ami Handler : middle e print hobo")
	}

