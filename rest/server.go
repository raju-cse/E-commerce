package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct{
	cnf            *config.Config
  productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NerServer(
	cnf        *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler  *review.Handler,
) *Server{
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
} 

func (server *Server) Start(){

	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cros,	
		middleware.Logger,
	)
  
	mux := http.NewServeMux()
	wrapperMux := manager.WrapMux(mux)


 server.productHandler.RegsterRoutes(mux, manager)
 server.userHandler.RegsterRoutes(mux, manager)
 server.reviewHandler.RegsterRoutes(mux, manager)


	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	fmt.Println("Server running on port", addr)
	err := http.ListenAndServe(addr, wrapperMux)

	if err != nil{
		fmt.Println("Error startin the Servr :", err)
		os.Exit(1)
	}
}