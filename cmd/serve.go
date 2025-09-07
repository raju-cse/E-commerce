package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
  
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)


	// mux.Handle("GET /route", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))
	
	// mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	// mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProduct))))
	// mux.Handle("GET /products/{productId}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductByID))))
	

  globalRouter := global_router.GlobalRouter(mux)
	
	
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", globalRouter)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
	}
}