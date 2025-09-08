package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve(){
  manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cros,	
		middleware.Logger,
	)
  
	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)


  initRoutes(mux, manager)
	
	fmt.Println("Server running on : 8080")
	err := http.ListenAndServe(":8080", wrapperMux)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
	}
}