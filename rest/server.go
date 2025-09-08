package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config){

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cros,	
		middleware.Logger,
	)
  
	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)


  initRoutes(mux, manager)


	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrapperMux)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
		os.Exit(1)
	}
}