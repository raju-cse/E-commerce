package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve(){

	mux := http.NewServeMux()
	
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))
	

  globalRouter := global_router.GlobalRouter(mux)
	
	
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
	}
}